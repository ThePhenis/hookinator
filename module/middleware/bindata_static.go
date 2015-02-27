package middleware

import (
	"mime"
	"net/http"
	"path/filepath"
)

type StaticLoader func(name string) ([]byte, error)

type Static struct {
	Loader StaticLoader
}

func NewStaticLoader(sl StaticLoader) *Static {
	return &Static{sl}
}

func (static *Static) ServeHTTP(rw http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	if r.Method != "GET" {
		n(rw, r)
		return
	}

	asset := r.URL.Path
	asset = asset[1:]

	data, err := static.Loader(asset)
	if err != nil {
		n(rw, r)
		return
	}

	ext := filepath.Ext(asset)

	var mt string

	if len(ext) > 0 {
		mt = mime.TypeByExtension(ext)
	}

	if 0 == len(mt) {
		mt = "application/octet-stream"
	}

	rw.Header().Set("Content-Type", mt)
	rw.Write(data)
}
