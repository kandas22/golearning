package basics

import (
	"fmt"
)

func Primitive() {
	fmt.Print("Primitive DataTypes\n")
	var (
		x    int
		y        = 20
		z    int = 30
		d, e     = 40, "hello"
		f, g string
	)
	fmt.Println(x, y, z, d, e, f, g)
}
