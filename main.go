package main

import "fmt"
//import "strconv"

func sendData(ch chan<-string){

	// Simulate api request that take time 
	ch <- "ammar"
}

func main(){
	ch :=make(chan string)
	go sendData(ch);
// recive data from channel 

	// Simulate api request to get the data from channel after finish 
	res := <- ch
	fmt.Println(res)

}