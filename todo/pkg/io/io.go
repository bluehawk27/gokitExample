package io

import "encoding/json"

// Todo : represents the todo resource
type Todo struct {
	ID       int    `json:"id,omitempty" db:"id"`
	Title    string `json:"title"  db:"title"`
	Complete bool   `json:"complete" db:"complete"`
}

func (t Todo) String() string {
	b, err := json.Marshal(t)
	if err != nil {
		return "unsupported value type"
	}
	return string(b)
}
