package src

import "errors"

type Ball struct {
	Diameter int
}

type Box struct {
	Height int
	Width  int
	Length int
}

func NewBall(diameter int) (*Ball, error) {
	if diameter < 1 || diameter > 10000 {
		return nil, errors.New("error when creating model")
	}
	return &Ball{Diameter: diameter}, nil
}

func NewBox(height, width, length int) (*Box, error) {
	if height < 1 || height > 10000 || width < 1 || width > 10000 || length < 1 || length > 10000 {
		return nil, errors.New("error when creating model")
	}
	return &Box{Height: height, Width: width, Length: length}, nil
}