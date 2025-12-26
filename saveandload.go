package main

import (
	"encoding/json"
	"os"
)

const dataFile = "tasks.json"

// Сохраняет задачи в JSON-файл
func saveTasks() error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, data, 0644)
}

// Загружает задачи из JSON-файла
func loadTasks() error {
	fileData, err := os.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			// Файл не найден — начинаем с пустого списка
			tasks = []Task{}
			currentID = 0
			return nil
		}
		return err
	}

	err = json.Unmarshal(fileData, &tasks)
	if err != nil {
		return err
	}

	// Обновляем currentID, чтобы новые задачи получали корректный ID
	for _, task := range tasks {
		if task.ID > currentID {
			currentID = task.ID
		}
	}
	return nil
}
