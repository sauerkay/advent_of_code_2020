package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	testData := strings.Split(string(input), "\n")

	for i := 0; i < len(testData); i++ {
		for x := 0; x < len(testData); x++ {
			if testData[i] != testData[x] && testData[i] != "" && testData[x] != "" {
				a, err := strconv.Atoi(testData[i])
				if err != nil {
					log.Fatal(err)
				}
				b, err := strconv.Atoi(testData[x])
				if err != nil {
					log.Fatal(err)
				}
				calc := a + b
				if calc == 2020 {
					fmt.Println(testData[i], " + ", testData[x], " = ", calc)
					multiply := a * b
					fmt.Println(multiply)
				}
			}
		}
	}
}
