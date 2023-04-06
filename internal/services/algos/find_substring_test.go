package algos

import "testing"

func TestAlgosService_FindSubstring(t *testing.T) {
	s := &AlgosService{}

	type args struct {
		str string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should return largest substring without repetitions",
			args: args{
				str: "aaaabbbcdeffffwqarrnan",
			},
			want: "bcdef",
		},
		{
			name: "should return empty string",
			args: args{
				str: "",
			},
			want: "",
		},
		{
			name: "should return largest substring without repetitions",
			args: args{
				str: "abcbafgeaterquiyo",
			},
			want: "aterquiyo",
		},
		{
			name: "should return only 1 caharacter",
			args: args{
				str: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			},
			want: "a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.FindSubstring(tt.args.str); got != tt.want {
				t.Errorf("AlgosService.FindSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
