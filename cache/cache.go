package cache

import "sync"

/*
PaladinCache exported variable
*/

/*
Cache exported structure
*/
type Cache struct {
	data   map[string][]byte
	sets   int64
	gets   int64
	hits   int64
	dels   int64
	misses int64
	mutex  sync.RWMutex
}

/*
Stats structure to descriibg cache statistic
*/
type Stats struct {
	Sets   int64 `json:"sets"`
	Gets   int64 `json:"gets"`
	Hits   int64 `json:"hits"`
	Misses int64 `json:"misses"`
	Dels   int64 `json:"deletes"`
}

// PaladinCache main cache object
var PaladinCache Cache

/*
Set updates a key:value pair in the cache.
*/
func (v *Cache) Set(key string, value []byte) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	v.sets++
	if v.data == nil {
		v.data = map[string][]byte{}
	}
	v.data[key] = value
	return nil
}

// Get returns the value stored at `key`.
func (v *Cache) Get(key string) []byte {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	v.gets++

	value, ok := v.data[key]
	if ok {
		v.hits++
		return value
	}
	v.misses++
	return nil
}

// Remove removes the provided key from the cache
func (v *Cache) Remove(key string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	_, ok := v.data[key]
	if ok {
		v.dels++
		delete(v.data, key)
	}
	return nil
}

// Has returns whether or not the `key` is in the cache without updating
func (v *Cache) Has(key string) bool {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	_, ok := v.data[key]
	return ok
}

//Stats return cache statistic
func (v *Cache) Stats() Stats {
	v.mutex.RLock()
	defer v.mutex.RUnlock()

	return Stats{
		Sets:   v.sets,
		Gets:   v.gets,
		Hits:   v.hits,
		Dels:   v.dels,
		Misses: v.misses,
	}
}
