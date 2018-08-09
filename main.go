package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	fmt.Println(generateTckn())
}

func generateTckn() string {
	rand.Seed(time.Now().Unix())
	var (
		digits [11]int
		tckn   string
	)

	digits[0] = random(1, 10)
	for i := 1; i < 9; i++ {
		digits[i] = random(0, 10)
	}

	digits[9] = ((digits[0]+digits[2]+digits[4]+digits[6]+digits[8])*7 -
		(digits[1] + digits[3] + digits[5] + digits[7])) % 10
	digits[10] = (digits[0] + digits[1] + digits[2] + digits[3] + digits[4] + digits[5] +
		digits[6] + digits[7] + digits[8] + digits[9]) % 10
	for i := 0; i < len(digits); i++ {
		tckn += strconv.Itoa(digits[i])
	}
	return tckn
}
