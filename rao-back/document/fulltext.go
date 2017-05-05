package document

type FullTextDocument struct {
	IDocument
	Content string
	Sum     string
	Bytes   int64
}

func (doc FullTextDocument) GetContent() string {
	return doc.Content
}

func (doc FullTextDocument) SetContent(c string) {
	doc.Content = c
}

func (doc FullTextDocument) GetSum() string {
	return doc.Sum
}

func (doc FullTextDocument) SetSum(s string) {
	doc.Sum = s
}

func (doc FullTextDocument) GetBytes() int64 {
	return doc.Bytes
}

func (doc FullTextDocument) SetBytes(b int64) {
	doc.Bytes = b
}
