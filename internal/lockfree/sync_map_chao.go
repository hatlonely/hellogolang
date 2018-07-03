package lockfree

// NewSyncMapChao create a new sync map
func NewSyncMapChao() *SyncMapChao {
	tokenBuf := make(chan struct{}, 1)
	tokenBuf <- struct{}{}
	rwToken := token{false, tokenBuf}
	return &SyncMapChao{
		rwToken: rwToken,
		m:       map[interface{}]interface{}{},
	}
}

// SyncMapChao sync map
type SyncMapChao struct {
	rwToken token
	m       map[interface{}]interface{}
}

type token struct {
	isReadOwner bool
	tokenBuf    chan struct{}
}

// Set key
func (m *SyncMapChao) Set(key interface{}, val interface{}) {
	if _, exist := m.m[key]; !exist {
		m.m[key] = val
	} else {
		if m.rwToken.isReadOwner {
			_ = <-m.rwToken.tokenBuf
			m.rwToken.isReadOwner = false
			m.m[key] = val
			m.rwToken.tokenBuf <- struct{}{}
		} else {
			m.m[key] = val
		}
	}
}

// Get key
func (m *SyncMapChao) Get(key interface{}) (interface{}, bool) {
	if m.rwToken.isReadOwner {
		val, ok := m.m[key]
		return val, ok
	}
	_ = <-m.rwToken.tokenBuf
	m.rwToken.isReadOwner = true
	val, ok := m.m[key]
	m.rwToken.tokenBuf <- struct{}{}
	return val, ok
}

// Del key
func (m *SyncMapChao) Del(key interface{}) {
	if _, exist := m.m[key]; !exist {
		m.m[key] = key
	} else {
		if m.rwToken.isReadOwner {
			_ = <-m.rwToken.tokenBuf
			m.rwToken.isReadOwner = false
			delete(m.m, key)
			m.rwToken.tokenBuf <- struct{}{}
		} else {
			delete(m.m, key)
		}
	}
}
