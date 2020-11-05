package count

import "testing"

func TestDo(t *testing.T) {
	type args struct {
		something interface{}
		target    interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// --- positive testing ---
		{
			name: "positive testing(string)",
			args: args{
				something: "something",
				target:    "a",
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "positive testing(string)",
			args: args{
				something: "something",
				target:    "s",
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "positive testing([]string)",
			args: args{
				something: []string{"hoge", "foo", "baz"},
				target:    "foo",
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "positive testing([]string)",
			args: args{
				something: []string{"hoge", "foo", "foo"},
				target:    "foo",
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "positive testing([]int)",
			args: args{
				something: []int{1, 2, 3, 4, 5},
				target:    2,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "positive testing([]int)",
			args: args{
				something: []int{1, 2, 3, 4, 4},
				target:    4,
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "positive testing([]interface)",
			args: args{
				something: []interface{}{1, "2", 3, "4", nil},
				target:    nil,
			},
			want:    1,
			wantErr: false,
		},
		// --- negative testing ---
		{
			name: "negative testing(unsupported type:nil)",
			args: args{
				something: nil,
				target:    nil,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "negative testing(unsupported type:map)",
			args: args{
				something: map[string]string{},
				target:    "test",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "negative testing(unsupported type:error)",
			args: args{
				something: []interface{}{1, "2", 3, "4", []int{1, 2, 3}},
				target:    []int{1, 2, 3},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt // Escape: Using the variable on range scope `tt` in function literal
		t.Run(tt.name, func(t *testing.T) {
			got, err := Do(tt.args.something, tt.args.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Do() got = %v, want %v", got, tt.want)
			}
		})
	}
}
