package task2

import "time"

type MyDuration struct {
	d time.Duration
}

func (dur MyDuration) MarshalJSON() ([]byte, error) {
	// TODO impl
}

func (dur *MyDuration) UnmarshalJSON(data []byte) error {
	// TODO impl
}
