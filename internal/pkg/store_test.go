package store

import (
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStore__Put_and_Get(t *testing.T) {
	testStore := &Store{
		data: map[string]string{},
		mu:   sync.RWMutex{},
	}

	var wg sync.WaitGroup

	t.Run("testing 100 Concurrent Put requests", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				n := strconv.Itoa(i)
				testStore.Put("key"+n, "value"+n)
			}(i)
		}
		wg.Wait()
		assert.Equal(t, len(testStore.data), 100)
	})
}

func TestStore_Delete(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		s    *Store
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Delete(tt.args.key); got != tt.want {
				t.Errorf("Store.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
