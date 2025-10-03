package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("PINGPANG")
		} else if i%3 == 0 {
			fmt.Println("PING")
		} else if i%5 == 0 {
			fmt.Println("PANG")
		} else {
			fmt.Printf("%d\n", i)
		}
	}
}
