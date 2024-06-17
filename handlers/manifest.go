package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/fgtago/fgweb/appsmodel"
)

func Manifest(w http.ResponseWriter, r *http.Request) {
	ws := appsmodel.GetWebservice()
	path := filepath.Join(ws.RootDir, "manifest.json")
	http.ServeFile(w, r, path)
}
