package handlers

import (
	"net/http"

	"github.com/fgtago/fgweb/appsmodel"
	"github.com/fgtago/fgweb/defaulthandlers"
)

func Home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pv := ctx.Value(appsmodel.PageVariableKeyName).(*appsmodel.PageVariable)
	pv.PageName = "home"

	defaulthandlers.SimplePageHandler(pv, w, r)
}
