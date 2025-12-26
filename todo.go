package main

import (
	"fmt"
	"strconv"
	"time"
)

var tasks []Task
var currentID int

// Добавление задачи
func AddTask(title string, dateStr string) {
	date, _ := time.Parse("2006-01-02", dateStr)
	currentID++
	// Записывает в срез tasks задачу
	// ID берет из currentID (для каждой новой задачи +1)
	// Title берет из title, который аргумент функции AddTask
	// Title берет из data, который аргумент функции data
	// Completed по умолчанию false, может быть перезаписан функцией CompleteTask
	tasks = append(tasks, Task{
		ID:        currentID,
		Title:     title,
		Date:      date,
		Completed: false,
	})
	fmt.Println("Задача добавлена!")

	err := saveTasks()
	if err != nil {
		fmt.Println("Ошибка сохранения:", err)
	}
}

// Удаление задачи
func DeleteTask(idStr string) {
	id, _ := strconv.Atoi(idStr)
	var newTasks []Task

	// Фильтруем удаляемую задачу
	for _, task := range tasks {
		if task.ID != id {
			newTasks = append(newTasks, task)
		}
	}

	// Переприсваиваем ID всем оставшимся задачам
	for i, task := range newTasks {
		task.ID = i + 1 // ID начинается с 1
		newTasks[i] = task
	}

	tasks = newTasks
	currentID = len(tasks) // Обновляем currentID до текущего количества задач
	fmt.Println("Задача удалена!")

	// Сохраняем
	err := saveTasks()
	if err != nil {
		fmt.Println("Ошибка сохранения:", err)
	}
}

// Вывод задач по дате
func ShowTasksByDate(dateStr string) {
	date, _ := time.Parse("2006-01-02", dateStr)
	fmt.Println("\nЗадачи на", date.Format("02.01.2006"))
	// Перебирает срез tasks, если Date.Format (дата текущей задачи)
	// совпадает с date.Format (указанная дата аргуметом функции ShowTasksByDate),
	// создает и записывает в переменную status ✓
	// если !task.Completed (то же что и task.Completed == false) переписывает переменную status на ✗
	for _, task := range tasks {
		if task.Date.Format("2006-01-02") == date.Format("2006-01-02") {
			status := "✓"
			if !task.Completed {
				status = "✗"
			}
			fmt.Printf("%d. [%s] %s\n", task.ID, status, task.Title)
		}
	}
}

// Отметить задачу как выполненную
func CompleteTask(idStr string) {
	id, _ := strconv.Atoi(idStr)
	// Перебирает срез tasks, если ID (ID ранее созданной задачи)
	// совпадает с id (id Указанный агрументом функции CompleteTask),
	// переписывает переменную Complited структуры Task в срезе tasks по этому id на true
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
			fmt.Println("Задача отмечена как выполненная!")

			err := saveTasks()
			if err != nil {
				fmt.Println("Ошибка сохранения:", err)
			}
			return
		}
	}
}
