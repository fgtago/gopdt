package handlers

import (
	"net/http"

	"github.com/agungdhewe/dwlog"
	"github.com/fgtago/fgweb/appsmodel"
	"github.com/fgtago/fgweb/defaulthandlers"
)

type LoginData struct {
	Email    string
	Password string
	Remember bool
}

func (hdr *Handler) Login(w http.ResponseWriter, r *http.Request) {
	ws := hdr.Webservice
	ctx := r.Context()
	pv := ctx.Value(appsmodel.PageVariableKeyName).(*appsmodel.PageVariable)

	var pagename = "login"
	if r.Method == "POST" {
		hdr.Webservice.Session.RenewToken(r.Context())
		err := r.ParseForm()
		if err != nil {
			if ws.ShowServerError {
				defaulthandlers.ErrorPageHandler(500, err.Error(), pv, w, r)
			} else {
				dwlog.Error(err.Error())
				defaulthandlers.ErrorPageHandler(500, "internal server error", pv, w, r)
			}
			return
		}

		data := LoginData{
			Email:    r.Form.Get("email"),
			Password: r.Form.Get("password"),
			Remember: r.Form.Get("rememberme") == "on",
		}

		form := appsmodel.NewForm(r.PostForm)
		form.Requires("email", "password")
		if !form.Valid() {
			pv.Form = form
			pv.Data = &data
			pv.PageName = pagename
			defaulthandlers.SimplePageHandler(pv, w, r)
			return
		}

		var authenticated bool
		authenticated = false
		if data.Email == "agung" && data.Password == "rahasia" {
			authenticated = true

			// set session
			ws.Session.RememberMe(r.Context(), data.Remember)
			ws.Session.Put(r.Context(), string(appsmodel.IsAuthenticatedKeyName), true)
			ws.Session.Put(r.Context(), string(appsmodel.UserIdKeyName), "xxxx")
			ws.Session.Put(r.Context(), string(appsmodel.UserNickNameKeyName), "agung")
			ws.Session.Put(r.Context(), string(appsmodel.UserFullNameKeyName), "Agung Nugroho")

		}

		if authenticated {
			// login berhasil, redirect ke halaman yang di refer sebelumnya
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		} else {
			form.Errors.Add("login", "email atau password salah")
			pv.Form = form
			pv.Data = &data
			pv.PageName = pagename
			defaulthandlers.SimplePageHandler(pv, w, r)
			return
		}

	} else {

		// jika sudah login, redirect ke halaman home
		authenticated := ws.Session.GetBool(ctx, string(appsmodel.IsAuthenticatedKeyName))
		if authenticated {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		pv.Data = &LoginData{}
		pv.Form = appsmodel.NewForm(nil)
		pv.PageName = pagename
		defaulthandlers.SimplePageHandler(pv, w, r)
	}

}
