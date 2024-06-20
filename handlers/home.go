package handlers

import (
	"net/http"

	"github.com/fgtago/fgweb/appsmodel"
	"github.com/fgtago/fgweb/defaulthandlers"
)

func (hdr *Handler) Home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pv := ctx.Value(appsmodel.PageVariableKeyName).(*appsmodel.PageVariable)
	pv.PageName = "home"
	pv.Use(hdr.LoginCheck)
	defaulthandlers.SimplePageHandler(pv, w, r)
}
