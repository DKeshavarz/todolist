package TodoApp

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var commandsMap map[string]func(command []string)(string,error)

func init(){
	commandsMap["add-board"] = addBoard
}
func Run(){
	for {
		input := getInput()
		process(input)
	}
}

func getInput()[]string{
	fmt.Print("todo-app > ")
	
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.Split(strings.TrimSpace(input), " ")
}

func process(input []string){
	fmt.Println(input)
}