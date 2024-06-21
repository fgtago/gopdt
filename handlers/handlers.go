package handlers

import (
	"fmt"
	"net/http"

	"github.com/agungdhewe/dwtpl"
	"github.com/fgtago/fgweb/appsmodel"
)

type Handler struct {
	Webservice *appsmodel.Webservice
}

var hdr *Handler

func New(ws *appsmodel.Webservice) *Handler {
	hdr = &Handler{
		Webservice: ws,
	}
	return hdr
}

func (hdr *Handler) LoginCheck(pv *appsmodel.PageVariable, pg *dwtpl.PageConfig) error {
	var authenticated bool

	r := pv.Request
	w := pv.Response
	ws := hdr.Webservice
	ctx := pv.Request.Context()

	authenticated = ws.Session.GetBool(ctx, string(appsmodel.IsAuthenticatedKeyName))
	if pg.Auth && !authenticated {
		url := fmt.Sprintf("/user/login?referer=%s", r.URL.Path)
		http.Redirect(w, r, url, http.StatusSeeOther)
		return nil
	}

	return nil
}
