package main

import "fmt"

type State struct {
	Name  string
	Count int
}

type Command struct {
	Type   string
	Value  string
	Result chan State
}

func manager(commands chan Command) {
	state := State{
		Name:  "Student",
		Count: 0,
	}

	for command := range commands {
		switch command.Type {
		case "setName":
			state.Name = command.Value
			command.Result <- state

		case "add":
			state.Count++
			command.Result <- state

		case "get":
			command.Result <- state
		}
	}
}

func main() {
	commands := make(chan Command)
	result := make(chan State)

	go manager(commands)

	commands <- Command{
		Type:   "setName",
		Value:  "Alex",
		Result: result,
	}

	fmt.Println(<-result)

	commands <- Command{
		Type:   "add",
		Result: result,
	}

	fmt.Println(<-result)

	commands <- Command{
		Type:   "add",
		Result: result,
	}

	fmt.Println(<-result)

	commands <- Command{
		Type:   "get",
		Result: result,
	}

	fmt.Println(<-result)
}
