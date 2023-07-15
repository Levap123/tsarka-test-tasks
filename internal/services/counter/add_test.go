package counter

import (
	"context"
	"testing"

	"github.com/Levap123/tsarka-test-tasks/mocks"
	"github.com/stretchr/testify/mock"
)

func TestCounterService_Add(t *testing.T) {
	counterRepo := mocks.NewCounterRepoI(t)
	counterRepo.EXPECT().Get(mock.Anything).Return(0, nil)
	counterRepo.EXPECT().Set(mock.Anything, mock.Anything).Return(nil)

	s := NewService(counterRepo)

	type args struct {
		ctx context.Context
		add int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ok",
			args: args{
				ctx: context.Background(),
				add: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.Add(tt.args.ctx, tt.args.add); err != nil {
				t.Errorf("expected not err, got: %v", err)
			}
		})
	}
}
