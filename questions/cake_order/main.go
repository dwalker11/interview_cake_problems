package main

import "fmt"

func order_checker(takeOutOrders []int, dineInOrders []int, combined []int) bool {
	i := 0
	j := 0

	for _, v := range combined {
		if i != len(takeOutOrders) && v == takeOutOrders[i] {
			i++
		} else if j != len(dineInOrders) && v == dineInOrders[j] {
			j++
		} else {
			return false
		}
	}

	return true
}

func main() {
	takeOutOrders := []int{1, 3, 5}
	dineInOrders := []int{2, 4, 6}
	combined := []int{1, 2, 4, 3, 6, 5}
	result := order_checker(takeOutOrders, dineInOrders, combined)
	fmt.Println(result)
}
