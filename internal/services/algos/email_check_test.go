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
			name: "should return all email matching the pattern «Email:__$email»",
			args: args{
				str: "Email: test@mail.ru hello heeeelloo abcd@g,ail.com, Email: qwer@qwerty.org Email: q1w3e4r5t5y@tre.com",
			},
			want: []string{"test@mail.ru", "qwer@qwerty.org", "q1w3e4r5t5y@tre.com"},
		},
		{
			name: "should return empty slice",
			args: args{
				str: "Email: abc@abc Email: @gmail.com Email: qwerty@ post@post.com test123@gmail.com",
			},
			want: []string{},
		},
		{
			name: "should work fine with any space symbols",
			args: args{
				str: "Email: asd asd@gmail.com TSARKA THE BEST Email: \n \r \t     qwe@rte.rte Email:    \ntest123@org.org Email:\t\n   abc@abc.abc",
			},
			want: []string{"qwe@rte.rte", "test123@org.org", "abc@abc.abc"},
		},
		{
			name: "should return only correct email adresses",
			args: args{
				str: "Email: abx@mail..ru, Email qwe@.., Email: correct@correct.com, Email:alotdomains@gmail.com.kz.org.uk Email: q..we@qwe Email: incorrect@@qwe.qwe",
			},
			want: []string{"correct@correct.com", "alotdomains@gmail.com.kz.org.uk"},
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
