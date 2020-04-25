package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func nextLine() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	inText := nextLine()
	inputs := strings.Split(inText, " ")

	price, _ := strconv.Atoi(inputs[0])
	rate, _ := strconv.ParseFloat(inputs[1], 10)
	rate /= 100

	totalPay := 0

	for {
		totalPay += price
		price = int(float64(price) * (1 - rate))

		if price < 1 {
			break
		}
	}

	fmt.Println(totalPay)

}
