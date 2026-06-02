package runner_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/cmontemuino/tokeval/internal/runner"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    string
		wantErr bool
	}{
		{
			name:    "simple echo",
			args:    []string{"echo", "hello", "world"},
			want:    "hello world\n",
			wantErr: false,
		},
		{
			name:    "invalid command",
			args:    []string{"nonexisting-command-12345"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "empty args",
			args:    []string{},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := runner.Run(context.Background(), tt.args, &buf)

			if (err != nil) != tt.wantErr {
				t.Fatalf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				got := buf.String()
				if got != tt.want {
					t.Errorf("Run() got = %q, want %q", got, tt.want)
				}
			}
		})
	}
}
