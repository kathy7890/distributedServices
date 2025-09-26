package server

import (
	"fmt"
	"sync"
)

type Log struct {
	mu sync.Mutex
	records []Record
}

func NewLog() *Log {
	return &Log{}
}

// START:types
type Record struct {
	Value  []byte `json:"value"`  // Struct tags are used to provide metadata about a struct field about how to serialize and deserialize. 
	Offset uint64 `json:"offset"` 
}

//END:types

func (c *Log) Append(record Record) (uint64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	record.Offset = uint64(len(c.records))
	c.records = append(c.records, record)
	return record.Offset,nil
}


func (c *Log) Read(offset uint64) (Record, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if offset >= uint64(len(c.records)) {
		return Record{}, ErrOffsetNotFound
	}
	return c.records[offset], nil
}


var ErrOffsetNotFound = fmt.Errorf("offset not found")




