package task1

import (
	"context"
)

var _ Worker = &Worker1{}

type Worker1 struct{}

func (w *Worker1) Run(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
