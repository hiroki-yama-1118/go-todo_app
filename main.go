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
// fmt.Println("User*",user)

// session,err :=user.CreateSession()
// if err!=nil{
// 	log.Println("Error*",err)
// }
// fmt.Println("Session*",session)

// valid,_ := session.CheckSession()
// fmt.Println(valid)
}