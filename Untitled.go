//
// Go Template: zExample\appGo.tpl
//
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Type in your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello %s", name)
}
