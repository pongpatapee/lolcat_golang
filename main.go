package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	// "strings"
)

func rgb(i int) (int, int, int) {
	f := 0.1

	r := int(math.Sin(f*float64(i)+0)*127 + 128)
	g := int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128)
	b := int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)

	return r, g, b
}

func printRainbow(output []rune) {
	for i := 0; i < len(output); i++ {

		r, g, b := rgb(i)
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, output[i])
	}

	fmt.Println()
}

func main() {
	info, _ := os.Stdin.Stat()

	var output []rune

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("Intended to be used with pipe")
		fmt.Println("Usage: some program printing to stdout | lolcat_golang")
		os.Exit(0)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		input, _, err := reader.ReadRune()

		if err != nil && err == io.EOF {
			break
		}

		output = append(output, input)
	}

	printRainbow(output)
}
