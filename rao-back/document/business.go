package document

type BusinessDocument struct {
	IDocument
	Client string
	Agence string
}

func (doc BusinessDocument) SetClient(c string) {
	doc.Client = c
}

func (doc BusinessDocument) GetClient() string {
	return doc.Client
}

func (doc BusinessDocument) SetAgence(a string) {
	doc.Agence = a
}

func (doc BusinessDocument) GetAgence() string {
	return doc.Agence
}