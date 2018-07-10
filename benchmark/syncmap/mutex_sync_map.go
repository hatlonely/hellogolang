package syncmap

import "sync"

func NewMutexSyncMap(shards int, hashFunc HashFunc) *MutexSyncMap {
	maps := make([]map[interface{}]interface{}, shards)
	for i := range maps {
		maps[i] = map[interface{}]interface{}{}
	}

	return &MutexSyncMap{
		maps:     maps,
		mutexs:   make([]sync.RWMutex, shards),
		shards:   shards,
		hashFunc: hashFunc,
	}
}

type HashFunc func(interface{}) int

type MutexSyncMap struct {
	shards   int
	maps     []map[interface{}]interface{}
	mutexs   []sync.RWMutex
	hashFunc HashFunc
}

func (m *MutexSyncMap) Set(key interface{}, val interface{}) {
	shard := m.hashFunc(key) % m.shards
	m.mutexs[shard].Lock()
	m.maps[shard][key] = val
	m.mutexs[shard].Unlock()
}

func (m *MutexSyncMap) Get(key interface{}) (interface{}, bool) {
	shard := m.hashFunc(key) % m.shards
	m.mutexs[shard].RLock()
	val, ok := m.maps[shard][key]
	m.mutexs[shard].RUnlock()
	return val, ok
}

func (m *MutexSyncMap) Del(key interface{}) {
	shard := m.hashFunc(key) % m.shards
	m.mutexs[shard].Lock()
	delete(m.maps[shard], key)
	m.mutexs[shard].Unlock()
}
