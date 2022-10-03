package main

import (
	"bufio"
	"converter_app/temperature"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	ans, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return ans
}

func main() {

	itemsToConvert := []string{"Temperature", "Length", "Mass"}

	fmt.Println("What do you want to convert?")

	for i, v := range itemsToConvert {
		fmt.Printf("For %s Press %d \n", v, i)
	}

	var whatToConvert int
	for {
		fmt.Print("Enter a number: ")

		ans := readInput()

		value, err := strconv.Atoi(strings.TrimSpace(ans))

		if err != nil {
			log.Fatal(err)
			fmt.Printf("%s is not an Integer \n", ans)
		} else {
			whatToConvert = value
			break
		}
	}

	switch whatToConvert {
	case 1:
		temperature.Start()
		break
	}

}
