package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Task представляет задачу
type Task struct {
	ID        int
	Title     string
	Completed bool
}

// TodoList хранит список задач
type TodoList struct {
	tasks  []Task
	nextID int
}

// NewTodoList создает новый список задач
func NewTodoList() *TodoList {
	return &TodoList{
		tasks:  make([]Task, 0),
		nextID: 1,
	}
}

func main() {
	todoList := NewTodoList()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("📋 Менеджер задач")
	fmt.Println("==================")

	for {
		fmt.Println("\nКоманды:")
		fmt.Println("  list    - показать все задачи")
		fmt.Println("  add     - добавить задачу")
		fmt.Println("  done    - отметить задачу как выполненную")
		fmt.Println("  delete  - удалить задачу")
		fmt.Println("  clear   - очистить выполненные задачи")
		fmt.Println("  exit    - выйти из программы")
		fmt.Print("\nВведите команду: ")

		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(strings.ToLower(command))

		switch command {
		case "list", "l":
			listTasks(todoList)
		case "add", "a":
			addTask(reader, todoList)
		case "done", "d":
			markTaskDone(reader, todoList)
		case "delete", "del":
			deleteTask(reader, todoList)
		case "clear", "c":
			clearCompleted(todoList)
		case "exit", "quit", "e", "q":
			fmt.Println("До свидания!")
			return
		default:
			fmt.Println("❌ Неизвестная команда")
		}
	}
}

func listTasks(todoList *TodoList) {
	if len(todoList.tasks) == 0 {
		fmt.Println("📭 Список задач пуст!")
		return
	}

	fmt.Println("\n📋 Ваши задачи:")
	for _, task := range todoList.tasks {
		status := "❌"
		if task.Completed {
			status = "✅"
		}
		fmt.Printf("  %s [%d] %s\n", status, task.ID, task.Title)
	}
}

func addTask(reader *bufio.Reader, todoList *TodoList) {
	fmt.Print("Введите название задачи: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	if name == "" {
		fmt.Println("❌ Название задачи не может быть пустым!")
		return
	}

	task := Task{
		ID:        todoList.nextID,
		Title:     name,
		Completed: false,
	}
	todoList.tasks = append(todoList.tasks, task)
	todoList.nextID++

	fmt.Printf("✅ Задача добавлена с ID %d\n", task.ID)
}

func markTaskDone(reader *bufio.Reader, todoList *TodoList) {
	if len(todoList.tasks) == 0 {
		fmt.Println("📭 Список задач пуст")
		return
	}

	fmt.Print("Введите ID задачи: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("❌ Некорректный ID")
		return
	}

	for i, task := range todoList.tasks {
		if task.ID == id {
			if task.Completed {
				fmt.Println("⚠️ Задача уже выполнена")
				return
			}
			todoList.tasks[i].Completed = true
			fmt.Printf("✅ Задача '%s' отмечена как выполненная\n", task.Title)
			return
		}
	}
	fmt.Println("❌ Задача с таким ID не найдена")
}

func deleteTask(reader *bufio.Reader, todoList *TodoList) {
	if len(todoList.tasks) == 0 {
		fmt.Println("📭 Список задач пуст")
		return
	}

	fmt.Print("Введите ID задачи для удаления: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("❌ Некорректный ID")
		return
	}

	for i, task := range todoList.tasks {
		if task.ID == id {
			// Удаляем элемент из среза
			todoList.tasks = append(todoList.tasks[:i], todoList.tasks[i+1:]...)
			fmt.Printf("✅ Задача '%s' удалена\n", task.Title)
			return
		}
	}
	fmt.Println("❌ Задача с таким ID не найдена")
}

func clearCompleted(todoList *TodoList) {
	if len(todoList.tasks) == 0 {
		fmt.Println("📭 Список задач пуст")
		return
	}

	// Подсчитываем количество выполненных задач
	completedCount := 0
	for _, task := range todoList.tasks {
		if task.Completed {
			completedCount++
		}
	}

	if completedCount == 0 {
		fmt.Println("📭 Нет выполненных задач")
		return
	}

	// Создаем новый срез только с невыполненными задачами
	newTasks := make([]Task, 0, len(todoList.tasks)-completedCount)
	for _, task := range todoList.tasks {
		if !task.Completed {
			newTasks = append(newTasks, task)
		}
	}
	todoList.tasks = newTasks

	fmt.Printf("✅ Удалено %d выполненных задач\n", completedCount)
}
