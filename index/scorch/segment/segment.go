package segment

import (
	"github.com/RoaringBitmap/roaring"
	"github.com/blevesearch/bleve/index"
)

// DocumentFieldValueVisitor defines a callback to be visited for each
// stored field value.  The return value determines if the visitor
// should keep going.  Returning true continues visiting, false stops.
type DocumentFieldValueVisitor func(field string, typ byte, value []byte, pos []uint64) bool

type Segment interface {
	Dictionary(field string) TermDictionary

	VisitDocument(num uint64, visitor DocumentFieldValueVisitor) error
	Count() uint64

	DocNumbers([]string) *roaring.Bitmap

	Fields() []string
}

type TermDictionary interface {
	PostingsList(term string, except *roaring.Bitmap) PostingsList

	Iterator() DictionaryIterator
	PrefixIterator(prefix string) DictionaryIterator
	RangeIterator(start, end string) DictionaryIterator
}

type DictionaryIterator interface {
	Next() (*index.DictEntry, error)
}

type PostingsList interface {
	Iterator() PostingsIterator

	Count() uint64

	// NOTE deferred for future work

	// And(other PostingsList) PostingsList
	// Or(other PostingsList) PostingsList
}

type PostingsIterator interface {
	Next() Posting
}

type Posting interface {
	Number() uint64

	Frequency() uint64
	Norm() float64

	Locations() []Location
}

type Location interface {
	Field() string
	Start() uint64
	End() uint64
	Pos() uint64
	ArrayPositions() []uint64
}