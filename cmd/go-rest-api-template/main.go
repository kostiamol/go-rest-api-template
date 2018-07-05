package main

import (
	"log"
	"os"

	"github.com/kostiamol/go-rest-api-template/storage"
	"github.com/kostiamol/go-rest-api-template/svc"

	"github.com/unrolled/render"
)

func main() {
	var (
		env      = os.Getenv("ENV")      // LOCAL, DEV, STAGE, PROD
		port     = os.Getenv("PORT")     // server traffic on this port
		version  = os.Getenv("VERSION")  // path to VERSION file
		fixtures = os.Getenv("FIXTURES") // path to fixtures file
	)
	if env == "" || env == svc.Local {
		env = svc.Local
		port = "3001"
		version = "../../VERSION"
		fixtures = "../../fixtures.json"
	} else if env == svc.Prod {
		env = svc.Prod
		port = "8080"
		version = "./rsc/VERSION"
		fixtures = "./rsc/fixtures.json"
	}
	version, err := svc.ParseVersionFile(version)
	if err != nil {
		log.Fatal(err)
	}
	db, err := storage.LoadFixturesIntoMockDB(fixtures)
	if err != nil {
		log.Fatal(err)
	}
	ctx := svc.Context{
		Render:  render.New(),
		Version: version,
		Env:     env,
		Port:    port,
		DB:      db,
	}
	svc.Run(ctx)
}
