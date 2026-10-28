package main

import "fmt"
//import "strconv"


// Function that take function 
func Callable(fn func(int) int) int{
	return fn(100);
}

func Double(number int) int{
	return (number*2);
}

func sendData(ch chan<-string){

	// Simulate api request that take time 
	ch <- "ammar"
}

// String Looping
func StringLooping(str []string)string {
	var result string;
		for _,val := range str{
			result+=val;
			fmt.Println(val);
		}
		return result;
}


func DisplayMap(mp map[string][]int){
	for  idx,val := range mp{
		fmt.Println("I",idx);
		fmt.Println("v",val);
	}
}
func main(){
	ch :=make(chan string)
	go sendData(ch);
// recive data from channel 

	// Simulate api request to get the data from channel after finish 
	res := <- ch
	fmt.Println(res)
	arr:=[]string{"Ammar","Ibrahim"}
	StringLooping(arr[:])

	fmt.Println("-------------")
	// Declare Map
	mp := map[string][]int{"a":{1,2}};

	DisplayMap(mp)

	// Callable function
	result :=Callable(Double);
	fmt.Println(result);
}