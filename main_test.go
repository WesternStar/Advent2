package main

import (
	"reflect"
	"testing"
)

func Test_doOp(t *testing.T) {

	type args struct {
		index  int
		tape   *[]int
		input  *[]int
		output *[]int
	}
	tests := []struct {
		name     string
		args     args
		want     *int
		wantErr  bool
		wantTape *[]int
	}{
		{"Return", args{4, &[]int{100, 2, 3, 5, 99}, &[]int{}, &[]int{}}, new(int), true, &[]int{100, 2, 3, 5, 99}},
		{"Add", args{0, &[]int{1, 2, 3, 5, 99, 50, 67, 72}, &[]int{}, &[]int{}}, nil, false, &[]int{1, 2, 3, 5, 99, 8, 67, 72}},
		{"Multiply", args{0, &[]int{2, 2, 3, 5, 99, 50, 67, 72}, &[]int{}, &[]int{}}, nil, false, &[]int{2, 2, 3, 5, 99, 15, 67, 72}},
	}
	*(tests[0].want) = 100
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := doOp(&tt.args.index, tt.args.tape, tt.args.input, tt.args.output)
			if (err != nil) != tt.wantErr {
				t.Errorf("doOp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doOp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printOp(t *testing.T) {
	type args struct {
		index int
		tape  []int
		input []int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := printOp(tt.args.index, tt.args.tape, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("printOp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("printOp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_runTape(t *testing.T) {
	type args struct {
		tape   *[]int
		input  *[]int
		output *[]int
	}
	tests := []struct {
		name       string
		args       args
		want       int
		wantErr    bool
		wantTape   *[]int
		wantOutput *[]int
	}{
		{"Advent Test1", args{&[]int{1, 0, 0, 0, 99}, &[]int{}, &[]int{}}, 2, false, &[]int{2, 0, 0, 0, 99},&[]int{}},
		{"Advent Test2", args{&[]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}, &[]int{}, &[]int{}}, 3500, false, &[]int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},&[]int{}},
		{"Advent Test3", args{&[]int{2, 3, 0, 3, 99}, &[]int{}, &[]int{}}, 2, false, &[]int{2, 3, 0, 6, 99},&[]int{}},
		{"Advent Test4", args{&[]int{2, 4, 4, 5, 99, 0}, &[]int{}, &[]int{}}, 2, false, &[]int{2, 4, 4, 5, 99, 9801},&[]int{}},
		{"Advent Test5", args{&[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, &[]int{}, &[]int{}}, 30, false, &[]int{30, 1, 1, 4, 2, 5, 6, 0, 99},&[]int{}},
		{"Advent Test Unknown", args{&[]int{100, 2, 3, 5, 99}, &[]int{}, &[]int{}}, 0, true, &[]int{100, 2, 3, 5, 99},&[]int{}},
		{"Input Test Equal1", args{&[]int{3,9,8,9,10,9,4,9,99,-1,8}, &[]int{8}, &[]int{}}, 3, false, &[]int{3,9,8,9,10,9,4,9,99,1,8},&[]int{1}},
		{"Input Test Equal2", args{&[]int{3,9,8,9,10,9,4,9,99,-1,8}, &[]int{100}, &[]int{}}, 3, false, &[]int{3,9,8,9,10,9,4,9,99,0,8},&[]int{0}},
		{"Input Test Less Than 1", args{&[]int{3,9,7,9,10,9,4,9,99,-1,8}, &[]int{7}, &[]int{}}, 3, false, &[]int{3,9,7,9,10,9,4,9,99,1,8},&[]int{1}},
		{"Input Test Less Than 2", args{&[]int{3,9,7,9,10,9,4,9,99,-1,8}, &[]int{100}, &[]int{}}, 3, false, &[]int{3,9,7,9,10,9,4,9,99,0,8},&[]int{0}},
		{"Input Test Imm Equal1", args{&[]int{3,3,1108,-1,8,3,4,3,99}, &[]int{8}, &[]int{}}, 3, false, &[]int{3,3,1108,1,8,3,4,3,99},&[]int{1}},
		{"Input Test Imm Equal2", args{&[]int{3,3,1108,-1,8,3,4,3,99}, &[]int{100}, &[]int{}}, 3, false, &[]int{3,3,1108,0,8,3,4,3,99},&[]int{0}},
		{"Input Test Imm Less Than 1", args{&[]int{3,3,1107,-1,8,3,4,3,99}, &[]int{7}, &[]int{}}, 3, false, &[]int{3,3,1107,1,8,3,4,3,99},&[]int{1}},
		{"Input Test Imm Less Than 2", args{&[]int{3,3,1107,-1,8,3,4,3,99}, &[]int{100}, &[]int{}}, 3, false, &[]int{3,3,1107,0,8,3,4,3,99},&[]int{0}},

		{"Xor ", args{&[]int{3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9}, &[]int{100}, &[]int{}}, 3, false, &[]int{3,12,6,12,15,1,13,14,13,4,13,99,100,1,1,9},&[]int{1}},
		{"Xor 2 ", args{&[]int{3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9}, &[]int{0}, &[]int{}}, 3, false, &[]int{3,12,6,12,15,1,13,14,13,4,13,99,0,0,1,9},&[]int{0}},
		{"Xor Imm ", args{&[]int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, &[]int{100}, &[]int{}}, 3, false, &[]int{3, 3, 1105, 100, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, &[]int{1}},
		{"Xor Imm 2", args{&[]int{3,3,1105,-1,9,1101,0,0,12,4,12,99,1}, &[]int{0}, &[]int{}}, 3, false, &[]int{3,3,1105,0,9,1101,0,0,12,4,12,99,0},&[]int{0}},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := runTape(tt.args.tape, tt.args.input, tt.args.output)
			if (err != nil) != tt.wantErr {
				t.Errorf("runTape() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("runTape() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(tt.args.tape, tt.wantTape) {
				t.Errorf("doOp() = %v, want %v", tt.args.tape, tt.wantTape)
			}
			if !reflect.DeepEqual(tt.args.output, tt.wantOutput) {
				t.Errorf("doOp() = %v, want %v", tt.args.output, tt.wantOutput)
			}
		})
	}
}

func Test_amplify(t *testing.T) {
	type args struct {
		program []int
		phases  []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"Test1",args{[]int{3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0},[]int{4,3,2,1,0}},43210,false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := amplify(tt.args.program, tt.args.phases)
			if (err != nil) != tt.wantErr {
				t.Errorf("amplify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("amplify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_phases(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1,2,3",args{3},[][]int{[]int{1,2,3},[]int{2,1,3},[]int{3,1,2},[]int{1,3,2},[]int{2,3,1},[]int{3,2,1},}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := phases(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("phases() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
