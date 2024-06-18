package handlers

import "github.com/fgtago/fgweb/appsmodel"

type Handler struct {
	Webservice *appsmodel.Webservice
}

var hnd *Handler

func New(ws *appsmodel.Webservice) *Handler {
	hnd = &Handler{
		Webservice: ws,
	}
	return hnd
}
