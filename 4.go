package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func do_4() {
	readFile, err := os.Open("./input.in")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	fileScanner.Scan()
	_, err = strconv.Atoi(fileScanner.Text())
	if err != nil {
		panic(err)
	}

	A := 0
	B := 0
	K := 0

	c := 1 // Case 1

	i := 0
	for fileScanner.Scan() {
		if i == 0 {
			A, err = strconv.Atoi(fileScanner.Text())
			if err != nil {
				panic(err)
			}
		}
		if i == 1 {
			B, err = strconv.Atoi(fileScanner.Text())
			if err != nil {
				panic(err)
			}
		}
		if i == 2 {
			K, err = strconv.Atoi(fileScanner.Text())
			if err != nil {
				panic(err)
			}

			i = 0

			count := 0
			for x := A; x <= B; x++ {
				if x%K == 0 {
					count += 1
				}
			}
			fmt.Println("Case " + strconv.Itoa(c) + ": " + strconv.Itoa(count))
			c++
			continue
		}
		i++
	}

	readFile.Close()
}
