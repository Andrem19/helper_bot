// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"time"
)

type Command struct {
	ID        int64     `json:"id"`
	Cmd       string    `json:"cmd"`
	CreatedAt time.Time `json:"created_at"`
}
