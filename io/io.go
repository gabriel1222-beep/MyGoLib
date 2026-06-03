package io

import(
	"os"
	"bufio"
	"strings"
)
func getString() (string,error){
	reader := bufio.NewReader(os.Stdin)
	text,err := reader.ReadString('\n')
	if err != nil{
		return "",err
	}
	return strings.TrimSpace(text),nil
}
