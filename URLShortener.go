package main

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"math/rand"
	"net/http"
	"net/url"
	"time"
	"strings"
)

var db *bolt.DB
var bucketName = []byte("url_mapping")

type URLRequest struct {
	LongURL string `json:"long_url"`
}

func initDB() error {
	var err error
	db, err = bolt.Open("url_shortener.db", 0600, nil)
	if err != nil {
		return err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketName)
		return err
	})

	return err
}

func generateShortKey() string {
	rand.Seed(time.Now().UnixNano())
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	shortKey := make([]byte, 6)
	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortKey)
}

func isValidURL(longURL string) bool {
	_, err := url.ParseRequestURI(longURL)
	return err == nil
}

func shortenURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	var urlReq URLRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&urlReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if !isValidURL(urlReq.LongURL) {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	shortKey := generateShortKey()

	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		return bucket.Put([]byte(shortKey), []byte(urlReq.LongURL))
	})

	if err != nil {
		http.Error(w, "Failed to save URL", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"short_key": shortKey})
}

func redirectToURLHandler(w http.ResponseWriter, r *http.Request) {
	shortKey := strings.TrimPrefix(r.URL.Path, "/")

	var longURL string
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		longURLBytes := bucket.Get([]byte(shortKey))
		if longURLBytes == nil {
			return fmt.Errorf("Short key not found")
		}
		longURL = string(longURLBytes)
		return nil
	})

	if err != nil {
		http.Error(w, "Short key not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}

func main() {
	if err := initDB(); err != nil {
		fmt.Println("Error initializing database:", err)
		return
	}
	defer db.Close()

	http.HandleFunc("/shorten", shortenURLHandler)           
	http.HandleFunc("/", redirectToURLHandler)               

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
