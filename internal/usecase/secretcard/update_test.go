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

type MockUpdateRepository struct {
	mock.Mock
}

func (m *MockUpdateRepository) UpdateSecretCard(ctx context.Context, card *entity.SecretCard) error {
	args := m.Called(card)
	return args.Error(0)
}

func TestUpdateUC_UpdateSecret(t *testing.T) {
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
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &MockUpdateRepository{}
			uc := &UpdateUC{
				l: l,
				r: r,
			}
			r.On("UpdateSecretCard", tt.args.card).Return(tt.err)
			err := uc.UpdateSecret(context.Background(), tt.args.card)
			if tt.wantErr {
				assert.Equal(t, tt.err, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestNewUpdateUC(t *testing.T) {
	l, err := logger.NewLogger(&config.Config{})
	r := &MockUpdateRepository{}
	assert.NoError(t, err)
	wantUC := &UpdateUC{l: l, r: r}
	uc := NewUpdateUC(l, r)
	assert.Equal(t, wantUC, uc)
}
