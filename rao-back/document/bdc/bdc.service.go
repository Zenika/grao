package bdc

import (
	"regexp"

	"github.com/Zenika/RAO/document"
	"github.com/Zenika/RAO/search"
	"github.com/Zenika/RAO/tree"
	"github.com/Zenika/RAO/utils"
)

var bdcFilterPattern string = `(?i)^.+/_{1,2}clients(_|\s){1}(?P<Agence>[\w&\s]+)/(?P<Client>[^/]+)/(?P<Projet>[^/]+)/BON DE COMMANDE/(?P<Consultant>[^/]+)`
var bdcPatternFilter = regexp.MustCompile(bdcFilterPattern)

// Adding support for docx documents:
// "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
var mimes = []string{"application/pdf"}

const INDEX_ID = "bdc"

type BdcService struct {
	searchService search.SearchService
	treeService   tree.TreeService
}

// func (bdc BdcService) Poll() {
// 	root := fmt.Sprintf("/%v", os.Getenv("RAO_DBX_ROOT"))
// 	bdc.treeService.Poll(root, bdc.docFilter, bdc.docHandler)
// }

func NewService(searchService search.SearchService, treeService tree.TreeService) *BdcService {
	return &BdcService{
		treeService:   treeService,
		searchService: searchService,
	}
}

func (bdc BdcService) DocFilter(doc document.IDocument) bool {
	if !utils.ArrayContainsString(mimes, doc.GetMime()) {
		return false
	}
	matches := bdcPatternFilter.FindStringSubmatch(doc.GetPath())
	if nil == matches {
		return false
	}
	return true
}

func (bdc BdcService) docMapper(doc document.IDocument) map[string]interface{} {
	return map[string]interface{}{
		"Title":      doc.GetTitle(),
		"Path":       doc.GetPath(),
		"Client":     doc.(BdcDocument).GetClient(),
		"Agence":     doc.(BdcDocument).GetAgence(),
		"Extension":  doc.GetExtension(),
		"Mime":       doc.GetMime(),
		"Mtime":      doc.GetMtime(),
		"Consultant": doc.(BdcDocument).GetConsultant(),
		"Projet":     doc.(BdcDocument).GetProjet(),
	}
}

func (bdc BdcService) DocHandler(doc document.IDocument) {
	matches := bdcPatternFilter.FindStringSubmatch(doc.GetPath())
	agence := matches[2]
	client := matches[3]
	projet := matches[4]
	consultant := matches[5]
	bdcDocument := BdcDocument{
		doc,
		document.BusinessDocument{
			doc,
			client,
			agence,
		},
		projet,
		consultant,
	}
	bdc.searchService.Store(INDEX_ID, bdcDocument, bdc.docMapper)
}
