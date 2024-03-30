package secretcard

import (
	"GophKeeper-Server/config"
	"GophKeeper-Server/internal/entity"
	"GophKeeper-Server/logger"
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockReadRepository struct {
	mock.Mock
}

func (m *MockReadRepository) GetSecretCardByID(ctx context.Context, secretId uuid.UUID) (*entity.SecretCard, error) {
	args := m.Called(secretId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.SecretCard), args.Error(1)
}

func TestNewReadUC(t *testing.T) {
	logger, err := logger.NewLogger(&config.Config{})
	assert.NoError(t, err)
	repo := &MockReadRepository{}
	uc := NewReadUC(logger, repo)
	assert.NotNil(t, uc)
}

func TestReadUC_ReadSecret(t *testing.T) {
	l, err := logger.NewLogger(&config.Config{})
	assert.NoError(t, err)
	type args struct {
		secretId uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		want    *entity.SecretCard
		wantErr bool
	}{
		{
			name:    "Get secret card",
			args:    args{secretId: uuid.MustParse("7db789d8-2fdc-4302-ab24-5eb898dd6ae4")},
			want:    &entity.SecretCard{ID: uuid.MustParse("7db789d8-2fdc-4302-ab24-5eb898dd6ae4")},
			wantErr: false,
		},
		{
			name:    "Get no exist secret card",
			args:    args{secretId: uuid.New()},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "Error",
			args:    args{secretId: uuid.New()},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &MockReadRepository{}
			if tt.wantErr {
				r.On("GetSecretCardByID", tt.args.secretId).Return(nil, assert.AnError)
			} else {
				r.On("GetSecretCardByID", tt.args.secretId).Return(tt.want, nil)
			}
			uc := NewReadUC(l, r)
			dbCard, err := uc.ReadSecret(context.Background(), tt.args.secretId)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, dbCard)
			}
		})
	}
}
