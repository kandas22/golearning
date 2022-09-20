package basics

import "fmt"

const x int64 = 10
const (
	idKey   = "id"
	nameKey = "name"
)
const z = 20 * 10

func ConstFunc() {
	fmt.Println(idKey)
	fmt.Println(nameKey)
}
