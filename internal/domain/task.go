package domain

type TaskID string

type Task struct {
	ID TaskID `json:"id"`
}
