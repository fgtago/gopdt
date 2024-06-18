package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/alexedwards/scs/v2"
	"github.com/fgtago/fgweb"
	"github.com/go-chi/chi/v5"
)

func main() {
	var err error

	fmt.Println("Starting Program ...")

	// baca parameter dari cli
	var cfgFileName string
	flag.StringVar(&cfgFileName, "conf", "config.yml", "nama file konfigurasi yang akan di load")
	flag.Parse()

	// set root direktori ke current working direktori
	rootDir, err := os.Getwd()
	if err != nil {
		// ada error saat inisiasi webservice, halt
		panic(err.Error())
	}

	// ambil file konfigurasi
	cfgpath := filepath.Join(rootDir, cfgFileName)

	// start jalankan web
	ws, err := fgweb.New(rootDir, cfgpath)
	if err != nil {
		// ada error saat inisiasi webservice, halt
		panic(err.Error())
	}

	// create session manager
	sessmgr := scs.New()
	fmt.Println(sessmgr)
	ws.SessMgr = sessmgr

	// info: memulai service
	port := ws.Configuration.Port
	fmt.Println("service running on port:", port)
	err = fgweb.StartService(port, func(mux *chi.Mux) error {
		return Router(mux)
	})
	if err != nil {
		// ada error saat service start, halt
		panic(err.Error())
	}
}
