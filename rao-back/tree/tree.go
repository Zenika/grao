package tree

import (
	"github.com/Zenika/RAO/document"
)

type TreeEngine interface {
	Poll(root string, filter document.DocumentFilter, handler document.DocumentHandler)
	LongPoll(root string, filter document.DocumentFilter, handler document.DocumentHandler)
}

type TreeService struct {
	engine TreeEngine
}

func (tree TreeService) Poll(root string, filter document.DocumentFilter, handler document.DocumentHandler) {
	tree.engine.Poll(root, filter, handler)
}

func (tree TreeService) LongPoll(root string, filter document.DocumentFilter, handler document.DocumentHandler) {
	tree.engine.LongPoll(root, filter, handler)
}

func (tree TreeService) GetEngine() TreeEngine {
	return tree.engine
}

func New(eng TreeEngine) *TreeService {
	return &TreeService{
		engine: eng,
	}
}
