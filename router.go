package main

import (
	"github.com/fgtago/fgweb"
	"github.com/fgtago/fgweb/appsmodel"
	"github.com/fgtago/fgweb/defaulthandlers"
	"github.com/fgtago/gopdt/handlers"
	"github.com/go-chi/chi/v5"
)

func Router(mux *chi.Mux) error {

	fgweb.Get(mux, "/favicon.ico", defaulthandlers.FaviconHandler)
	fgweb.Get(mux, "/asset/*", defaulthandlers.AssetHandler)
	fgweb.Get(mux, "/template/*", defaulthandlers.TemplateHandler)

	hnd := handlers.New(appsmodel.GetWebservice())

	fgweb.Get(mux, "/", hnd.Home)
	fgweb.Get(mux, "/manifest.json", hnd.Manifest)
	fgweb.Get(mux, "/module/{modulename}", hnd.Module)
	fgweb.Get(mux, "/account", hnd.Account)

	return nil
}
