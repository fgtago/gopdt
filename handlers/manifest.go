package handlers

import (
	"net/http"
	"path/filepath"
)

func (hdr *Handler) Manifest(w http.ResponseWriter, r *http.Request) {
	ws := hdr.Webservice
	path := filepath.Join(ws.RootDir, "manifest.json")
	http.ServeFile(w, r, path)
}
