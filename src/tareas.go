package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) addToList(tl *task) {
	t.tasks = append(t.tasks, tl)
}

func (t *taskList) deleteFromList(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (t *taskList) printList() {
	for _, task := range t.tasks {
		println("Nombre: ", task.name)
		println("Descripción: ", task.description)
	}
}

func (t *taskList) printListComplete() {
	println("Tareas completadas")
	for _, task := range t.tasks {
		if task.completed {
			println("Nombre: ", task.name)
			println("Descripción: ", task.description)
		}
	}
}

type task struct {
	name        string
	description string
	completed   bool
}

func (t *task) markComplete() {
	t.completed = true
}

func (t *task) updateDescription(description string) {
	t.description = description
}

func (t *task) updateName(name string) {
	t.name = name
}

func main() {
	t := &task{
		name:        "Completar curso",
		description: "Completar el curso de Go de Platzi en esta semana",
		completed:   false,
	}

	t1 := &task{
		name:        "Completar curso javascript",
		description: "Completar el curso de javascript de Platzi en esta semana",
		completed:   false,
	}

	list := &taskList{
		tasks: []*task{
			t, t1,
		},
	}

	t2 := &task{
		name:        "Completar curso Node.js",
		description: "Completar el curso de Node.js de Platzi en esta semana",
	}
	list.addToList(t2)
	list.printList()
	list.tasks[0].markComplete()
	list.printListComplete()

	mapTasks := make(map[string]*taskList)

	mapTasks["Juan"] = list

	t3 := &task{
		name:        "Completar curso C#",
		description: "Completar el curso de C# de Platzi en esta semana",
		completed:   false,
	}

	t4 := &task{
		name:        "Completar curso C++",
		description: "Completar el curso de C++ de Platzi en esta semana",
		completed:   false,
	}

	list2 := &taskList{
		tasks: []*task{
			t3, t4,
		},
	}

	mapTasks["Pedro"] = list2

	fmt.Println("Tareas de Juan")
	mapTasks["Juan"].printList()
	fmt.Println("Tareas de Pedro")
	mapTasks["Pedro"].printList()
}
