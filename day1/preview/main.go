package main

import (
	"fmt"
)

// Task 1
// func main() {

// 	// Number
// 	for i := 1; i <= 10; i++ {
// 		go func(i int) {
// 			fmt.Println(i)
// 		}(i)
// 	}

// 	// Alphabet
// 	for i := 'a'; i <= 'j'; i++ {
// 		go func(i rune) {
// 			fmt.Println(string(i))
// 		}(i)
// 	}

// 	time.Sleep(3 * time.Second)
// }

// Task 2
// func main() {
// 	var wg, wg2 sync.WaitGroup

// 	// Number
// 	for i := 1; i <= 10; i++ {
// 		wg.Add(1)
// 		go func(i int) {
// 			defer wg.Done()
// 			fmt.Println(i)
// 		}(i)
// 	}

// 	// Alphabet
// 	for i := 'a'; i <= 'j'; i++ {
// 		wg2.Add(1)
// 		go func(i rune) {
// 			defer wg2.Done()
// 			fmt.Println(string(i))
// 		}(i)
// 	}

// 	wg.Wait()
// 	wg2.Wait()
// }

// Task 3
// func produce(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 1; i <= 10; i++ {
// 		ch <- i
// 	}
// 	close(ch)
// }

// func consume(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for num := range ch {
// 		fmt.Println(num)
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup
// 	ch := make(chan int)

// 	wg.Add(1)
// 	go produce(ch, &wg)
// 	wg.Add(1)
// 	go consume(ch, &wg)

// 	wg.Wait()
// }

// Task 4
// func produce(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 1; i <= 10; i++ {
// 		ch <- i
// 	}
// 	close(ch)
// }

// func consume(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for num := range ch {
// 		fmt.Println(num)
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup
// 	ch := make(chan int, 5)

// 	wg.Add(1)
// 	go produce(ch, &wg)
// 	wg.Add(1)
// 	go consume(ch, &wg)

// 	wg.Wait()
// }

// Task 5
func sendNumbers(evenCh, oddCh chan int) {
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			evenCh <- i
		} else {
			oddCh <- i
		}
	}
	close(evenCh)
	close(oddCh)
}

func main() {
	evenCh := make(chan int, 5)
	oddCh := make(chan int, 5)

	go sendNumbers(evenCh, oddCh)

	for {
		select {
		case evenNum, err := <-evenCh:
			if !err {
				evenCh = nil
			} else {
				fmt.Printf("Received an even number: %d\n", evenNum)
			}
		case oddNum, err := <-oddCh:
			if !err {
				oddCh = nil
			} else {
				fmt.Printf("Received an odd number: %d\n", oddNum)
			}
		}

		// Stop after all numbers have been printed
		if evenCh == nil && oddCh == nil {
			break
		}
	}
}
