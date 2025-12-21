package main

import "time"

// Структура для хранения задачи
type Task struct {
	ID        int
	Title     string
	Date      time.Time
	Completed bool
}
