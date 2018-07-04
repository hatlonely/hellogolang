package syncmap

import "sync"

type StdSyncMap struct {
	m sync.Map
}

func (m *StdSyncMap) Set(key interface{}, val interface{}) {
	m.m.Store(key, val)
}

func (m *StdSyncMap) Get(key interface{}) (interface{}, bool) {
	return m.m.Load(key)
}

func (m *StdSyncMap) Del(key interface{}) {
	m.m.Delete(key)
}
