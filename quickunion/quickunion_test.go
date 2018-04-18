package main

import "testing"

func Test_quickunion_root(t *testing.T) {
	type fields struct {
		data []int
	}
	type args struct {
		p int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{name: "simple test", fields: fields{[]int{0, 1, 2}}, args: args{1}, want: 1},
		{name: "simple test", fields: fields{[]int{1, 1, 1, 2}}, args: args{2}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &quickunion{
				data: tt.fields.data,
			}
			if got := u.root(tt.args.p); got != tt.want {
				t.Errorf("quickunion.root() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_quickunion_connected(t *testing.T) {
	type fields struct {
		data []int
	}
	type args struct {
		p int
		q int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{name: "simple", fields: fields{[]int{0, 1, 2}}, args: args{1, 2}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &quickunion{
				data: tt.fields.data,
			}
			if got := u.connected(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("quickunion.connected() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_quickunion_union(t *testing.T) {
	type fields struct {
		data []int
	}
	type args struct {
		p int
		q int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "simple", fields: fields{[]int{0, 1, 2}}, args: args{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &quickunion{
				data: tt.fields.data,
			}
			u.union(tt.args.p, tt.args.q)
			if !u.connected(tt.args.p, tt.args.q) {
				t.Errorf("quickunion.union() = %v, and %v not connected", tt.args.p, tt.args.q)
			}
		})
	}
}
