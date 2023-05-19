package main

import "fmt"

func main() {
	notas := make(map[string][]int)

	notas["guilherme"] = []int{10, 8, 2, 3}

	notas["namoradadoguilherme"] = []int{10, 5, 2, 4}

	notas["mayumi"] = []int{5, 3, 4, 5}
	fmt.Println(notas)

	mediagui := 10 + 8 + 2 + 3/4

	medianamoradadogui := 10 + 5 + 2 + 4/4

	mediamayumi := 5 + 3 + 4 + 5/4

	mediadasnotas := make(map[string]int)

	mediadasnotas["guilherme"] = mediagui

	mediadasnotas["namoradadoguilherme"] = medianamoradadogui

	mediadasnotas["mayumi"] = mediamayumi

	fmt.Println(mediadasnotas)

}
