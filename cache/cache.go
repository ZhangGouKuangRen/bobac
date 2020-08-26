package cache

type Cache interface {
	Set(string, []byte) error
	Get(string) ([]byte, error)
	Del(string) error
	GetStat() Stat
}

func New(cache string)Cache {
	var c Cache
	if cache == "inMemoryCache" {
		c = newInMemoryCache()
	}else {
		c = nil
	}
	return c
}
