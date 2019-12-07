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
		fmt.Println(input)
		loc1 := (*tape)[*index+1]
		if len(*input) > 0 {
			(*tape)[loc1] = (*input)[0]
			*input=(*input)[1:]
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
		if len(input )>0{
			return fmt.Sprintf("Input:%d written to %d\n",input[0],loc1 ), nil
		}
		return fmt.Sprintf("Input already consumed\n"),EmptyInputError{}
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
		return fmt.Sprintf("Output:%d from %v Written to output\n Halt at next Instruction:%v\n",val1,s,tape[index+2]==99), nil
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
		return fmt.Sprintf("JumpTrue: %v!=0 :%v Inst -> %v\n",val1, val1!=0, s), nil
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
		return fmt.Sprintf("JumpZero: %v==0:%v Inst -> %v\n", val1,val1==0, s), nil
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
		if err!=nil{
			return 0,fmt.Errorf("Invalid Operand at %d\n Status:%v\n Tape:%v %w",i,s,*tape,err)
		}
		result,err := doOp(&i, tape,input,output)
		if result!=nil{
			return *result, nil
		}else if err!=nil{
			if !errors.Is(err,HaltError{}){
				return 0, fmt.Errorf("Failed at %d\n Tape:%v %w", i, tape, err)
			}
		}

	}
	return 0,fmt.Errorf("Unexpected Termination")
}
func amplify(program, phases[]int)(int,error){
	in_out:=0
	var err error
	for i,v:= range phases{
		working:=make([]int,len(program))
		copy(working,program)
		input:=[]int{v,in_out}
		output:=[]int{}
		_,err=runTape(&working,&input,&output)
		in_out=output[0]
		if err!=nil {
			return in_out,fmt.Errorf("Amplify:Failed at stage %v:%w",i,err)
		}
		fmt.Println(working)
	}
	return in_out,nil
}
func phases(n int) [][]int{
	start:=make([]int,n)
	output:=make([][]int,0)
	for i:=0;i<n;i++{
		start[i]=i
	}
	var permutations func(k int, A *[]int)
	permutations=func(k int, A *[]int){
		if k ==1 {
		o:=make([]int,n)
		copy(o,start)
		output=append(output,o)
		}else{
			permutations(k-1,A)
			for i:=0; i<k-1;i++{
				if (k & 1)==0{
					(*A)[i],(*A)[k-1]=(*A)[k-1],(*A)[i]
				}else {
					(*A)[0],(*A)[k-1]=(*A)[k-1],(*A)[0]
				}
				permutations(k-1,A)

			}
		}
	}
	permutations(n,&start)
	return output
}
func main() {
	data, err := ioutil.ReadFile("programs3.txt")
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

	tape=[]int{3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0}
	output :=0
	var phase []int
	ps:=phases(5)
	for _,v := range ps{
		t, err := amplify(tape,v)
		if err != nil {
			fmt.Printf("Amplify Error %v", err)
			os.Exit(1)
		}
		if t>output{
			output=t
			phase=v
		}

	}
	fmt.Printf("Highest output:%v,Phase:%v",output,phase)
}
