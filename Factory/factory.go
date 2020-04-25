package main

import "errors"

type ShapeFactory struct{}

func (sf *ShapeFactory) GetShape(shapeName string) (Figure, error) {
	switch shapeName {
	case CIRCLE:
		return Circle{}, nil
	case SQUARE:
		return Square{}, nil
	default:
		return nil, errors.New("shape not supported")
	}
}
