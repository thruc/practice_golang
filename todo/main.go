package main

import (
	"fmt"

	"todo/app/controllers"
	"todo/app/models"
)

func TestConnection() {

}

func main() {
	fmt.Println(models.Db)
	go controllers.StartMainServer()

	for {

	}
}
