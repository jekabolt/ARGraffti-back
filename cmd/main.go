package main

import (
	"github.com/jekabolt/ARGraffti-back"
	"github.com/jekabolt/config"
	"github.com/jekabolt/slf"
)

var (
	log = slf.WithContext("main")

	branch    string
	commit    string
	buildtime string
	lasttag   string
	// TODO: add all default params
	globalOpt = graffity.Configuration{
		// Name: "my-test-back",
		// Database: store.Conf{
		// 	Address:             "localhost:27017",
		// 	DBUsers:             "userDB-test",
		// 	DBFeeRates:          "BTCMempool-test",
		// 	DBTx:                "DBTx-test",
		// 	DBStockExchangeRate: "dev-DBStockExchangeRate",
		// 	Username:            "Username",
		// 	Password:            "Password",
		// },
		RestAddress: "localhost:7778",
	}
)

func main() {
	config.ReadGlobalConfig(&globalOpt, "multy configuration")

	log.Error("--------------------------------new multy back server session")
	log.Infof("CONFIGURATION=%+v", globalOpt)

	log.Infof("branch: %s", branch)
	log.Infof("commit: %s", commit)
	log.Infof("build time: %s", buildtime)
	log.Infof("tag: %s", lasttag)

	mu, err := graffity.Init(&globalOpt)
	if err != nil {
		log.Fatalf("Server initialization: %s\n", err.Error())
	}

	if err = mu.Run(); err != nil {
		log.Fatalf("Server running: %s\n", err.Error())
	}

}
