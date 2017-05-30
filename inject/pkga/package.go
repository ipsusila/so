package pkga

import "github.com/ipsusila/so/inject/factory"

type thisFactory struct {
}

type Alpha struct{}
type Beta struct{}
type Gamma struct{}

func (f *thisFactory) New(name string) (interface{}, bool) {
	switch name {
	case "Alpha":
		return &Alpha{}, true
	case "Beta":
		return &Beta{}, true
	case "Gamma":
		return &Gamma{}, true
	}
	return nil, false
}

func init() {
	factory.Register("pkga", &thisFactory{})
}
