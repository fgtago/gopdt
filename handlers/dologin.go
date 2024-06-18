package handlers

import "net/http"

func (hdr *Handler) DoLogin(w http.ResponseWriter, r *http.Request) {
	hdr.Webservice.Session.RenewToken(r.Context())

	w.Write([]byte("login di post"))
}
