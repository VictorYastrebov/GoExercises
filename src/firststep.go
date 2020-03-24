package main

import (
	Exercises "./Exercises"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	Exercises.GetFileNamesWithSameStrings(os.Args[1:])

	tryStuff()
}

func tryStuff() {
	fmt.Println("Всем привет!")

	rand.Seed(time.Now().UnixNano())

	fmt.Println("Сейчас случайное число - это ", rand.Intn(100))

	fmt.Printf("Корень квадртаный из 8: %g\n", math.Sqrt(8))

	x := rand.Intn(50)
	y := rand.Intn(60)
	fmt.Printf("Сумма двух чисел %d и %d = %d\n", x, y, add(x, y))

	s1 := "Строка1"
	s2 := "Строка2"
	s1New, s2New := swap(s1, s2)
	fmt.Printf("Была строка %s и %s, а стали %s и %s", s1, s2, s1New, s2New)
}
