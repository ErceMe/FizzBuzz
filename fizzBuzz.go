package main

import "fmt"

func main() {
	var i, j int
	fmt.Println("Masukkan Angka : ")
	fmt.Scan(&i)
	fmt.Println("Hasil :")
	for j = 1; j <= i; j++ {
		if (j%3 == 0) && (j%5 == 0) {
			fmt.Println("FizzBuzz")
		} else if j%3 == 0 {
			fmt.Println("Fizz")
		} else if j%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(j)
		}
	}
}
