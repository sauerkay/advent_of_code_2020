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

	data := strings.Split(string(input), "\n")

	for x := 0; x < len(data); x++ {
		for y := 0; y < len(data); y++ {
			if data[x] != data[y] && data[x] != "" && data[y] != "" {
				a, err := strconv.Atoi(data[x])
				if err != nil {
					log.Fatal(err)
				}
				b, err := strconv.Atoi(data[y])
				if err != nil {
					log.Fatal(err)
				}
				calc := a + b
				if calc == 2020 {
					fmt.Println(data[x], " + ", data[y], " = ", calc)
					multiply := a * b
					fmt.Println(multiply)
				}
				for z := 0; z < len(data); z++ {
					if data[z] != data[x] && data[z] != data[y] && data[z] != "" {
						c, err := strconv.Atoi(data[z])
						if err != nil {
							log.Fatal(err)
						}
						calc_2 := calc + c
						if calc_2 == 2020 {
							fmt.Println(data[x], " + ", data[y], "+", data[z], " = ", calc_2)
							multiply_2 := a * b * c
							fmt.Println(multiply_2)
						}
					}
				}
			}
		}
	}
}
