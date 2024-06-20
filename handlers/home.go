package handlers

import (
	"fmt"
	"net/http"

	"github.com/fgtago/fgweb/appsmodel"
	"github.com/fgtago/fgweb/defaulthandlers"
)

func (hdr *Handler) Home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pv := ctx.Value(appsmodel.PageVariableKeyName).(*appsmodel.PageVariable)
	pv.PageName = "home"
	pv.Request = r
	pv.Response = w
	pv.Use(hdr.LoginCheck)

	fmt.Println("username:", pv.UserNickName)

	defaulthandlers.SimplePageHandler(pv, w, r)
}
