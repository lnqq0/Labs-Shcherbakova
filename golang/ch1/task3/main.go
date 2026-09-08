package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// Медленный способ через цикл
	start1 := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Printf("Цикл занял: %v\n", time.Since(start1))

	// Быстрый способ
	start2 := time.Now()
	_ = strings.Join(os.Args[1:], " ")
	fmt.Printf("strings.Join занял: %v\n", time.Since(start2))
}
