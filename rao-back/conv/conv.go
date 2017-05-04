package conv

type ConvEngine interface {
	Convert(input []byte, mimeType string) ([]byte, error)
}

type ConvService struct {
	engine ConvEngine
}

func (conv ConvService) Convert(input []byte, mimeType string) ([]byte, error) {
	return conv.engine.Convert(input, mimeType)
}

func New(eng ConvEngine) *ConvService {
	return &ConvService{
		engine: eng,
	}
}
