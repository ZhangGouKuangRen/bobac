package cache

import "sync"

type inMemoryCache struct {
	container map[string][]byte
	mutex     sync.RWMutex
	Stat
}

func newInMemoryCache() *inMemoryCache {
	return &inMemoryCache{
		container: make(map[string][]byte),
		mutex: sync.RWMutex{},
		Stat: Stat{},
	}
}

func (imc *inMemoryCache)Set(key string, value []byte) error {
	imc.mutex.Lock()
	defer imc.mutex.Unlock()
	//更新Stat状态
	_, exist := imc.container[key]
	if exist {
		imc.del(key, value)
	}
	imc.container[key]=value
	imc.add(key, value)
	return nil
}

func (imc *inMemoryCache)Get(key string)([]byte, error)  {
	imc.mutex.RLock()
	defer imc.mutex.RUnlock()
	return imc.container[key], nil
}

func (imc *inMemoryCache)Del(key string)error  {
    imc.mutex.Lock()
    defer imc.mutex.Unlock()
    value, exist := imc.container[key]
    //如果要删除的键值对存在，更新Stat状态
    if exist {
    	imc.del(key, value)
		delete(imc.container, key)
	}
	return nil
}

func (imc *inMemoryCache)GetStat() Stat {
    return imc.Stat
}
