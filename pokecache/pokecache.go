package cache

import (
	"time"
	"sync"
)

type cacheEntry struct {
	createdAt time.Time 
	val []byte
}

type Cache struct {
	entry map[string]cacheEntry
	mux sync.Mutex
}
var interval = time.Second * 7

func (c Cache) Add(key string, val []byte) {
	newEntry := cacheEntry{time.Now(), val}
	c.entry[key] = newEntry
}
// get elem, ok = Cache[key]
func (c Cache) Get(key string) ([]byte, bool){
	entry, ok := c.entry[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (c Cache) reapLoop() {
	for _, elem := range c.entry {
		if time.Now().Sub(elem.createdAt) > interval {
			//add a delete here 
			// ?delete(c.entry, c.entry[elem])?
		}
	}
}
func NewCache(createdAt int, val []byte) {

}