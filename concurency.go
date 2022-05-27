package main

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func test1(ch1 chan int) {
// 	time.Sleep(1 * time.Second)
// 	ch1 <- 2
// }
// func test2(ch2 chan int) {
// 	time.Sleep(2 * time.Second)
// 	ch2 <- 1
// }
// func main() {
// 	// Creating channel of string type , we are passing chan as the int
// 	c1 := make(chan int)
// 	c2 := make(chan int)
// 	go test1(c1)
// 	go test2(c2)
// 	// Here we are calling two functions as the goroutine ,
// 	// where next call be made when send and receive has been done for others.
// 	select {
// 	case operation1 := <-c1:
// 		fmt.Println("The operation1 value is", operation1, "And it is the channel number")
// 	case operation2 := <-c2:
// 		fmt.Println("The operation1 value is", operation2, "And it is the channel number")
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var wg = sync.WaitGroup{}

// func main() {
// 	var myChan chan string = make(chan string, 1)

// 	wg.Add(4)
// 	go printA(myChan)
// 	go printB(myChan)
// 	go printC(myChan)
// 	go printD(myChan)
// 	time.Sleep(time.Second * 1)
// 	fmt.Println(<-myChan)
// 	fmt.Println(<-myChan)
// 	fmt.Println(<-myChan)
// 	fmt.Println(<-myChan)
// 	wg.Wait()
// }

// func printA(c chan string) {
// 	defer wg.Done()
// 	c <- "A"
// 	fmt.Println("PrintA() done")
// }

// func printB(c chan string) {
// 	defer wg.Done()
// 	c <- "B"
// 	fmt.Println("PrintB() done")
// }

// func printC(c chan string) {
// 	defer wg.Done()
// 	c <- "C"
// 	fmt.Println("PrintC() done")
// }

// func printD(c chan string) {
// 	defer wg.Done()
// 	c <- "D"
// 	fmt.Println("PrintD() done")
// }
