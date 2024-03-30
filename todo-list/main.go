package main

import (
	"fmt"
	"os"
)

type todoItem struct {
	id    int
	title string
}

var todoList []todoItem

func addTodoItem(title string) {
	newItem := todoItem{
		id:    len(todoList) + 1,
		title: title,
	}
	todoList = append(todoList, newItem)
}

func listTodoItems() {
	fmt.Println("To-do List:")
	for _, item := range todoList {
		fmt.Printf("%d. %s\n", item.id, item.title)
	}
}

func main() {
	fmt.Println("Welcome to To-doooo!")
	fmt.Println("How do you do?")
	fmt.Println("Is there a task I can store for you?")

	for {
		fmt.Println("\nSelect and option:")
		fmt.Println("1. Add a todo list item")
		fmt.Println("2. List all todo items")
		fmt.Println("3. Exit")
		var choice int
		fmt.Println("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter the title of the todo item: ")
			var title string
			fmt.Scanln(&title)
			addTodoItem(title)
			fmt.Println("Todo item added successfully!")
		case 2:
			listTodoItems()
		case 3:
			println("Exiting...")
			os.Exit(0)
		}
	}
}
