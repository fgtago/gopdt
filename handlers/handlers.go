package handlers

import "github.com/fgtago/fgweb/appsmodel"

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
