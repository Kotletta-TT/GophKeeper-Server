package secretcard

import (
	"GophKeeper-Server/config"
	"GophKeeper-Server/internal/entity"
	"GophKeeper-Server/logger"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAddRepository struct {
	mock.Mock
}

func (m *MockAddRepository) CreateSecretCard(ctx context.Context, card *entity.SecretCard) error {
	args := m.Called(card)
	return args.Error(0)
}

func TestNewAddUC(t *testing.T) {
	l, err := logger.NewLogger(&config.Config{})
	if err != nil {
		assert.Error(t, err)
	}
	type args struct {
		l logger.Logger
		r AddRepository
	}
	tests := []struct {
		name string
		args args
		want *AddUC
	}{
		{
			name: "Create new AddUC",
			args: args{
				l: l,
				r: &MockAddRepository{},
			},
			want: &AddUC{
				l: l,
				r: &MockAddRepository{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewAddUC(tt.args.l, tt.args.r))
		})
	}
}

func TestAddUC_AddSecret(t *testing.T) {
	l, err := logger.NewLogger(&config.Config{})
	if err != nil {
		assert.Error(t, err)
	}
	type args struct {
		card *entity.SecretCard
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "Add secret card",
			args: args{
				card: &entity.SecretCard{},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Add secret card error",
			args: args{
				card: &entity.SecretCard{},
			},
			wantErr: true,
			err:     assert.AnError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &MockAddRepository{}
			uc := NewAddUC(l, r)
			uc.r.(*MockAddRepository).On("CreateSecretCard", tt.args.card).Return(tt.err)
			if tt.wantErr {
				assert.Error(t, uc.AddSecret(context.Background(), tt.args.card))
			} else {
				assert.NoError(t, uc.AddSecret(context.Background(), tt.args.card))
			}
		})
	}
}
