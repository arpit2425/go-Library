package main

import (
	"log"

	"golang.org/x/text/message"
)
func LoggerInfo(message string){
	log.Println(message)
}
func LoggerError(message string){
	log.Printf("This is an error %v",message)
}