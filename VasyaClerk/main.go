package main

func verifiCambioDisponible(bill int, dollarAvailable *map[int]int) bool {
	if bill == 100 {
		if (*dollarAvailable)[50] > 0 && (*dollarAvailable)[25] > 0 {
			(*dollarAvailable)[50]--
			(*dollarAvailable)[25]--
			(*dollarAvailable)[100]++
			return true
		} else if (*dollarAvailable)[25] > 2 {
			(*dollarAvailable)[100]++
			(*dollarAvailable)[25] -= 3
			return true
		}
	} else if bill == 50 {
		if (*dollarAvailable)[25] > 0 {
			(*dollarAvailable)[25]--
			(*dollarAvailable)[50]++
			return true
		}
	} else {
		(*dollarAvailable)[25]++
		return true
	}
	return false
}

func main() {
}
