package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Do the given operation on the tape
func doOp(index int, tape []int) (bool, *int) {
	op := tape[index]
switch op {
case 99:
	return true, &tape[0]
case 1:
	loc1 := tape[index+1]
	loc2 := tape[index+2]
	loc3 := tape[index+3]
	tape[loc3] = tape[loc1] + tape[loc2]
	return false, nil
case 2:
	loc1 := tape[index+1]
	loc2 := tape[index+2]
	loc3 := tape[index+3]
	tape[loc3] = tape[loc1] * tape[loc2]
	return false, nil
default:
	return false, &tape[0]
}
}
// printOp ...
func printOp(index int, tape []int) (string,error) {
	op := tape[index]
	switch op {
	case 99:
		return fmt.Sprintf("Halt\n"),nil
	case 1:
		loc1 := tape[index+1]
		loc2 := tape[index+2]
		return fmt.Sprintf("Add:%d + %d = %d written to %d\n", tape[loc1],tape[loc2],tape[loc1] + tape[loc2],tape[index+3]),nil
	case 2:
		loc1 := tape[index+1]
		loc2 := tape[index+2]
		return fmt.Sprintf("Mul:%d * %d =%d written to %d\n", tape[loc1],tape[loc2],tape[loc1] * tape[loc2],tape[index+3]),nil
	default:
		return string(""), fmt.Errorf("Unknown Operation %d",op)
	}

}
// runTape Runs the current tape
func runTape(tape []int) (int,error) {
	finished := false
	var result *int
	for i := 0; finished != true; i += 4 {
		// op,err := printOp(i,tape)
		// if err != nil {
		// 	return 0,fmt.Errorf("Failed to print operation:%w",err)

		// }
		//fmt.Printf(op)
		finished,result = doOp(i, tape)
		if !finished && result!=nil{
			s,err:=printOp(i,tape)
			return 0,fmt.Errorf("Invalid Operand at %d\n Tape:%v %w",i,s,err)
		}
	}
	return *result,nil
}
func main() {
	fmt.Printf("hello, world\n")
	data, err := ioutil.ReadFile("programs.txt")
	if err != nil {
		fmt.Printf("File Error %w", err)
		os.Exit(4)
	}

	program := strings.Split(strings.TrimSpace(string(data)), ",")
	tape := make([]int,len(program))
	for i, s := range program {
		tape[i], err = strconv.Atoi(s)
		if err != nil {
			fmt.Printf("Strconv Error %v", err)
			os.Exit(5)
		}
	}
	for i:=0;i<100;i++{
		for j:=0;j<100;j++{
			working := make([]int,len(tape))
			copy(working,tape)
			working[1]=i
			working[2]=j

			result,err := runTape(working)
			if err != nil {
				fmt.Printf("Tape Error %w", err)
				os.Exit(1)
			}
			if result == 19690720 {
				fmt.Printf("Answer: %d",i*100+j)
			}

		}
	}
}
