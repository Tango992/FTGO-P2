package main

import (
	"fmt"
	// "sync"
	"time"
)


/*
	TASK 1
*/
func main() {
	go numbers()
	go alphabet()
	time.Sleep(time.Second)
}

func numbers() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func alphabet() {
	for i := 97; i < 107; i++ {
		// Melakukan konversi dari ASCII ke string 
		// agar mudah dalam mengiterasi a - j
		fmt.Println(string(byte(i)))
	}
}


/*
	TASK 2
*/
// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	go numbers(&wg)
// 	go alphabet(&wg)
// 	wg.Wait()
// }

// func numbers(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 1; i <= 10; i++ {
// 		fmt.Println(i)
// 	}
// }

// func alphabet(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 97; i < 107; i++ {
// 		fmt.Println(string(byte(i)))
// 	}
// }


/*
	TASK 3
*/
// func main() {
// 	ch := make(chan int)
// 	go numbers(ch)

// 	for v := range ch {
// 		fmt.Println(v)
// 	}
// }

// func numbers(ch chan<- int) {
// 	for i := 1; i <= 10; i++ {
// 		ch <- i
// 	}
// 	close(ch)
// }


/*
	TASK 4
	Dengan buffered channel, function terasa lebih cepat
	karena ketika ketika function ingin melakukan sending data melalui channel,
	tidak perlu saling menunggu seperti pada unbuffered channel
*/
// func main() {
// 	ch := make(chan int, 5)
// 	go numbers(ch)

// 	for v := range ch {
// 		fmt.Println(v)
// 	}
// }

// func numbers(ch chan<- int) {
// 	for i := 1; i <= 10; i++ {
// 		ch <- i
// 	}
// 	close(ch)
// }


/*
	TASK 5
*/
// func main() {
// 	oddCh := make(chan int)
// 	evenCh := make(chan int)
// 	go numbers(oddCh, evenCh)

// 	for i := 0; i < 20; i ++ {
// 		select {
// 		case even := <- evenCh:
// 			fmt.Printf("Received an even number\t: %v\n", even)

// 		case odd := <- oddCh:
// 			fmt.Printf("Received an odd number\t: %v\n", odd)
// 		}
// 	}
// }

// func numbers(oddCh chan<- int, evenCh chan<- int) {
// 	for i := 1; i <= 20; i++ {
// 		if i%2 == 0  {
// 			evenCh <- i
// 		} else {
// 			oddCh <- i
// 		}
// 	}
// 	close(oddCh)
// 	close(evenCh)
// }

/*
	TASK 6
*/
// func main() {
// 	oddCh := make(chan int)
// 	evenCh := make(chan int)
// 	errorCh := make(chan int)
// 	go numbers(oddCh, evenCh, errorCh)

// 	for i := 0; i < 30; i ++ {
// 		select {
// 		case even := <- evenCh:
// 			fmt.Printf("Received an even number\t: %v\n", even)

// 		case odd := <- oddCh:
// 			fmt.Printf("Received an odd number\t: %v\n", odd)

// 		case error := <- errorCh:
// 			fmt.Printf("Error: number %v is greater than 20\n", error)
// 		}
// 	}
// }

// func numbers(oddCh chan<- int, evenCh chan<- int, errorCh chan<- int) {
// 	for i := 1; i <= 30; i++ {
// 		if i > 20 {
// 			errorCh <- i
// 		} else if i%2 == 0  {
// 			evenCh <- i
// 		} else {
// 			oddCh <- i
// 		}
// 	}
// 	close(errorCh)
// 	close(oddCh)
// 	close(evenCh)
// }