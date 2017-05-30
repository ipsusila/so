package pkga

import "github.com/ipsusila/so/inject/factory"

type thisFactory struct {
}

type Alpha struct{}
type Beta struct{}
type Gamma struct{}

func (f *thisFactory) New(name string) interface{} {
	switch name {
	case "Alpha":
		return &Alpha{}
	case "Beta":
		return &Beta{}
	case "Gamma":
		return &Gamma{}
	}
	return nil
}

func init() {
	factory.Register("pkga", &thisFactory{})
}
