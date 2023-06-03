package models

type Task struct {
	Id     int    `gorm:"primaryKey AutoIncrement" json:"id"`
	Task   string `json:"task"`
	IsDone bool   `json:"isDone"`
}
