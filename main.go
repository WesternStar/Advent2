package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"errors"
)

type HaltError struct {
}

func (e HaltError) Error() string {
	return fmt.Sprintf("Halted!")
}
type UnknownModeError struct {
}

func (e UnknownModeError) Error() string {
	return fmt.Sprintf("Unknown Mode!")
}

type EmptyInputError struct{}

func (e EmptyInputError) Error() string {
	return fmt.Sprintf("Empty Input!")
}

type UnknownOpError struct {
	op int
}

func (e UnknownOpError) Error() string {
	return fmt.Sprintf("Unknown Operation:%v", e.op)
}


// Do the given operation on the tape
func doOp(index *int, tape, input, output *[]int) (*int, error) {
	splitOp:=func (op int) (int,[]int){
		hundredThou:=op/100000
		tenThou:=(op- hundredThou*100000)/10000
		thou:=(op - hundredThou*100000-tenThou*10000)/1000
		hun := (op - hundredThou*100000-tenThou*10000 - thou*1000)/100
		res := (op - hundredThou*100000-tenThou*10000 - thou*1000 - hun*100)
		return res,[]int{tenThou,thou,hun}
	}
	getValue:= func(loc ,mode int) (int ,error){
		if mode ==0{
			return (*tape)[loc],nil
		}else if mode ==1{
			return loc,nil
		}else{
			return 0, UnknownModeError{}
		}
	}


	op,mode := splitOp((*tape)[*index])
	switch op {
	case 99:
		return &(*tape)[0], HaltError{}
	case 1:
		loc1 := (*tape)[*index+1]
		loc2 := (*tape)[*index+2]
		loc3 := (*tape)[*index+3]
	        //fmt.Printf("Opcode:%v, mode:%v, loc1:%v, loc2:%v\n",op,mode[2], loc1,loc2)
		val1,err:=getValue(loc1,mode[2])
		if err!=nil {
			return nil,err
		}
		val2,err:=getValue(loc2,mode[1])
		if err!=nil {
			return nil,err
		}
		(*tape)[loc3] = val1 + val2
		*index+=4
		return nil, nil
	case 2:
		loc1 := (*tape)[*index+1]
		loc2 := (*tape)[*index+2]
		loc3 := (*tape)[*index+3]
		val1,err:=getValue(loc1,mode[2])
		if err!=nil {
			return nil,err
		}
		val2,err:=getValue(loc2,mode[1])
		if err!=nil {
			return nil,err
		}
		(*tape)[loc3] = val1*val2
		*index+=4
		return nil, nil
	case 3:
		loc1 := (*tape)[*index+1]
		if len(*input) > 0 {
			(*tape)[loc1] = (*input)[0]
			*index+=2
			return nil, nil
		} else {
			return nil, EmptyInputError{}

		}
	case 4:
		loc1 := (*tape)[*index+1]
		val1,err:=getValue(loc1,mode[2])
		if err!=nil {
			return nil,err
		}
		*output = append(*output, val1)
		*index+=2
		return nil, nil
	case 5:
		loc1 := (*tape)[*index+1]
		loc2 := (*tape)[*index+2]
		val1,err:=getValue(loc1,mode[2])
		if err!=nil {
			return nil,err
		}
		val2,err:=getValue(loc2,mode[1])
		if err!=nil {
			return nil,err
		}
		if val1 !=0{
			*index=val2
		}else{
			*index+=3
		}
		return nil, nil
	case 6:
		loc1 := (*tape)[*index+1]
		loc2 := (*tape)[*index+2]
		val1,err:=getValue(loc1,mode[2])
		if err!=nil {
			return nil,err
		}
		val2,err:=getValue(loc2,mode[1])
		if err!=nil {
			return nil,err
		}
		if val1 ==0{
			*index=val2
		}else{
			*index+=3
		}
		return nil, nil
	case 7:
		loc1 := (*tape)[*index+1]
		loc2 := (*tape)[*index+2]
		loc3 := (*tape)[*index+3]
		val1,err:=getValue(loc1,mode[2])
		if err!=nil {
			return nil,err
		}
		val2,err:=getValue(loc2,mode[1])
		if err!=nil {
			return nil,err
		}
		if val1<val2{
			(*tape)[loc3]=1
		}else {
			(*tape)[loc3]=0
		}
		*index+=4
		return nil, nil
	case 8:
		loc1 := (*tape)[*index+1]
		loc2 := (*tape)[*index+2]
		loc3 := (*tape)[*index+3]
		val1,err:=getValue(loc1,mode[2])
		if err!=nil {
			return nil,err
		}
		val2,err:=getValue(loc2,mode[1])
		if err!=nil {
			return nil,err
		}
		if val1==val2{
			(*tape)[loc3]=1
		}else {
			(*tape)[loc3]=0
		}
		*index+=4
		return nil, nil
	default:
		return &(*tape)[0], UnknownOpError{op}
	}
}

// printOp ...
func printOp(index int, tape,input []int) (string, error) {
	splitOp:=func (op int) (int,[]int){
		hundredThou:=op/100000
		tenThou:=(op- hundredThou*100000)/10000
		thou:=(op - hundredThou*100000-tenThou*10000)/1000
		hun := (op - hundredThou*100000-tenThou*10000 - thou*1000)/100
		res := (op - hundredThou*100000-tenThou*10000 - thou*1000 - hun*100)
		return res,[]int{tenThou,thou,hun}
	}
	getValue:= func(loc ,mode int) (int ,error){
		if mode ==0{
			return tape[loc],nil
		}else if mode ==1{
			return loc,nil
		}else{
			return 0, UnknownModeError{}
		}
	}


	op,mode := splitOp(tape[index])
	switch op {
	case 99:
		return fmt.Sprintf("Halt\n"), nil
	case 1:
		loc1 := tape[index+1]
		loc2 := tape[index+2]
		loc3 := tape[index+3]
		val1,err:=getValue(loc1,mode[2])
		if err!=nil {
			return "",err
		}
		val2,err:=getValue(loc2,mode[1])
		if err!=nil {
			return "",err
		}
		return fmt.Sprintf("Add:%d + %d = %d written to %d\n", val1, val2, val1+val2, loc3), nil
	case 2:
		loc1 := tape[index+1]
		loc2 := tape[index+2]
		loc3 := tape[index+3]
		val1,err:=getValue(loc1,mode[2])
		if err!=nil {
			return "",err
		}
		val2,err:=getValue(loc2,mode[1])
		if err!=nil {
			return "",err
		}
		return fmt.Sprintf("Mul:%d + %d = %d written to %d\n", val1, val2, val1*val2, loc3), nil
	case 3:
		loc1 := tape[index+1]
		return fmt.Sprintf("Input:%d written to %d\n",input[0],loc1 ), nil
	case 4:
		loc1 := tape[index+1]
		val1,err:=getValue(loc1,mode[2])
		if err!=nil {
			return "",err
		}
		var s string
		if mode[2]==1{
			s="Imm"
		} else{
			s=strconv.Itoa(loc1)

		}
		return fmt.Sprintf("Output:%d from %v Written to output\n",val1,s), nil
	case 5:
		loc1 := tape[index+1]
		loc2 := tape[index+2]
		val1,err:=getValue(loc1,mode[2])
		if err!=nil {
			return "",err
		}
		val2,err:=getValue(loc2,mode[1])
		if err!=nil {
			return "",err
		}
		var s string
		if val1 !=0{
			s=fmt.Sprintf("Set to %d",val2)
		}else{
			s=fmt.Sprintf("Set to %d",index+3)
		}
		return fmt.Sprintf("JumpTrue:%v Inst -> %v\n", val1==1, s), nil
	case 6:
		loc1 := tape[index+1]
		loc2 := tape[index+2]
		val1,err:=getValue(loc1,mode[2])
		if err!=nil {
			return "",err
		}
		val2,err:=getValue(loc2,mode[1])
		if err!=nil {
			return "",err
		}
		var s string
		if val1 ==0{
			s=fmt.Sprintf("Set to %d",val2)
		}else{
			s=fmt.Sprintf("Set to %d",index+3)
		}
		return fmt.Sprintf("JumpTrue:%v Inst -> %v\n", val1==1, s), nil
	case 7:
		loc1 := tape[index+1]
		loc2 := tape[index+2]
		loc3 := tape[index+3]
		val1,err:=getValue(loc1,mode[2])
		if err!=nil {
			return "",err
		}
		val2,err:=getValue(loc2,mode[1])
		if err!=nil {
			return "",err
		}
		return fmt.Sprintf("Less:%d < %d = %v written to %d\n", val1, val2, val1<val2, loc3), nil
	case 8:
		loc1 := tape[index+1]
		loc2 := tape[index+2]
		loc3 := tape[index+3]
		val1,err:=getValue(loc1,mode[2])
		if err!=nil {
			return "",err
		}
		val2,err:=getValue(loc2,mode[1])
		if err!=nil {
			return "",err
		}
		return fmt.Sprintf("Equals:%d == %d = %v written to %d\n", val1, val2, val1==val2, loc3), nil
	default:
		return string(""), fmt.Errorf("Unknown Operation %d", op)
	}

}

// runTape Runs the current tape
func runTape(tape, input,output*[]int) (int, error) {
	finished := false
	for i := 0; finished != true;{


		s,err:=printOp(i,*tape,*input)
		fmt.Println(s)
		fmt.Printf("return:%d\n",(*tape)[0])
		if err!=nil{
			return 0,fmt.Errorf("Invalid Operand at %d\n Tape:%v %w",i,s,err)
		}
		result,err := doOp(&i, tape,input,output)
		if result!=nil{
			fmt.Println(output)
			return *result, nil
		}
		if err!=nil{
			if errors.Is(err,HaltError{}){
				finished=true
			}else {
				return 0, fmt.Errorf("Failed at %d\n Tape:%v %w", i, tape, err)
			}
		}

	}
	return 0,fmt.Errorf("Unexpected Termination")
}
func main() {
	fmt.Printf("hello, world\n")
	data, err := ioutil.ReadFile("programs2.txt")
	if err != nil {
		fmt.Printf("File Error %w", err)
		os.Exit(4)
	}

	program := strings.Split(strings.TrimSpace(string(data)), ",")
	tape := make([]int, len(program))
	for i, s := range program {
		tape[i], err = strconv.Atoi(s)
		if err != nil {
			fmt.Printf("Strconv Error %v", err)
			os.Exit(5)
		}
	}
	_, err = runTape(&tape,&[]int{5},&[]int{})
	if err != nil {
		fmt.Printf("Tape Error %v", err)
		os.Exit(1)
	}
}
