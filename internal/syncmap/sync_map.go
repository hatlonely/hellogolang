package syncmap

import "sync"

// NewSyncMap create a new sync map
func NewSyncMap(shards int, hashFunc HashFunc) *SyncMap {
	maps := make([]map[interface{}]interface{}, shards)
	for i := range maps {
		maps[i] = map[interface{}]interface{}{}
	}

	return &SyncMap{
		maps:     maps,
		mutexs:   make([]sync.RWMutex, shards),
		shards:   shards,
		hashFunc: hashFunc,
	}
}

// SyncMap sync map
type SyncMap struct {
	shards   int
	maps     []map[interface{}]interface{}
	mutexs   []sync.RWMutex
	hashFunc HashFunc
}

// Set key
func (m *SyncMap) Set(key interface{}, val interface{}) {
	shard := m.hashFunc(key) % m.shards
	m.mutexs[shard].Lock()
	m.maps[shard][key] = val
	m.mutexs[shard].Unlock()
}

// Get key
func (m *SyncMap) Get(key interface{}) (interface{}, bool) {
	shard := m.hashFunc(key) % m.shards
	m.mutexs[shard].RLock()
	val, ok := m.maps[shard][key]
	m.mutexs[shard].RUnlock()
	return val, ok
}

// Del key
func (m *SyncMap) Del(key interface{}) {
	shard := m.hashFunc(key) % m.shards
	m.mutexs[shard].Lock()
	delete(m.maps[shard], key)
	m.mutexs[shard].Unlock()
}
