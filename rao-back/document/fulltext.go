package document

// FullTextDocument adds the following fields to IDocument
//
// Content stands for a fulltext content attached to the document
//
// Sum stands for the computed hash of a fulltext content that
//
// Bytes stands for content length expressed in bytes
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
