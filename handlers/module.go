package handlers

import (
	"net/http"

	"github.com/fgtago/fgweb/appsmodel"
	"github.com/fgtago/fgweb/defaulthandlers"
	"github.com/go-chi/chi/v5"
)

func (hdr *Handler) Module(w http.ResponseWriter, r *http.Request) {
	modulename := chi.URLParam(r, "modulename")

	ctx := r.Context()
	pv := ctx.Value(appsmodel.PageVariableKeyName).(*appsmodel.PageVariable)
	pv.PageName = modulename
	pv.Form = appsmodel.NewForm(nil)
	pv.Use(hdr.LoginCheck)

	defaulthandlers.SimplePageHandler(pv, w, r)
}
