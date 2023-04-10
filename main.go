package main

import (
	"fmt"
	"log"
	"todo_app/app/controllers"
	"todo_app/app/models"
)

func main() {
fmt.Println(models.Db)

err := controllers.StartMainServer()
if err != nil {
		log.Fatalln(err)
}

}