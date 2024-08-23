package task1

import "context"

type Worker interface {
	Run(ctx context.Context) error
}

type WorkerRunner struct {
	workers []Worker
}

func (r *WorkerRunner) Add(worker Worker) {
	// TODO зарегистрировать воркер
}

func (r *WorkerRunner) Run(ctx context.Context) error {
	// TODO запустить все воркеры параллельно
}

func (r *WorkerRunner) Stop() {
	// TODO остановить выполнение
}
