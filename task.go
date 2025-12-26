package main

import "time"

// Структура для хранения задачи
type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Date      time.Time `json:"date"`
	Completed bool      `json:"completed"`
}
