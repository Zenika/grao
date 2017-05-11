# document
--
    import "."

Package document contains common document interfaces, with subpackages related
to document implementations and their associated services

## Usage

#### type BusinessDocument

```go
type BusinessDocument struct {
	IDocument
	Client string
	Agence string
}
```

BusinessDocument adds the following fields to IDocument


Client stands for a customer

Agence stands for a business area division

#### func (BusinessDocument) GetAgence

```go
func (doc BusinessDocument) GetAgence() string
```

#### func (BusinessDocument) GetClient

```go
func (doc BusinessDocument) GetClient() string
```

#### func (BusinessDocument) SetAgence

```go
func (doc BusinessDocument) SetAgence(a string)
```

#### func (BusinessDocument) SetClient

```go
func (doc BusinessDocument) SetClient(c string)
```

#### type Document

```go
type Document struct {
	Title     string
	Path      string
	Mime      string
	Extension string
	Mtime     time.Time
}
```


#### func (Document) GetExtension

```go
func (doc Document) GetExtension() string
```

#### func (Document) GetMime

```go
func (doc Document) GetMime() string
```

#### func (Document) GetMtime

```go
func (doc Document) GetMtime() time.Time
```

#### func (Document) GetPath

```go
func (doc Document) GetPath() string
```

#### func (Document) GetTitle

```go
func (doc Document) GetTitle() string
```

#### func (Document) SetExtension

```go
func (doc Document) SetExtension(e string)
```

#### func (Document) SetMime

```go
func (doc Document) SetMime(m string)
```

#### func (Document) SetMtime

```go
func (doc Document) SetMtime(t time.Time)
```

#### func (Document) SetPath

```go
func (doc Document) SetPath(p string)
```

#### func (Document) SetTitle

```go
func (doc Document) SetTitle(t string)
```

#### type DocumentMapper

```go
type DocumentMapper func(doc IDocument) map[string]interface{}
```

DocumentMapper is a mapper function that can be used by external services to
convert any implentation of IDocument into a map

#### type FullTextDocument

```go
type FullTextDocument struct {
	IDocument
	Content string
	Sum     string
	Bytes   int64
}
```

FullTextDocument adds the following fields to IDocument


Content stands for a fulltext content attached to the document

Sum stands for the computed hash of a fulltext content that

Bytes stands for content length expressed in bytes

#### func (FullTextDocument) GetBytes

```go
func (doc FullTextDocument) GetBytes() int64
```

#### func (FullTextDocument) GetContent

```go
func (doc FullTextDocument) GetContent() string
```

#### func (FullTextDocument) GetSum

```go
func (doc FullTextDocument) GetSum() string
```

#### func (FullTextDocument) SetBytes

```go
func (doc FullTextDocument) SetBytes(b int64)
```

#### func (FullTextDocument) SetContent

```go
func (doc FullTextDocument) SetContent(c string)
```

#### func (FullTextDocument) SetSum

```go
func (doc FullTextDocument) SetSum(s string)
```

#### type IDocument

```go
type IDocument interface {
	GetTitle() string
	GetPath() string
	GetMime() string
	GetExtension() string
	GetMtime() time.Time
	SetTitle(t string)
	SetPath(p string)
	SetMime(m string)
	SetExtension(e string)
	SetMtime(t time.Time)
}
```

IDocument is the common inherited interface for documents
# tree
--
    import "."

Tree Package contains tree service interfaces with subpackages related to their
implementations.

A tree service compose a TreeEngine interface implementation, provided as an
argument to the factory call

## Usage

#### type TreeEngine

```go
type TreeEngine interface {
	Poll(root string, pairs [][]interface{})
	LongPoll(root string, pairs [][]interface{})
}
```

TreeEngine implementation own the responsability of implementing tree service
core methods

Both **Poll** and **LongPoll** methods take a root path as their first argument
and an array of function pair as their second argument:

*pairs[0]* is of type ```go func(IDocument)(bool)``` and acts as a filter

*pairs[1]* is of type func(IDocument) and is called only if *pairs[0]* evaluates
to true

#### type TreeService

```go
type TreeService struct {
}
```


#### func  New

```go
func New(eng TreeEngine) *TreeService
```

#### func (TreeService) GetEngine

```go
func (tree TreeService) GetEngine() TreeEngine
```

#### func (TreeService) LongPoll

```go
func (tree TreeService) LongPoll(root string, pairs [][]interface{})
```

#### func (TreeService) Poll

```go
func (tree TreeService) Poll(root string, pairs [][]interface{})
```
# conv
--
    import "."

conv Package contains conv service interfaces with subpackages related to their
implementations.

A conv service compose a ConvEngine interface implementation, provided as an
argument to the factory call

## Usage

#### type ConvEngine

```go
type ConvEngine interface {
	Convert(input []byte, mimeType string) ([]byte, error)
}
```

ConvEngine implementation own the responsability of implementing conv service
core method

**Convert** takes a binary content as an input and convert its content as a
readable fulltext stream using mimetype to guess which conversion strategy to
use

#### type ConvService

```go
type ConvService struct {
}
```


#### func  New

```go
func New(eng ConvEngine) *ConvService
```

#### func (ConvService) Convert

```go
func (conv ConvService) Convert(input []byte, mimeType string) ([]byte, error)
```
# search
--
    import "."

Search Package contains search service interfaces with subpackages related to
their implementations.

A search service compose a SearchEngine interface implementation, provided as an
argument to the factory call

## Usage

#### type Query

```go
type Query struct {
	Query        string
	Facets       string
	FacetFilters string
	Filters      string
	Page         int
	HitsPerPage  int
	Restriction  string
}
```


#### type Response

```go
type Response struct {
	Data interface{}
}
```


#### type SearchEngine

```go
type SearchEngine interface {
	Store(index string, doc document.IDocument, docMapper document.DocumentMapper)
	Search(index string, query Query) (*Response, error)
	Configure(index string, settings map[string]interface{}) error
}
```

SearchEngine implementation own the responsability of implementing search
service core methods

**Store** stores the provided document to an index referenced by its first
argument. docMapper function may be used to convert document to a
map[string]interface{} complying with the underlying implementation

**Search** will perform on an indexed referenced by its first argument a query
provided as a seconde argument under the form of a Query object

**Configure** should be used to tune index before performing queries if needed

#### type SearchService

```go
type SearchService struct {
}
```


#### func  New

```go
func New(eng SearchEngine) *SearchService
```

#### func (SearchService) Configure

```go
func (search SearchService) Configure(index string, settings map[string]interface{}) error
```

#### func (SearchService) Search

```go
func (search SearchService) Search(index string, query Query) (*Response, error)
```

#### func (SearchService) Store

```go
func (search SearchService) Store(index string, doc document.IDocument, docMapper document.DocumentMapper)
```
