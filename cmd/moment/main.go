package main

import (
	"moment/pkg"
	"moment/pkg/db"
	"moment/pkg/web"

	"gitlab.com/knopkalab/go/logger"
	"gitlab.com/knopkalab/go/utils"
)

func main() {
	conf, _ := pkg.LoadConfig()
	log := logger.New(conf.Log, nil)
	if err := conf.Save(); err != nil {
		log.Err(err).Msg("Error save config file")
	}
	db, err := db.Open(conf.DBReindexer, log)
	check(err, log, "error connect database")
	defer db.Close()
	check(db.Migrate(), log, "error run migration")

	controllers := web.NewController(log, db)

	srv, err := web.StartServer(conf, log, controllers)
	check(err, log, "error connect server")

	defer srv.Stop()
	utils.WaitForExit()

}

func check(err error, log logger.Logger, errMessage string) {
	if err != nil {
		log.Fatal().Err(err).Msg(errMessage)
	}
}
