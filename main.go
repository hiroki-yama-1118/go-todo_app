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

// user,_ :=models.GetUserByEmail("test@test.com")
// fmt.Println(user)

// session,err :=user.CreateSession()
// if err!=nil{
// 	log.Println(err)
// }
// fmt.Println(session)

// valid,_ := session.CheckSession()
// fmt.Println(valid)
}