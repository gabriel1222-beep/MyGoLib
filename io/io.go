package io

import(
	"os"
	"bufio"
	"strings"
)
func GetString() (string,error){
	reader := bufio.NewReader(os.Stdin)
	text,err := reader.ReadString('\n')
	if err != nil{
		return "",err
	}
	return strings.TrimSpace(text),nil
}
func GetByte() (byte,error){
	reader := bufio.NewReader(os.Stdin)
	b,err := reader.ReadByte()
	if err != nil{
		return byte(0),err
	}
	return b,nil
}
func GetRune() (rune,error){
	reader := bufio.NewReader(os.Stdin)
	r,_,err := reader.ReadRune()
	if err != nil{
		return rune(0),err
	}
	return r,nil
}
