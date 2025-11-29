package cmd

import (
	"ToDoList/internal/todo"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Run() {
	scanner := bufio.NewScanner(os.Stdin) //это используется чтобы построчно читать введённый текст из терминала

	for { // бесконечный цикл/меню
		fmt.Println()
		fmt.Println("<<===My ToDo-list===>>")
		fmt.Println("1. Add task")
		fmt.Println("2. Show all tasks")
		fmt.Println("3. Mark as completed")
		fmt.Println("4. Exit")

		if !scanner.Scan() { // условие при котором если пользователь не ввёл ничего а scanner.Scan() равен false то цикл закрывается
			break
		}
		choice := scanner.Text() //так сканируется текст который ввел пользователь

		switch choice {

		case "1":
			fmt.Print("Write the task's name")
			if !scanner.Scan() { //опять же если пользователь не ввел ничего то переходим к следующей итеррации цикла. итерация это одно полное прохождение цикла
				continue
			}
			title := scanner.Text() // то что ввёл пользователь засовывается в эту переменную
			todo.AddTask(title)     // введенное пользователем отдаётся функции в другом файле для последующей обработки
			fmt.Println("The task successfully added")

		case "2":
			tasks := todo.ListTasks() //берём актуальный список задач из ядра(пакет todo.go)
			if len(tasks) == 0 {      // если список пуст ты вывести ошибку
				fmt.Println("The list is empty")

			} else {
				fmt.Println()
				fmt.Println("Task's list")
				for _, t := range tasks { //пройтись по каждой задаче для проставки меток выполнения
					status := ""     //по умолчанию пусто так как задача не выполнена
					if t.Completed { //если выполнена поставить эту галочку
						status = "✅"
					}
					fmt.Printf("%s %d. %s\n", status, t.ID, t.Title) //вывести на экран статус, номер задачи, и имя
				}
			}
		case "3":
			fmt.Println("Write the task's number/ID")
			if !scanner.Scan() { //если пользователь не ввёл ничего перейти к следующей итерации цикла
				continue
			}
			input := scanner.Text()         // если всё же есть ввод то передать её сюда
			num, err := strconv.Atoi(input) //преобразование строк в числа
			if err != nil || num < 1 {      //проверка на ошибку пустой или меньше 1
				fmt.Println("Wrong number/ID")
				continue
			}
			err = todo.CompleteTask(num) //если неь такого номера то вывод ошибки из ядра проекта
			if err != nil {              //если ошибка имеется то вывести её
				fmt.Println("Error: ", err)
			} else {
				fmt.Println("The task is marked as 'Completed' ") //если всё ок то вывести сообщение об успехе
			}
		case "4": // при этом сценарии завершить программу и выйти
			fmt.Println("Have a productive day")
			return
		default: //по умолчанию вернуть сообщение что неправильный номер выбора в меню
			fmt.Println("Choose the options between 1-4")
		}

	}
}
