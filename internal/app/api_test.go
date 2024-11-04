package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler_GetHandler(t *testing.T) {
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	type calls struct {
		getCalls int
	}
	tests := []struct {
		name           string
		h              *Handler
		args           args
		want           string
		wantStatusCode int
		calls          calls
	}{
		{
			name: "When given a correct key, 200 status returned",
			h: &Handler{
				Get: func(key string) ([]string, bool) {
					return []string{"value1"}, true
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/key1", nil),
			},
			want:           "[\"value1\"]",
			wantStatusCode: 200,
			calls: calls{
				getCalls: 1,
			},
		},
		{
			name: "When given an empty key all keys are returned",
			h: &Handler{
				Get: func(key string) ([]string, bool) {
					return []string{"key1", "key2"}, true
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/", nil),
			},
			want:           "[\"key1\",\"key2\"]",
			wantStatusCode: 200,
			calls: calls{
				getCalls: 1,
			},
		},
		{
			name: "when given an invalid key, 404 response sent",
			h: &Handler{
				Get: func(key string) ([]string, bool) {
					return nil, false
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/invalidKey", nil),
			},
			want:           "key not found\n",
			wantStatusCode: 404,
			calls: calls{
				getCalls: 1,
			},
		},
	}
	for _, tt := range tests {
		calls := calls{}
		h := &Handler{
			Get: func(key string) ([]string, bool) {
				calls.getCalls++
				return tt.h.Get(key)
			},
		}
		t.Run(tt.name, func(t *testing.T) {
			h.GetHandler(tt.args.w, tt.args.r)
			resp := tt.args.w
			if resp.Result().StatusCode != tt.wantStatusCode {
				t.Errorf("h.GetHandlerStatusCode= got %v, want %v", resp.Result().StatusCode, tt.want)
			}
			if resp.Body.String() != tt.want {
				t.Errorf("h.GetHandler= got %v, want %v", resp.Body.String(), tt.want)
			}
			assert.Equal(t, calls, tt.calls)
		})
	}
}

func TestHandler_PutHandler(t *testing.T) {
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	type calls struct {
		putCalls int
	}
	tests := []struct {
		name           string
		args           args
		want           string
		wantStatusCode int
		calls          calls
	}{
		{
			name: "When given a put request, it store successfully",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer([]byte(`"value1"`))),
			},
			want:           "request successful, information added to data store",
			wantStatusCode: 200,
			calls: calls{
				putCalls: 1,
			},
		},
		{
			name: "When given an invalid json, 400 status code is returned",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer([]byte(`{}"value1"}`))),
			}, want: "invalid request body\n",
			wantStatusCode: 400,
			calls: calls{
				putCalls: 0,
			},
		},
	}
	for _, tt := range tests {
		calls := calls{}
		h := &Handler{
			Put: func(key, value string) {
				calls.putCalls++
			},
		}
		t.Run(tt.name, func(t *testing.T) {
			h.PutHandler(tt.args.w, tt.args.r)
			resp := tt.args.w
			if resp.Result().StatusCode != tt.wantStatusCode {
				t.Errorf("h.PutHandlerStatusCode= got %v, want %v", resp.Result().StatusCode, tt.want)
			}
			if resp.Body.String() != tt.want {
				t.Errorf("h.PutHandler= got %v, want %v", resp.Body.String(), tt.want)
			}
			assert.Equal(t, calls, tt.calls)
		})
	}
}

func TestHandler_DeleteHandler(t *testing.T) {
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	type calls struct {
		deleteCalls int
	}
	tests := []struct {
		name           string
		h              *Handler
		args           args
		want           string
		wantStatusCode int
		calls          calls
	}{
		{
			name: "When given a valid key, deletes successfully and returns 200 status code",
			h: &Handler{
				Delete: func(key string) bool {
					return true
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodDelete, "/key1", nil),
			},
			want:           " removed successfully from store",
			wantStatusCode: 200,
			calls: calls{
				deleteCalls: 1,
			},
		},
		{
			name: "given an invalidKey, 404 status code returned",
			h: &Handler{
				Delete: func(key string) bool {
					return false
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodDelete, "/key1", nil),
			},
			want:           " not found in store\n",
			wantStatusCode: 404,
			calls: calls{
				deleteCalls: 1,
			},
		},
	}
	for _, tt := range tests {
		calls := calls{}
		h := &Handler{
			Delete: func(key string) bool {
				calls.deleteCalls++
				return tt.h.Delete(key)
			},
		}
		t.Run(tt.name, func(t *testing.T) {
			h.DeleteHandler(tt.args.w, tt.args.r)
			resp := tt.args.w
			if resp.Result().StatusCode != tt.wantStatusCode {
				t.Errorf("h.DeleteStatusCode= got %v, want %v", resp.Result().StatusCode, tt.want)
			}
			if resp.Body.String() != tt.want {
				t.Errorf("h.DeleteHandler= got %v, want %v", resp.Body.String(), tt.want)
			}
			assert.Equal(t, calls, tt.calls)
		})
	}
}
