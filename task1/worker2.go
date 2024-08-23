package task1

import "context"

var _ Worker = &Worker2{}

type Worker2 struct{}

func (w *Worker2) Run(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
