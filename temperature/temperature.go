package temperature

import (
	"fmt"
)

type Temperature struct {
	value     float64
	convertId int
}

func Start() {
	fmt.Println()
	fmt.Println("Loading Temperature Converter ....")
	fmt.Println()
}
