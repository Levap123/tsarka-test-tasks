package algos

import (
	"reflect"
	"testing"
)

func TestAlgosService_EmailCheck(t *testing.T) {
	s := &AlgosService{}

	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: `should return all email matching the pattern «Email:__$email»`,
			args: args{
				str: "Email: test@mail.ru hello heeeelloo abcd@g,ail.com , Email: qwer@qwerty.org",
			},
			want: []string{"test@mail.ru", "qwer@qwerty.org"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.EmailCheck(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AlgosService.EmailCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}
