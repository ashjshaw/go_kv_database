package api

import (
	"encoding/json"
	"net/http"

	store "github.com/ashjshaw/go_kv_database/internal/pkg"
	"github.com/gorilla/mux"
)

type Handler struct {
	Get    func(key string) ([]string, bool)
	Put    func(key, value string)
	Delete func(key string) bool
}

func New() *Handler {
	store := store.NewStore()
	newHandler := &Handler{}
	newHandler.Get = store.Get
	newHandler.Put = store.Put
	newHandler.Delete = store.Delete
	return newHandler
}

func (h *Handler) GetHandler(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	value, exists := h.Get(key)
	if !exists {
		http.Error(w, "key not found", http.StatusNotFound)
		return
	}
	jsonData, err := json.Marshal(value)
	if err != nil {
		http.Error(w, "error encoding json", http.StatusInternalServerError)
		return
	}
	w.WriteHeader((http.StatusOK))
	w.Write(jsonData)
}

func (h *Handler) PutHandler(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	value := ""
	if err := json.NewDecoder(r.Body).Decode(&value); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	h.Put(key, value)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("request successful, information added to data store"))
}

func (h *Handler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	if exists := h.Delete(key); !exists {
		errorString := key + " not found in store"
		http.Error(w, errorString, http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	responseString := key + " removed successfully from store"
	w.Write([]byte(responseString))

}
