package syncmap

type Map interface {
	Set(key interface{}, val interface{})
	Get(key interface{}) (interface{}, bool)
	Del(key interface{})
}
