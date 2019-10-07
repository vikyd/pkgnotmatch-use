package main

import (
	"fmt"

	// can use `github.com/vikyd/pkgnotmatch` directly
	// but `abc "github.com/vikyd/pkgnotmatch"` will be better
	"github.com/vikyd/pkgnotmatch"
)

func main() {
	fmt.Println("main")
	abc.F01()

}
