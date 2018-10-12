package service

import (
	"os"
	"regexp"

	"github.com/Zenika/rao/rao-back/conv"
	"github.com/Zenika/rao/rao-back/document"
	"github.com/Zenika/rao/rao-back/document/rao"
	"github.com/Zenika/rao/rao-back/log"
	"github.com/Zenika/rao/rao-back/search"
	"github.com/Zenika/rao/rao-back/tree"
	"github.com/Zenika/rao/rao-back/tree/dropbox2"
	"github.com/Zenika/rao/rao-back/utils"
)

var RAO_FILTER_PATTERN string = os.Getenv("RAO_POLL_FROM")
var RAO_PATTERN_FILTER = regexp.MustCompile(RAO_FILTER_PATTERN)
var MIMES = []string{"application/pdf"}
var INDEX_ID = "rao"


type RaoService struct {
	searchService search.SearchService
	convService   conv.ConvService
	treeService   tree.TreeService
}

func New(searchService search.SearchService, convService conv.ConvService, treeService tree.TreeService) *RaoService {
	return &RaoService{
		treeService:   treeService,
		convService:   convService,
		searchService: searchService,
	}
}

func (service RaoService) DocFilter(doc document.IDocument) bool {
	if !utils.ArrayContainsString(MIMES, doc.GetMime()) {
		return false
	}
	log.Debug("Document path : " + doc.GetPath())
	log.Debug("Document title: " + doc.GetTitle())
	matches := RAO_PATTERN_FILTER.FindStringSubmatch(doc.GetPath())
	if nil == matches {
		return false
	}
	log.Debug("doc complies with filter assertion, processing: " + doc.GetPath() + doc.GetTitle())
	return true
}

func (service RaoService) docMapper(doc document.IDocument) map[string]interface{} {
	return map[string]interface{}{
		"Content":   doc.(rao.RaoDocument).GetContent(),
		"Title":     doc.GetTitle(),
		"Path":      doc.GetPath(),
		"Client":    doc.(rao.RaoDocument).GetClient(),
		"Agence":    doc.(rao.RaoDocument).GetAgence(),
		"Extension": doc.GetExtension(),
		"Mime":      doc.GetMime(),
		"Mtime":     doc.GetMtime(),
		"Bytes":     doc.(rao.RaoDocument).GetBytes(),
		"Sum":       doc.(rao.RaoDocument).GetSum(),
	}
}

func (service RaoService) DocHandler(doc document.IDocument) {
	bytes, size := service.treeService.GetEngine().(*dropbox2.Dropbox2).DownloadFile(doc)
	b, err := service.convService.Convert(bytes, doc.GetMime())
	log.Error(err, log.ERROR)
	content := string(b[:])
	log.Debug("[rao.service] - content : " + content)
	if "" == content {
		return // Shall we index the document if we could not extract its content ?
	}
	matches := RAO_PATTERN_FILTER.FindStringSubmatch(doc.GetPath())
	agence := matches[2]
	client := matches[3]
	chunks := utils.ChunkString(content, 10000)
	for _, chunk := range chunks {
		raoDocument := rao.RaoDocument{
			doc,
			document.BusinessDocument{
				doc,
				utils.NormalizeString(client),
				utils.NormalizeString(agence),
			},
			document.FullTextDocument{
				Bytes:   size,
				Sum:     utils.Md5Sum(content),
				Content: chunk,
			},
		}
		service.searchService.Store(INDEX_ID, raoDocument, service.docMapper)
	}
}
