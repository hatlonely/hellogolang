package syncmap

import (
	"fmt"
	"strings"
)

// NewLockfreeMap create a new map
func NewLockfreeMap(bucketSize int, hashFunc HashFunc) *LockfreeMap {
	wops := make(chan *WriteOperation, 100)

	lm := &LockfreeMap{
		buckets:  make([]Entry, bucketSize),
		capacity: bucketSize,
		hash:     hashFunc,
		wops:     wops,
	}

	go func() {
		for op := range wops {
			if op.op == OPSet {
				lm.SyncSet(op.key, op.val)
			} else if op.op == OPDel {
				lm.SyncDel(op.key)
			}
		}
	}()

	return lm
}

// HashFunc hash function
type HashFunc func(interface{}) int

// LockfreeMap lockfree map
type LockfreeMap struct {
	buckets  []Entry
	capacity int
	hash     HashFunc
	wops     chan *WriteOperation
}

// OperationName set or del
type OperationName int

// OperationNames
const (
	OPSet OperationName = 1
	OPDel OperationName = 2
)

// WriteOperation write operation
type WriteOperation struct {
	op  OperationName
	key interface{}
	val interface{}
}

// Entry entry for key value
type Entry struct {
	key  interface{}
	val  interface{}
	next *Entry
}

// Get key
func (m *LockfreeMap) Get(key interface{}) (interface{}, bool) {
	h := m.hash(key) % m.capacity
	e := m.buckets[h].getEntry(key)
	if e != nil {
		return e.val, true
	}

	return nil, false
}

// Set key val
func (m *LockfreeMap) Set(key interface{}, val interface{}) {
	m.wops <- &WriteOperation{
		op:  OPSet,
		key: key,
		val: val,
	}
}

// Del key
func (m *LockfreeMap) Del(key interface{}) {
	m.wops <- &WriteOperation{
		op:  OPDel,
		key: key,
	}
}

// SyncSet sync set
func (m *LockfreeMap) SyncSet(key interface{}, val interface{}) {
	h := m.hash(key) % m.capacity
	m.buckets[h].setEntry(key, val)
}

// SyncDel sync del
func (m *LockfreeMap) SyncDel(key interface{}) {
	h := m.hash(key) % m.capacity
	m.buckets[h].delEntry(key)
}

func (m *LockfreeMap) show() {
	for i, e := range m.buckets {
		fmt.Printf("[%v]: %v\n", i, strings.Join(e.show(), ", "))
	}
}

func (e *Entry) show() []string {
	var l []string
	e = e.next
	for e != nil {
		l = append(l, fmt.Sprintf("(%v, %v)", e.key, e.val))
		e = e.next
	}

	return l
}

func (e *Entry) setEntry(key interface{}, val interface{}) {
	for e.next != nil {
		if e.next.key == key {
			e.next.val = val
			return
		}
		e = e.next
	}
	e.next = &Entry{
		key:  key,
		val:  val,
		next: nil,
	}
}

func (e *Entry) getEntry(key interface{}) *Entry {
	e = e.next
	for e != nil {
		if e.key == key {
			return e
		}
		e = e.next
	}

	return nil
}

func (e *Entry) delEntry(key interface{}) {
	for e.next != nil {
		if e.next.key == key {
			e.next = e.next.next
			return
		}
		e = e.next
	}
}
