package main

import (
	"jacq/config"
	"jacq/serverutils"
)

var c config.Config

func main() {
	c = config.ImportConfig(config.OSSource{})
	store, client := serverutils.SetUpDatabase(c.DatabaseURL, c.DatabaseName)
	handler := serverutils.SetUpHandler(store)
	server := serverutils.SetUpServer(&handler)
	router := serverutils.SetupRouter(&server)
	serverutils.SetupSwagger(router)
	serverutils.StartServer(router, client)
}
