package handlers

import (
	"fmt"
	"net/http"

	"github.com/fgtago/fgweb/appsmodel"
	"github.com/fgtago/fgweb/defaulthandlers"
)

func (hdr *Handler) Account(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pv := ctx.Value(appsmodel.PageVariableKeyName).(*appsmodel.PageVariable)
	pv.PageName = "account"
	pv.Request = r
	pv.Response = w
	pv.Use(hdr.LoginCheck)

	fmt.Println("get username:", pv.UserNickName)

	defaulthandlers.SimplePageHandler(pv, w, r)

}
