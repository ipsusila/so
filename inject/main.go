package main

import (
	"fmt"

	"github.com/ipsusila/so/inject/factory"
	_ "github.com/ipsusila/so/inject/pkga"
	_ "github.com/ipsusila/so/inject/pkgb"
)

func main() {
	a := factory.New("pkga.Alpha")
	b := factory.New("pkgb.Beta")
	c := factory.New("pkga.Gamma")

	fmt.Printf("%T %T %T\n", a, b, c)
}
