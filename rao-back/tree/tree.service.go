// Tree Package contains tree service interfaces
// with subpackages related to their implementations.
//
// A tree service compose a TreeEngine interface implementation,
// provided as an argument to the factory call
package tree

import "github.com/Zenika/rao/rao-back/log"

var REFERER = log.GetReferer()
// TreeEngine implementation own the responsability of
// implementing tree service core methods
//
// Both **Poll** and **LongPoll** methods take a root path as their
// first argument and an array of function pair as their second
// argument:
//
// *pairs[0]* is of type ```go func(IDocument)(bool)``` and acts as a filter
//
// *pairs[1]* is of type func(IDocument) and is called only
// if *pairs[0]* evaluates to true
type TreeEngine interface {
	Poll(root string, pairs [][]interface{})
	LongPoll(root string, pairs [][]interface{})
}

type TreeService struct {
	engine TreeEngine
}

func New(eng TreeEngine) *TreeService {
	return &TreeService{
		engine: eng,
	}
}

func (tree TreeService) Poll(root string, pairs [][]interface{}) {
	log.Debug("root :" +root, REFERER)
	tree.engine.Poll(root, pairs)
}

func (tree TreeService) LongPoll(root string, pairs [][]interface{}) {
	tree.engine.LongPoll(root, pairs)
}

func (tree TreeService) GetEngine() TreeEngine {
	return tree.engine
}
