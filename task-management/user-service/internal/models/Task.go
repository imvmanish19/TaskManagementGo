package models

type Task struct {
    ID     string `json:"id"`
    Name   string `json:"name"`
    Status string `json:"status"`
    UserID string `json:"user_id"`
}
