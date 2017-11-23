package service

import (
	"os"
	"regexp"

	"github.com/Zenika/RAO/conv"
	"github.com/Zenika/RAO/document"
	"github.com/Zenika/RAO/document/refs"
	"github.com/Zenika/RAO/log"
	"github.com/Zenika/RAO/search"
	"github.com/Zenika/RAO/tree"
	"github.com/Zenika/RAO/tree/dropbox"
	"github.com/Zenika/RAO/utils"
)

var RAO_FILTER_PATTERN string = os.Getenv("RAO_POLL_FROM")
var RAO_PATTERN_FILTER = regexp.MustCompile(RAO_FILTER_PATTERN)
var MIMES = []string{"application/pdf"}
var INDEX_ID = "refs"

type RefsService struct {
	searchService search.SearchService
	convService   conv.ConvService
	treeService   tree.TreeService
}

func New(searchService search.SearchService, convService conv.ConvService, treeService tree.TreeService) *RefsService {
	return &RefsService{
		treeService:   treeService,
		convService:   convService,
		searchService: searchService,
	}
}

func (service RefsService) DocFilter(doc document.IDocument) bool {
	if !utils.ArrayContainsString(MIMES, doc.GetMime()) {
		return false
	}
	matches := RAO_PATTERN_FILTER.FindStringSubmatch(doc.GetPath())
	if nil == matches {
		return false
	}
	log.Debug("doc complies with filter assertion, processing: " + doc.GetPath())
	return true
}

func (service RefsService) docMapper(doc document.IDocument) map[string]interface{} {
	return map[string]interface{}{
		"Tag":      doc.(refs.RefsDocument).GetTag(),
		"Path":      doc.GetPath(),
		"Client":    doc.(refs.RefsDocument).GetClient(),
		"Extension": doc.GetExtension(),
		"Mime":      doc.GetMime(),
		"Mtime":     doc.GetMtime(),
		"Bytes":     doc.(refs.RefsDocument).GetBytes(),
	}
}

func (service RefsService) DocHandler(doc document.IDocument) {
	bytes, size := service.treeService.GetEngine().(*dropbox.Dropbox).DownloadFile(doc)
	b, err := service.convService.Convert(bytes, doc.GetMime())
	log.Error(err, log.ERROR)
	content := string(b[:])
	if "" == content {
		return // Shall we index the document if we could not extract its content ?
	}
	matches := RAO_PATTERN_FILTER.FindStringSubmatch(doc.GetPath())
	client := matches[0]
	tag := matches[1]
	
	chunks := utils.ChunkString(content, 10000)
	for _, chunk := range chunks {
		refsDocument := refs.RefsDocument{
			doc,
			document.ReferencesDocument{
				doc,
				utils.NormalizeString(client),
				tag,
			},
			document.FullTextDocument{
				Bytes:   size,
				Sum:     utils.Md5Sum(content),
				Content: chunk,
			},
		}
		service.searchService.Store(INDEX_ID, refsDocument, service.docMapper)
	}
}
