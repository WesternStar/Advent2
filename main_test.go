package main

import "testing"
import "github.com/google/go-cmp/cmp"

func Test_doOp(t *testing.T) {
	resultRet:=100
	resultUnk:=101
	type args struct {
		index int
		tape  []int
	}
	tests := []struct {
		name         string
		args         args
		wantFinished bool
		wantResult   *int
		wantTape []int
	}{
		{"Return",args{4,[]int{100,2,3,5,99}},true,&resultRet,[]int{100,2,3,5,99}},
		{"Add",args{0,[]int{1,2,3,5,99,50,67,72}},false,nil,[]int{1,2,3,5,99,8,67,72}},
		{"Multiply",args{0,[]int{2,2,3,5,99,50,67,72}},false,nil,[]int{2,2,3,5,99,15,67,72}},
		{"Unknown Op",args{0,[]int{101,2,3,5,99}},false,&resultUnk,[]int{101,2,3,5,99}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFinished, gotResult := doOp(tt.args.index, tt.args.tape)
			if gotFinished != tt.wantFinished {
				t.Errorf("doOp() gotFinished = %v, want %v", gotFinished, tt.wantFinished)
			}
			if (gotResult ==nil || tt.wantResult==nil) && gotResult != tt.wantResult {
				t.Errorf("doOp() gotResult = %v, want %v", gotResult, tt.wantResult)
			}else if gotResult!=nil && *gotResult != *tt.wantResult {
				t.Errorf("doOp() gotResult = %v, want %v", *gotResult, *tt.wantResult)
			}
			if diff := cmp.Diff(tt.args.tape, tt.wantTape);diff  != ""{
				t.Errorf("doOp() Tape mismatch = %v",diff )
			}
		})
	}
}


func Test_runTape(t *testing.T) {
	type args struct {
		tape []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
		wantTape []int
	}{
		{"Advent Test1",args{[]int{1,0,0,0,99}},2,false,[]int{2,0,0,0,99}},
		{"Advent Test2",args{[]int{1,9,10,3,2,3,11,0,99,30,40,50}},3500,false,[]int{3500,9,10,70,2,3,11,0,99,30,40,50}},
		{"Advent Test3",args{[]int{2,3,0,3,99}},2,false,[]int{2,3,0,6,99}},
		{"Advent Test4",args{[]int{2,4,4,5,99,0}},2,false,[]int{2,4,4,5,99,9801}},
		{"Advent Test5",args{[]int{1,1,1,4,99,5,6,0,99}},30,false,[]int{30,1,1,4,2,5,6,0,99}},
		{"Advent Test Unknown",args{[]int{100,2,3,5,99}},0,true,[]int{100,2,3,5,99}},
	}
	for _, tt := range tests  {
		t.Run(tt.name, func(t *testing.T) {
			got, err := runTape(tt.args.tape)
			if (err != nil) != tt.wantErr {
				t.Errorf("runTape() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("runTape() = %v, want %v", got, tt.want)
			}
			if diff := cmp.Diff(tt.args.tape, tt.wantTape);diff  != ""{
				t.Errorf("doOp() Tape mismatch = %v",diff )
			}
		})
	}
}
