package handlers

import (
	"net/http"
	"path/filepath"
)

func (hnd *Handler) Manifest(w http.ResponseWriter, r *http.Request) {
	ws := hnd.Webservice
	path := filepath.Join(ws.RootDir, "manifest.json")
	http.ServeFile(w, r, path)
}
