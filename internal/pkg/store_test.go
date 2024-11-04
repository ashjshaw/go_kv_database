package store

import (
	"reflect"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
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

	t.Run("testing 100 Concurrent Get Requests", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				n := strconv.Itoa(i)
				got, boolResponse := testStore.Get("key" + n)
				if !reflect.DeepEqual(got, []string{"value" + n}) {
					t.Errorf("Store.Get(), got = %v, want = %v", got, []string{"value" + n})
				}
				if !boolResponse {
					t.Errorf("Store.Get(), got = %v, want = %v", boolResponse, true)
				}
			}(i)
			wg.Wait()
		}

		t.Run("testing for false response from invalidKey", func(t *testing.T) {
			testFalse, boolResponse := testStore.Get("invalidKey")
			assert.Nil(t, testFalse)
			if boolResponse {
				t.Errorf("Store.Get(), got = %v, want = %v", boolResponse, false)
			}
		})
		t.Run("testing for all keys when empty string given to get", func(t *testing.T) {
			testAll, _ := testStore.Get("")
			assert.Equal(t, len(testAll), 100)
			assert.True(t, strings.Contains(testAll[0], "key"))
		})
	})

	t.Run("testing 100 Concurrent Delete Requests", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				n := strconv.Itoa(i)
				boolResponse := testStore.Delete("key" + n)
				if !boolResponse {
					t.Errorf("Store.Get(), got = %v, want = %v", boolResponse, true)
				}
			}(i)
		}
		wg.Wait()
		assert.Equal(t, len(testStore.data), 0)
	})

	t.Run("When given an invalidKey, delete returns false", func(t *testing.T) {
		assert.False(t, testStore.Delete("invalidKey"))
	})
}
