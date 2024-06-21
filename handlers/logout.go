package handlers

import (
	"net/http"

	"github.com/fgtago/fgweb/appsmodel"
	"github.com/fgtago/fgweb/defaulthandlers"
)

func (hdr *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pv := ctx.Value(appsmodel.PageVariableKeyName).(*appsmodel.PageVariable)
	pv.PageName = "logout"
	pv.Use(hdr.LoginCheck)

	pv.IsAuthenticated = false
	pv.UserId = ""
	pv.UserNickName = ""
	pv.UserFullName = ""

	// logout dari sistem
	ws := hdr.Webservice
	ws.Session.Remove(r.Context(), string(appsmodel.IsAuthenticatedKeyName))
	ws.Session.Remove(r.Context(), string(appsmodel.UserIdKeyName))
	ws.Session.Remove(r.Context(), string(appsmodel.UserNickNameKeyName))
	ws.Session.Remove(r.Context(), string(appsmodel.UserFullNameKeyName))
	ws.Session.RenewToken(r.Context())

	defaulthandlers.SimplePageHandler(pv, w, r)
}
