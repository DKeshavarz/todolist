package TodoApp

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Run(){
	for {
		input := getInput()
		fmt.Println(input)
	}
}

func getInput()[]string{
	fmt.Print("todo-app > ")
	
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.Split(strings.TrimSpace(input), " ")
}