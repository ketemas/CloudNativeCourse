package lru

import "errors"

type Cacher interface {
	Get(interface{}) (interface{}, error)
	Put(interface{}, interface{}) error
}

type lruCache struct {
	size      int
	remaining int
	cache     map[string]string
	queue     []string
}

func NewCache(size int) Cacher {
	return &lruCache{size: size, remaining: size, cache: make(map[string]string), queue: make([]string, size)}
}

func (lru *lruCache) Get(key interface{}) (interface{}, error) {
	// Your code here....	
	element, ok := lru.cache[key.(string)]
	if ok{
		return key, errors.New("element is in queue")
	}
	return key, errors.New("element not in queue")
	temp 
	for element := range lru.cache{
		
	}
}

func (lru *lruCache) Put(key, val interface{}) error {
	// Your code here....
	value, ok := lru.cache[key.(string)]
	if ok{
		return key, errors.New("element is in queue")
	}
	return key, errors.New("element not in queue")
}

// Delete element from queue
func (lru *lruCache) qDel(ele string) {
	for i := 0; i < len(lru.queue); i++ {
		if lru.queue[i] == ele {
			oldlen := len(lru.queue)
			copy(lru.queue[i:], lru.queue[i+1:])
			lru.queue = lru.queue[:oldlen-1]
			break
		}
	}
}



/*
	for element := range lru.cache{
		if key.(string) := lru.cache[element]{
			lru.cache.MoveBefore(element, lru.cache.Front())
			return lru.cache()
		} 
	}
	
	

	ele, has:=lru.cache[key.(string)]
	if !has{
		return lru.cache, errors.New("element not in queue")
	}
	lru.cache.MoveBefore(ele, lru.cache.Front())
	return ele.value.([]string), error
*/
