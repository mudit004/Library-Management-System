package main

import (
	"fmt"
	"lms/pkg/api"
	"lms/pkg/models"
)

func main(){
	_, err := models.Connection()
	fmt.Println((err))
	api.Start()

}

