package tree

import (
	"github.com/Zenika/RAO/document"
)

type TreeEngine interface {
  Walk(root string, handler document.DocumentHandler)
  Poll(root string, handler document.DocumentHandler)
}

type TreeService struct {
	engine TreeEngine
}

func (tree TreeService) Walk(root string, handler document.DocumentHandler) {
	tree.engine.Walk(root, handler)
}

func (tree TreeService) Poll(root string, handler document.DocumentHandler) {
	tree.engine.Poll(root, handler)
}

func New(eng TreeEngine) *TreeService {
	return &TreeService{
		engine: eng,
	}
}
