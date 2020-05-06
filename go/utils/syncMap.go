package utils

import "sync"


type SyncMap struct {
	data map[string]string
	*sync.RWMutex
}

func NewSyncMap(data map[string]string) *SyncMap {
	return &SyncMap{data, &sync.RWMutex{}}
}

func (d *SyncMap) Len() int {
	d.RLock()
	defer d.RUnlock()
	return len(d.data)
}

func (d *SyncMap) Put(key string, value string) (string, bool) {
	d.Lock()
	defer d.Unlock()
	oldValue, ok := d.data[key]
	d.data[key] = value
	return oldValue, ok
}

func (d *SyncMap) Get(key string) (string, bool) {
	d.RLock()
	defer d.RUnlock()
	oldValue, ok := d.data[key]
	return oldValue, ok
}

func (d *SyncMap) Delete(key string) (string, bool) {
	d.Lock()
	defer d.Unlock()
	oldValue, ok := d.data[key]
	if ok {
		delete(d.data, key)
	}
	return oldValue, ok
}
