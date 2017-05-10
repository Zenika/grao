// conv Package contains conv service interfaces
// with subpackages related to their implementations
//
// A conv service compose a ConvEngine interface implementation,
// provided as an argument to the factory call
package conv

// ConvEngine implementation own the responsability of
// implementing conv service core method
//
// **Convert** takes a binary content as an input
// and convert its content as a readable fulltext
// stream using mimetype to guess which conversion
// strategy to use
type ConvEngine interface {
	Convert(input []byte, mimeType string) ([]byte, error)
}

type ConvService struct {
	engine ConvEngine
}

func New(eng ConvEngine) *ConvService {
	return &ConvService{
		engine: eng,
	}
}

func (conv ConvService) Convert(input []byte, mimeType string) ([]byte, error) {
	return conv.engine.Convert(input, mimeType)
}
