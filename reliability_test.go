package kind

import (
	"reflect"
	"sync"
	"testing"
)

func TestConcurrentCacheAccess(t *testing.T) {
	types := []reflect.Type{
		reflect.TypeFor[int](),
		reflect.TypeFor[[]int](),
		reflect.TypeFor[map[string][]int](),
		reflect.TypeFor[chan *int](),
		reflect.TypeFor[func(int) string](),
	}

	var wg sync.WaitGroup
	for range 64 {
		for _, typ := range types {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for range 256 {
					k := OfType(typ)
					if k.Type() != typ {
						t.Errorf("Type() = %v, want %v", k.Type(), typ)
						return
					}
				}
			}()
		}
	}
	wg.Wait()
}
