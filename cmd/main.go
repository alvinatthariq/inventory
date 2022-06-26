package main

import (
	restserver "inventory/rest"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	restserver.Init(router)
}
