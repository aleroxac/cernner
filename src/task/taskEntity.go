package task

import (
	"time"

	"github.com/google/uuid"
)

var (
	Id          uuid.UUID
	Name        string
	Description string
	Owner       string
	Labels      []string
	Status      string
	CreatedAt   time.Time
	StartedAt   string
	StopedAt    string
	CanceledAt  string
	FinishedAt  string
	Duration    time.Duration
)

type Task struct {
	Id          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Owner       string        `json:"owner"`
	Labels      []string      `json:"labels"`
	Status      string        `json:"status"`
	CreatedAt   time.Time     `json:"created_at"`
	StartedAt   string        `json:"started_at"`
	StopedAt    string        `json:"stoped_at"`
	CanceledAt  string        `json:"canceled_at"`
	FinishedAt  string        `json:"finished_at"`
	Duration    time.Duration `json:"duration"`
}
