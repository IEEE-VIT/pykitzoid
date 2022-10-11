package main

import (
   "C"
   "log"
)

//export helloWorld
func helloWorld(){
   log.Println("Hello World")
}

func main(){

}