package config

import (
	"flag"
	"log"
)

var flagDB bool

func init() {
	flag.BoolVar(&flagDB, "db", false, "create table")
	flag.Parse()

	createTokenAuth()

	if err := getEnv(); err != nil {
		log.Println("Error get env: ", err)
	}

	if err := connectDatabase(flagDB); err != nil {
		log.Println("Error connect database: ", err)
	} else {
		log.Println("Conected!")
	}

}
