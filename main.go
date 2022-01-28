package main

import (
	"emalm/controller"
	"emalm/model"
)

func main() {
	model.Init()
	controller.Start()
}
