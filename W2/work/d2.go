package work

import (
	"fmt"
)

func ChannelDemo() {
	oneC := make(chan int)

	go func(){
		oneC <- 32
	}()
	v := <- oneC

	fmt.Println(v)
}

func SelectDemo(){
	oneC := make(chan int)
	go func ()  {
		for i := 0; i < 10; i++{
			oneC <- i
		}	
		defer close(oneC)
	}()
	// for {
	// 	select {
	// 	case valC, ok := <-oneC:
	// 		if ok {
	// 			fmt.Println("valC: ", valC)
	// 		} else {
	// 			fmt.Println("end")
	// 			return
	// 		}
			
	// 	}
	// }
	for valC := range oneC {
		fmt.Println("valC: ", valC)
	}
	fmt.Println("end")
}

var myC chan int

func Producer(){
	for i := 0; i < 20; i++ {
		myC <- i
		fmt.Println("produce", i)
	}
	close(myC)
}

func Consumer(){
	for valC := range myC {
		fmt.Println("consume", valC)
	}
}

func RunDemo(){
	myC = make(chan int, 2)
	go Producer()
	Consumer()
}