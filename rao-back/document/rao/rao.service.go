package rao

import (
	"regexp"

	"github.com/Zenika/RAO/conv"
	"github.com/Zenika/RAO/document"
	"github.com/Zenika/RAO/log"
	"github.com/Zenika/RAO/search"
	"github.com/Zenika/RAO/tree"
	"github.com/Zenika/RAO/tree/dropbox"
	"github.com/Zenika/RAO/utils"
)

var raoFilterPattern string = `(?i)^.+/_{1,2}clients(_|\s){1}(?P<Agence>[\w&\s]+)/(?P<Client>[^/]+)/(?P<Projet>[^/]+)/REPONSE ENVOYEE`
var raoPatternFilter = regexp.MustCompile(raoFilterPattern)

// Adding support for docx documents:
// "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
var mimes = []string{"application/pdf"}

const INDEX_ID = "rao"

type RaoService struct {
	searchService search.SearchService
	convService   conv.ConvService
	treeService   tree.TreeService
}

func NewService(searchService search.SearchService, convService conv.ConvService, treeService tree.TreeService) *RaoService {
	return &RaoService{
		treeService:   treeService,
		convService:   convService,
		searchService: searchService,
	}
}

func (rao RaoService) DocFilter(doc document.IDocument) bool {
	if !utils.ArrayContainsString(mimes, doc.GetMime()) {
		return false
	}
	matches := raoPatternFilter.FindStringSubmatch(doc.GetPath())
	if nil == matches {
		return false
	}
	return true
}

func (rao RaoService) docMapper(doc document.IDocument) map[string]interface{} {
	return map[string]interface{}{
		"Content":   doc.(RaoDocument).GetContent(),
		"Title":     doc.GetTitle(),
		"Path":      doc.GetPath(),
		"Client":    doc.(RaoDocument).GetClient(),
		"Agence":    doc.(RaoDocument).GetAgence(),
		"Extension": doc.GetExtension(),
		"Mime":      doc.GetMime(),
		"Mtime":     doc.GetMtime(),
		"Bytes":     doc.(RaoDocument).GetBytes(),
		"Sum":       doc.(RaoDocument).GetSum(),
	}
}

func (rao RaoService) DocHandler(doc document.IDocument) {
	bytes, size := rao.treeService.GetEngine().(*dropbox.Dropbox).DownloadFile(doc)
	b, err := rao.convService.Convert(bytes, doc.GetMime())
	log.Error(err, log.ERROR)
	content := string(b[:])
	if "" == content {
		return // Shall we index the document if we could not extract its content ?
	}
	matches := raoPatternFilter.FindStringSubmatch(doc.GetPath())
	agence := matches[2]
	client := matches[3]
	chunks := utils.SplitString(content, 10000)
	for _, chunk := range chunks {
		raoDocument := RaoDocument{
			doc,
			document.BusinessDocument{
				doc,
				client,
				agence,
			},
			document.FullTextDocument{
				Bytes:   size,
				Sum:     utils.Md5Sum(content),
				Content: chunk,
			},
		}
		rao.searchService.Store(INDEX_ID, raoDocument, rao.docMapper)
	}
}
