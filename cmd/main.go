package main

import (
	"fmt"
	"github.com/jasurbek-suyunov/udevs_project/config"
	"github.com/jasurbek-suyunov/udevs_project/src/handler"
)

func main() {
	cnf := config.NewConfig()
	r := handler.SetupRouter(cnf)
	r.Run(fmt.Sprintf(":%s", cnf.HTTPPort))
}