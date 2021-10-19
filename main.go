package main

import "fmt"

func main() {
	fmt.Println("Введите четырехзначный номер билета")
	var number int

	fmt.Scan(&number)

	digit1 := (number / 1000)
	digit2 := (number % 1000) / 100
	digit3 := (number % 100) / 10
	digit4 := (number % 10)

	if digit1 == digit4 && digit2 == digit3{
    fmt.Println("Зеркальный билет")
  }else if digit1 + digit2 == digit3 + digit4{
    fmt.Println("Счастливый билет")
  }else{
    fmt.Println("Простой билет")
  }

}
