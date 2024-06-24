package handlers

import (
	"fmt"
	"net/http"

	"github.com/fgtago/fgweb/appsmodel"
	"github.com/fgtago/fgweb/defaulthandlers"
)

type PricingPageData struct {
}

func (hdr *Handler) Pricing(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pv := ctx.Value(appsmodel.PageVariableKeyName).(*appsmodel.PageVariable)
	pv.PageName = "pricing"
	pv.Form = appsmodel.NewForm(nil)
	pv.Use(hdr.LoginCheck)

	defaulthandlers.SimplePageHandler(pv, w, r)

}

func (hdr *Handler) PricingPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pv := ctx.Value(appsmodel.PageVariableKeyName).(*appsmodel.PageVariable)
	pv.PageName = "pricing"
	pv.Form = appsmodel.NewForm(nil)
	pv.Use(hdr.LoginCheck)

	fmt.Fprintf(w, "ini data yang di post")

}
