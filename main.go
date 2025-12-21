package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Список дел. Введите help для списка команд")

	// Используем bufio для более надежной работы с вводом
	reader := bufio.NewReader(os.Stdin)

	for {
		// Добавляет > для удобство чтения
		fmt.Print("\n> ")
		// ReadString читает до новой строки и возвращает в значения: ввод (в input) и ошибку (в err)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			continue
		}

		// Удаляем символ новой строки
		input = strings.TrimSpace(input)
		// Пропускает пустые строки
		if input == "" {
			continue
		}

		// Разделяет ввод на отдельные элементы используя пробел как разделитель
		// Создает срез строк в который записывает элементы
		parts := strings.Split(input, " ")
		// Приводит первый элемент среза к нижнему регистру
		command := strings.ToLower(parts[0])

		switch command {
		case "add":
			// Проверка на правильность ввода. В срезе должно быть 3 элемента
			if len(parts) < 3 {
				fmt.Println("Используйте: add <название> <дата>")
				continue
			}
			// Берет элементы с 2 (индекс 1) до предпоследнего учитывая пробел как разделитель
			// Обьединяет и записывает в title
			title := strings.Join(parts[1:len(parts)-1], " ")
			// Берет последний элемент среза и записывает в data
			date := parts[len(parts)-1]
			AddTask(title, date)

		case "delete":
			if len(parts) != 2 {
				fmt.Println("Используйте: delete <id>")
				continue
			}
			DeleteTask(parts[1])

		case "show":
			if len(parts) != 2 {
				fmt.Println("Используйте: show <дата>")
				continue
			}
			ShowTasksByDate(parts[1])

		case "complete":
			if len(parts) != 2 {
				fmt.Println("Используйте: complete <id>")
				continue
			}
			CompleteTask(parts[1])

		case "help":
			fmt.Println("Команды:")
			fmt.Println("  add <название> <дата> - добавить задачу")
			fmt.Println("  delete <id> - удалить задачу")
			fmt.Println("  show <дата> - показать задачи")
			fmt.Println("  complete <id> - отметить задачу как выполненную")
			fmt.Println("  Примечание: Формат даты вводить вида Год-Месяц-День")

		case "exit":
			fmt.Println("=ЧАО ПОКА!")
			time.Sleep(3 * time.Second) // Задержка в 3 секунды
			return

		default:
			fmt.Println("Неизвестная команда. Введите help для списка команд")
		}
	}
}
