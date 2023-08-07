package main

import (
	"fmt"

	"github.com/Thashmi03/gopackagedemo/calc"
)

func main(){
	a:=40
	b:=20
	x:=calc.Add(a,b)	
	fmt.Println(x)
}