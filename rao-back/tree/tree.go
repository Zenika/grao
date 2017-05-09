package tree

type TreeEngine interface {
	Poll(root string, pairs [][]interface{})
	LongPoll(root string, pairs [][]interface{})
}

type TreeService struct {
	engine TreeEngine
}

func (tree TreeService) Poll(root string, pairs [][]interface{}) {
	tree.engine.Poll(root, pairs)
}

func (tree TreeService) LongPoll(root string, pairs [][]interface{}) {
	tree.engine.LongPoll(root, pairs)
}

func (tree TreeService) GetEngine() TreeEngine {
	return tree.engine
}

func New(eng TreeEngine) *TreeService {
	return &TreeService{
		engine: eng,
	}
}
