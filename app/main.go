package main

import (
	"github.com/sammervalgas/api-go-gin-clubs/datasources"
	"github.com/sammervalgas/api-go-gin-clubs/routes"
)

func main() {
	datasources.DbConnection()
	routes.HandleRequest()
}
