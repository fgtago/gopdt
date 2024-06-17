package main

import (
	"github.com/fgtago/fgweb"
	"github.com/fgtago/fgweb/defaulthandlers"
	"github.com/fgtago/gopdt/handlers"
	"github.com/go-chi/chi/v5"
)

func Router(mux *chi.Mux) error {

	fgweb.Get(mux, "/favicon.ico", defaulthandlers.FaviconHandler)
	fgweb.Get(mux, "/asset/*", defaulthandlers.AssetHandler)
	fgweb.Get(mux, "/template/*", defaulthandlers.TemplateHandler)
	fgweb.Get(mux, "/manifest.json", handlers.Manifest)

	fgweb.Get(mux, "/", handlers.Home)

	return nil
}
