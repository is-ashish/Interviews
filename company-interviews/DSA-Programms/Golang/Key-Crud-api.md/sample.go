package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/twinj/uuid"
)

type Key struct {
	ID      string    `json:"id"`
	Expires time.Time `json:"expires"`
}

// createKeyRequest is the optional request body for CreateKey.
type createKeyRequest struct {
	Expires *time.Time `json:"expires"`
}

var (
	store   = make(map[string]Key)
	storeMu sync.RWMutex
)

func main_() {

	router := mux.NewRouter()

	sub := router.PathPrefix("/keys").Subrouter()

	sub.HandleFunc("", GetKeys).
		Methods("GET")

	sub.HandleFunc("/{key_id}", GetKey).
		Methods("GET")

	sub.HandleFunc("", CreateKey).
		Methods("POST").
		HeadersRegexp("Content-Type", "application/json")

	sub.HandleFunc("/{key_id}", DeleteKey).
		Methods("DELETE")

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusNotFound)

	})

	srv := &http.Server{
		Handler:      router,
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Server is running at http://%s", srv.Addr)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}

func CreateKey(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	// parse optional request body

	var request createKeyRequest

	if r.Body != nil && r.ContentLength != 0 {
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
		}
	}
	// determine expiration time
	expiration := time.Now().UTC().Add(time.Hour)
	if request.Expires != nil {
		expiration = request.Expires.UTC()
	}

	key := Key{
		ID:      uuid.NewV4().String(),
		Expires: expiration,
	}

	storeMu.Lock()
	store[key.ID] = key
	storeMu.Unlock()

	marshalled, err := json.Marshal(key)

	if err != nil {
		http.Error(w, `{"error":"failed to marshal key"}`, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(marshalled); err != nil {
		log.Printf("CreateKey: failed to write response: %v", err)
	}

	log.Printf("CreateKey: created key %s expiring at %s", key.ID, key.Expires)

}

// GetKey returns a single key by ID. Returns 404 if not found or expired.
// GET /keys/<uuid>

func GetKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract the ID from the URL path: /keys/<uuid>
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, `{"error":"missing key id"}`, http.StatusBadRequest)
		return
	}

	storeMu.RLock()
	key, ok := store[id]
	storeMu.RUnlock()

	if !ok || key.Expires.Before(time.Now().UTC()) {
		http.Error(w, `{"error":"key not found"}`, http.StatusNotFound)
		return
	}

	marshalled, err := json.Marshal(key)
	if err != nil {
		http.Error(w, `{"error":"failed to marshal key"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(marshalled); err != nil {
		log.Printf("GetKey: failed to write response: %v", err)
	}

	log.Printf("GetKey: returned key %s", key.ID)
}

// GetKeys returns all unexpired keys as a JSON array.
// GET /keys
func GetKeys(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	now := time.Now().UTC()

	storeMu.RLock()
	keys := make([]Key, 0, len(store))
	for _, k := range store {
		if k.Expires.After(now) {
			keys = append(keys, k)
		}
	}
	storeMu.RUnlock()

	marshalled, err := json.Marshal(keys)
	if err != nil {
		http.Error(w, `{"error":"failed to marshal keys"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(marshalled); err != nil {
		log.Printf("GetKeys: failed to write response: %v", err)
	}

	log.Printf("GetKeys: returned %d unexpired keys", len(keys))
}

// DeleteKey deletes a key by ID. Returns 404 if not found or expired, 204 on success.
// DELETE /keys/<uuid>
func DeleteKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")
	if id == "" {
		http.Error(w, `{"error":"missing key id"}`, http.StatusBadRequest)
		return
	}

	storeMu.Lock()
	key, ok := store[id]
	if ok && !key.Expires.Before(time.Now().UTC()) {
		delete(store, id)
	}
	storeMu.Unlock()

	if !ok || key.Expires.Before(time.Now().UTC()) {
		http.Error(w, `{"error":"key not found"}`, http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	log.Printf("DeleteKey: deleted key %s", id)
}
