package main

import (
	"fmt"

	"github.com/ipsusila/so/inject/factory"
	_ "github.com/ipsusila/so/inject/pkga"
	_ "github.com/ipsusila/so/inject/pkgb"
)

func main() {
	a, _ := factory.New("pkga.Alpha")
	b, _ := factory.New("pkgb.Beta")
	c, _ := factory.New("pkga.Gamma")
	d, _ := factory.New("mypackage.NotExists")

	fmt.Printf("%T %T %T %T\n", a, b, c, d)
}
