package main

import (
    "fmt"
    "strings"
)

type PaymentProcessor interface{
    ProcessPayment(uint) string 
}
//func Process[T PaymentProcessor] (processType T,amount uint) string{
//    return processType.ProcessPayment(amount)
//}

//func ProcessPayment[T Visa | Telda | Card ](m uint,typ T) T{
//    return typ;
//}

func process(P PaymentProcessor,amount uint) string{
    return P.ProcessPayment(amount)
}

type Visa struct {
    VisaNumber string
}

func (v Visa) ProcessPayment(amount uint)string{
    return "Visa Result"
}

type Telda struct {
    TeldaNumber string
}

type Card struct {
    CardNumber string
}


type Shape interface{
    getArea() uint
}

type Triple struct {
    a uint
    b uint
    c uint
}
type Example struct {
    a uint
    b uint
}

func (t Triple) getArea() uint{
    return t.a + t.b + t.c 
}

func (e Example) getArea() uint{
    return e.a + e.b 
}



// Person represents a simple sample model.
type Person struct {
    name   string
    age    uint16
    salary float64
    fn     func(string)
}

type Product struct{
    price float64
}

func HaveStruct(p Person) Person {
    p.age = 22
    return p
}

// EfficientFunction demonstrates a value receiver method.
func (p Person) EfficientFunction(p1 Person) Person {
    p1.name = "Ali"
    return p1
}

func (p Product) CalculatePrice(price float64) float64{
    return price + (p.price* 0.4)
}

// Callable accepts a function and executes it.
func Callable(fn func(int) int) int {
    return fn(100)
}

func Double(number int) int {
    return number * 2
}

func sendData(ch chan<- string) {
    // Simulate an API request that takes time.
    ch <- "ammar"
    close(ch)
}

// StringLooping joins strings efficiently.
func StringLooping(str []string) string {
    var builder strings.Builder
    for _, val := range str {
        builder.WriteString(val)
        fmt.Println(val)
    }
    return builder.String()
}

func DisplayMap(mp map[string][]int) {
    for idx, val := range mp {
        fmt.Println("I", idx)
        fmt.Println("v", val)
    }
}

func main() {
    ch := make(chan string, 1)
    go sendData(ch)

    // Receive data from the channel after it completes.
    res := <-ch
    fmt.Println(res)

    arr := []string{"Ammar", "Ibrahim"}
    fmt.Println(StringLooping(arr))

    fmt.Println("-------------")

    mp := map[string][]int{"a": {1, 2}}
    DisplayMap(mp)

    result := Callable(Double)
    fmt.Println(result)

    person := Person{
        name:   "ammar",
        age:    20,
        salary: 12.290,
        fn: func(s string) {
            fmt.Println(s)
        },
    }
    fmt.Println("Struct ", person)
    person.fn("Hello from function field")

    price := Product{12.500}
    final_product_price := Product.CalculatePrice(price,12.500);
    fmt.Println("===================")
    fmt.Println(final_product_price);


    shapes := []Shape{Triple{a:1,b:2,c:3},Example{a:1,b: 2}}
    finalArea :=uint(0);
    for _,shape := range shapes{
        finalArea += uint(shape.getArea())
    }
    fmt.Println(finalArea); // 6 - 3 
    fmt.Println("===================")


    // Generics
    VisaResult := process(Visa{"10"},10);
    fmt.Println(VisaResult)
}
