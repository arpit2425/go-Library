package mylogger

import (
	"log"
)
func LoggerInfo(message string){
	log.Println(message)
}
func LoggerError(message string){
	log.Printf("This is an error %v",message)
}