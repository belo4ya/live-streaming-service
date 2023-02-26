package biz

import "context"

type Stream struct {
	Id       int
	Name     string
	Username string
}

type StreamRepo interface {
	ListAll(context.Context) ([]*Stream, error)
}

type StreamUseCase struct {
}

func NewStreamUseCase() {
}
