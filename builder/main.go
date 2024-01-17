package main

import (
	"fmt"
)

func main() {
	p := NewPersonBuilder()

	p.
		Lives().At("15th Avenue").In("New York").WithPostcode("200124587").
		Works().At("WEX").AsA("Senior Software Engineer").Earning(120000)

	fmt.Println(p.Build())
}
