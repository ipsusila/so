package factory

import (
	"strings"
	"sync"
)

type Factory interface {
	New(name string) interface{}
}

var (
	mu        sync.RWMutex
	factories = make(map[string]Factory)
)

func Register(pkgName string, f Factory) {
	mu.Lock()
	defer mu.Unlock()

	if f == nil {
		panic("Factory is nil")
	}
	if _, exist := factories[pkgName]; exist {
		panic("Factory already registered")
	}

	factories[pkgName] = f
}

func New(typeName string) interface{} {
	items := strings.Split(typeName, ".")
	if len(items) < 2 {
		return nil
	}

	mu.RLock()
	defer mu.RUnlock()
	if f, exist := factories[items[0]]; exist {
		return f.New(items[1])
	}
	return nil
}
