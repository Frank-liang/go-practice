package urlshort

import (
	"encoding/json"
	"net/http"

	"github.com/boltdb/bolt"
	"gopkg.in/yaml.v2"
)

func MapHandler(pathToUrls map[string]string, fallback http.Handle) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if path, ok := pathToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, path, http.StatusFound)
		}
		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yamldata []byte, fallback http.Handle) (http.HandlerFunc, error) {
	var pathToUrls []struct {
		Path string `yaml:"path"`
		URL  string `yaml:"url"`
	}
	if err := yaml.Unmarshal(yamldata, &pathToUrls); err != nil {
		return nil, err
	}
	return func(w http.ResponseWriter, r *http.Request) {
		for _, pathtourl := range pathToUrls {
			if pathtourl.Path == r.URL.Path {
				http.Redirect(w, r, pathtourl.URL, http.StatusFound)
				return
			}
		}
		fallback.ServeHTTP(w, r)
	}, nil
}

func JSONHandler(jsondata []byte, fallback http.Handle) (http.HandlerFunc, error) {
	var pathToUrls []struct {
		Path string `json:"path"`
		URL  string `json:"url"`
	}
	if err := json.Unmarshal(jsondata, &pathToUrls); err != nil {
		return nil, err
	}
	return func(w http.ResponseWriter, r *http.Request) {
		for _, pathtourl := range pathToUrls {
			if pathtourl.Path == r.URL.Path {
				http.Redirect(w, r.pathtourl.URL, http.StatusFound)
				return
			}
		}
		fallback.ServeHTTP(w, r)
	}, nil
}

func BOLTHandler(db *bolt.DB, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := db.View(func(tx *bolt.Tx) error {
			bucket := tx.Bucket([]byte("pathstourls"))
			if bucket != nil {
				cursor := bucket.Cursor()
				for path, url := cursor.First(); path != nil; path, url = cursor.Next() {
					if string(path) == r.URL.Path {
						http.Redirect(w, r, string(url), http.StatusFound)
						return nil
					}
				}
			}
			return nil
		}); err != nil {
			panic(err)
		}
		fallback.ServeHTTP(w, r)
	}
}
