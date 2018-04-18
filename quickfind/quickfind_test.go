package main

import "testing"

func Test_quickfind_connected(t *testing.T) {
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
			i := &quickfind{
				data: tt.fields.data,
			}
			if got := i.connected(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("quickfind.connected() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_quickfind_union(t *testing.T) {
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
			i := &quickfind{
				data: tt.fields.data,
			}
			i.union(tt.args.p, tt.args.q)
			if !i.connected(tt.args.p, tt.args.q) {
				t.Errorf("quickunion.union() = %v, and %v not connected", tt.args.p, tt.args.q)
			}
		})
	}
}
