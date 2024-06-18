package handlers

import (
	"fmt"
	"net/http"

	"github.com/fgtago/fgweb/appsmodel"
	"github.com/fgtago/fgweb/defaulthandlers"
)

func Account(w http.ResponseWriter, r *http.Request) {
	ws := appsmodel.GetWebservice()
	ctx := r.Context()
	pv := ctx.Value(appsmodel.PageVariableKeyName).(*appsmodel.PageVariable)
	pv.PageName = "account"

	username := ws.Session.GetString(r.Context(), "username")
	fmt.Println("get username:", username)

	defaulthandlers.SimplePageHandler(pv, w, r)

}
