package store

import (
	"reflect"
	"testing"
)

func TestStore_Get(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name  string
		s     *Store
		args  args
		want  []string
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Get(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store.Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Store.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestStore_Put(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name string
		s    *Store
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Put(tt.args.key, tt.args.value)
		})
	}
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
