package main

import "fmt"

func main() {
	fmt.Println("Loops in golang")

	// days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	// fmt.Println(days)

	// Way-1
	// for d:= 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// Way-2
	// for i := range days{
	// 	fmt.Println(days[i])
	// }

	// Way-3
	// for index, day := range days {
	// 	fmt.Printf("index is %v and the value is %v\n", index, day)
	// }

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 2 {
			goto lco
		}

		if rougueValue == 3 {
			rougueValue++
			continue
		}

		if rougueValue == 9 {
			break
		}

		fmt.Println("Value is ", rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("Jumping to the statement")
}
