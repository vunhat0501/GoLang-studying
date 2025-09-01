package main

import "fmt"

func main(){
	var username string = "john"
	var password string = "123456"
	
	fmt.Println("Authorization: Basic", username + ":" + password)
}