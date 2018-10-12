package service

import (
	"os"
	"regexp"

	"github.com/Zenika/rao/rao-back/document"
	"github.com/Zenika/rao/rao-back/document/bdc"
	"github.com/Zenika/rao/rao-back/log"
	"github.com/Zenika/rao/rao-back/search"
	"github.com/Zenika/rao/rao-back/tree"
	"github.com/Zenika/rao/rao-back/utils"
)

//var BDC_FILTER_PATTERN = `(?i)^.+/_{1,2}clients(_|\s){1}(?P<Agence>[\w&\s]+)/(?P<Client>[^/]+)/(?P<Projet>[^/]+)/BON DE COMMANDE/(?P<Consultant>[^/]+)`
var BDC_FILTER_PATTERN = os.Getenv("BDC_POLL_FROM")
var BDC_PATTERN_FILTER = regexp.MustCompile(BDC_FILTER_PATTERN)
var MIMES = []string{"application/pdf"}
var INDEX_ID = "bdc"
var REFERER  = log.GetReferer()

type BdcService struct {
	searchService search.SearchService
	treeService   tree.TreeService
}

func New(searchService search.SearchService, treeService tree.TreeService) *BdcService {
	return &BdcService{
		treeService:   treeService,
		searchService: searchService,
	}
}

func (service BdcService) DocFilter(doc document.IDocument) bool {
	if !utils.ArrayContainsString(MIMES, doc.GetMime()) {
		log.Debug("bad mime " + doc.GetMime(), REFERER)
		return false
	}
	matches := BDC_PATTERN_FILTER.FindStringSubmatch(doc.GetPath())
	if nil == matches {
		log.Debug("regexp filter: " + BDC_PATTERN_FILTER.String(), REFERER)
		log.Debug("no match " + doc.GetTitle(), REFERER)
		return false
	}
	log.Debug("doc complies with filter assertion, processing: " + doc.GetPath(), REFERER)
	return true
}

func (service BdcService) docMapper(doc document.IDocument) map[string]interface{} {
	return map[string]interface{}{
		"Title":      doc.GetTitle(),
		"Path":       doc.GetPath(),
		"Client":     doc.(bdc.BdcDocument).GetClient(),
		"Agence":     doc.(bdc.BdcDocument).GetAgence(),
		"Extension":  doc.GetExtension(),
		"Mime":       doc.GetMime(),
		"Mtime":      doc.GetMtime(),
		"Consultant": doc.(bdc.BdcDocument).GetConsultant(),
		"Projet":     doc.(bdc.BdcDocument).GetProjet(),
	}
}

func (service BdcService) DocHandler(doc document.IDocument) {
	matches := BDC_PATTERN_FILTER.FindStringSubmatch(doc.GetPath())
	agence := matches[2]
	client := matches[3]
	projet := matches[4]
	consultant := matches[5]
	bdcDocument := bdc.BdcDocument{
		doc,
		document.BusinessDocument{
			doc,
			utils.NormalizeString(client),
			utils.NormalizeString(agence),
		},
		projet,
		utils.NormalizeString(consultant),
	}
	service.searchService.Store(INDEX_ID, bdcDocument, service.docMapper)
}
