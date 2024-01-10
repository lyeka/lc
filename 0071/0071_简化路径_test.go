package simplifyPath

import (
	"testing"
)

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "0",
			args: args{path: "/home/"},
			want: "/home",
		},
		{
			name: "1",
			args: args{path: "/home//foo/"},
			want: "/home/foo",
		},
		{
			name: "2",
			args: args{path: "/a/./b/../../c/"},
			want: "/c",
		},
		{
			name: "3",
			args: args{path: "/../"},
			want: "/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
