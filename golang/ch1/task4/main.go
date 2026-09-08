package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	// Мапа, чтобы запоминать, в каких файлах встретилась конкретная строка
	lineFiles := make(map[string][]string)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, "stdin", lineFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "task4: %v\n", err)
				continue
			}
			countLines(f, counts, arg, lineFiles)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("Повторов: %d\tСтрока: %q\tФайлы: %v\n", n, line, lineFiles[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, filename string, lineFiles map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++

		// Проверяем, записывали ли мы уже этот файл для этой строки
		alreadyAdded := false
		for _, fName := range lineFiles[line] {
			if fName == filename {
				alreadyAdded = true
				break
			}
		}
		if !alreadyAdded {
			lineFiles[line] = append(lineFiles[line], filename)
		}
	}
}
