package main

import "fmt"

// Iterable defines an interface for iterable collections.
type Iterable interface {
	CreateIterator() Iterator
}

// Record defines a record structure.
type Record struct {
	Field string
	Value string
}

// Iterator defines an interface for iterating over a collection of items.
type Iterator interface {
	HasNext() bool
	GetNext() *Record
}

// RecordIterator defines an iterator for a collection of records.
type RecordIterator struct {
	records []*Record
	index   int
}

// HasNext returns true if there are more items to iterate over.
func (u *RecordIterator) HasNext() bool {
	return u.index < len(u.records)
}

// GetNext returns the next item in the collection.
func (u *RecordIterator) GetNext() *Record {
	if u.HasNext() {
		record := u.records[u.index]
		u.index++
		return record
	}

	return nil
}

// RecordCollection defines a collection of records.
type RecordCollection struct {
	records []*Record
}

// CreateIterator returns an iterator for the collection.
func (u *RecordCollection) CreateIterator() Iterator {
	return &RecordIterator{
		records: u.records,
	}
}

func main() {
	r1 := &Record{
		Field: "field1",
		Value: "value1",
	}
	r2 := &Record{
		Field: "field2",
		Value: "value2",
	}
	r3 := &Record{
		Field: "field3",
		Value: "value3",
	}
	r4 := &Record{
		Field: "field4",
		Value: "value4",
	}

	records := &RecordCollection{
		records: []*Record{r1, r2, r3, r4},
	}

	iterator := records.CreateIterator()

	for iterator.HasNext() {
		rec := iterator.GetNext()
		fmt.Printf("Record is %+v\n", rec)
	}
}
