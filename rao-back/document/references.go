package document

// ReferencesDocument adds the following fields to IDocument
//
// Tags stands for a list of relevant tags that relate to the reference 

type ReferencesDocument struct {
	IDocument
	Client string
	Tag string
}

func (doc ReferencesDocument) SetClient(c string) {
	doc.Client = c
}

func (doc ReferencesDocument) GetClient() string {
	return doc.Client
}

func (doc ReferencesDocument) GetTag() string {
	return doc.Tag
}

func (doc ReferencesDocument) SetTag(tag string){
	doc.Tag = tag
}