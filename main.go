package main

import (
	`label-backend/config`
	`label-backend/lib/db`
	`label-backend/router`
)

func main(){
	config.InitConf()
	db.InitDB()
	router.Init()
}