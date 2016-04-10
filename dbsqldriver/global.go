package shuntdbsqldriver


import (
	"sync"
)


var (
	// To make it so `globalMap` can be modified from multiple goroutines, we make
	// it "thread safe" by guarding it with the read write mutex `globalMapMutex`.
	globalMapMutex sync.RWMutex

	// globalMap is used to store the shunted iterators.
	//
	// In general they should not be in here for a very long time.
	globalMap map[string]Iterator = make(map[string]Iterator)
)


func deleteIterator(name string) {

	globalMapMutex.Lock()
	delete(globalMap, name)
	globalMapMutex.Unlock()
}


func getIterator(name string) Iterator {

	globalMapMutex.RLock()
	iterator, ok := globalMap[name]
	globalMapMutex.RUnlock()
	if !ok {
		return nil
	}

	return iterator
}


func insertIterator(name string, iterator Iterator) error {

	if nil == iterator {
		return nil
	}

	globalMapMutex.Lock()
	_, ok := globalMap[name]
	if ok {
		globalMapMutex.Unlock()
		return errAlreadyExists
	}
	globalMap[name] = iterator
	globalMapMutex.Unlock()

	return nil
}
