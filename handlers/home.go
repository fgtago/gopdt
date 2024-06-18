package handlers

import (
	"fmt"
	"net/http"

	"github.com/fgtago/fgweb/appsmodel"
	"github.com/fgtago/fgweb/defaulthandlers"
)

func Home(w http.ResponseWriter, r *http.Request) {
	ws := appsmodel.GetWebservice()
	ctx := r.Context()
	pv := ctx.Value(appsmodel.PageVariableKeyName).(*appsmodel.PageVariable)
	pv.PageName = "home"

	// coba test session
	fmt.Println("set username")
	ws.Session.Put(r.Context(), "username", "Agung")

	defaulthandlers.SimplePageHandler(pv, w, r)
}
