package repository

import (
	"goproject_port/datastruct"
	"reflect"
	"testing"
)

func TestGetIn(t *testing.T) {

	testsSet1 := []struct {
		name   string
		arg    int
		want   int
		indata []*datastruct.Port
	}{
		{
			name:   "test1",
			arg:    0,
			want:   9,
			indata: []*datastruct.Port{&datastruct.Port{Value: 9}},
		},
		{
			name:   "test2",
			arg:    1,
			want:   10,
			indata: []*datastruct.Port{&datastruct.Port{Value: 3}, &datastruct.Port{Value: 10}},
		},
		{
			name:   "test3",
			arg:    0,
			want:   6,
			indata: []*datastruct.Port{&datastruct.Port{Value: 6}},
		},
	}

	for _, tt := range testsSet1 {
		t.Run(tt.name, func(t *testing.T) {
			r := &repo{In: tt.indata}
			got, _ := r.GetIn(tt.arg)
			if got != tt.want {
				t.Errorf("Find() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostOut(t *testing.T) {

	testsSet2 := []struct {
		name    string
		arg1    int
		arg2    int
		want    []*datastruct.Port
		outdata []*datastruct.Port
	}{
		{
			name:    "test4",
			arg1:    0,
			arg2:    39,
			want:    []*datastruct.Port{&datastruct.Port{Value: 39}},
			outdata: []*datastruct.Port{&datastruct.Port{Value: 1}},
		},
		{
			name:    "test5",
			arg1:    1,
			arg2:    40,
			want:    []*datastruct.Port{&datastruct.Port{Value: 1}, &datastruct.Port{Value: 40}},
			outdata: []*datastruct.Port{&datastruct.Port{Value: 1}, &datastruct.Port{Value: 2}},
		},
		{
			name:    "test6",
			arg1:    1,
			arg2:    60,
			want:    []*datastruct.Port{&datastruct.Port{Value: 2}, &datastruct.Port{Value: 60}},
			outdata: []*datastruct.Port{&datastruct.Port{Value: 2}, &datastruct.Port{Value: 2}},
		},
	}

	for _, tt := range testsSet2 {
		t.Run(tt.name, func(t *testing.T) {
			r := &repo{Out: tt.outdata}
			_ = r.PostOut(tt.arg1, tt.arg2)
			if !reflect.DeepEqual(r.Out, tt.want) {
				t.Errorf("Find() got = %v, want %v", r.Out, tt.want)
			}
		})
	}
}

func TestGetChIn(t *testing.T) {

	ch_0 := make(chan datastruct.Reqest)
	ch_1 := make(chan datastruct.Reqest)
	ch_2 := make(chan datastruct.Reqest)

	testsSet3 := []struct {
		name   string
		arg    int
		want   chan datastruct.Reqest
		indata []*datastruct.Port
	}{
		{
			name:   "test7",
			arg:    0,
			want:   ch_0,
			indata: []*datastruct.Port{&datastruct.Port{Ch: ch_0}},
		},
		{
			name:   "test8",
			arg:    1,
			want:   ch_1,
			indata: []*datastruct.Port{&datastruct.Port{Value: 9, Ch: ch_0}, &datastruct.Port{Value: 10, Ch: ch_1}},
		},
		{
			name:   "test9",
			arg:    0,
			want:   ch_2,
			indata: []*datastruct.Port{&datastruct.Port{Ch: ch_2}},
		},
	}

	for _, tt := range testsSet3 {
		t.Run(tt.name, func(t *testing.T) {
			r := &repo{In: tt.indata}
			got, _ := r.GetChIn(tt.arg)
			if got != tt.want {
				t.Errorf("Find() got = %v, want %v", got, tt.want)
			}
		})
	}

}
