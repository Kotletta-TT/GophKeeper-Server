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

type MockListRepository struct {
	mock.Mock
}

func (m *MockListRepository) ListSecretCardByUserId(ctx context.Context, userId uuid.UUID) ([]*entity.SecretCard, error) {
	args := m.Called(userId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.SecretCard), args.Error(1)
}

func (m *MockListRepository) ListSecretCardByName(ctx context.Context, name string) ([]*entity.SecretCard, error) {
	args := m.Called(name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.SecretCard), args.Error(1)
}

func TestListUC_ListSecretByUserId(t *testing.T) {
	l, err := logger.NewLogger(&config.Config{})
	assert.NoError(t, err)
	type args struct {
		userId uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		want    []*entity.SecretCard
		wantErr bool
	}{
		{
			name: "Return normal list",
			args: args{
				userId: uuid.New(),
			},
			want:    []*entity.SecretCard{},
			wantErr: false,
		},
		{
			name: "Return error",
			args: args{
				userId: uuid.New(),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &MockListRepository{}
			if tt.wantErr {
				r.On("ListSecretCardByUserId", tt.args.userId).Return(nil, assert.AnError)
			} else {
				r.On("ListSecretCardByUserId", tt.args.userId).Return(tt.want, nil)
			}
			uc := NewListUC(l, r)
			data, err := uc.ListSecretByUserId(context.Background(), tt.args.userId)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, data)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, data)
			}
		})
	}
}

func TestListUC_ListSecretByName(t *testing.T) {
	l, err := logger.NewLogger(&config.Config{})
	assert.NoError(t, err)
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    []*entity.SecretCard
		wantErr bool
	}{
		{
			name: "Return normal list",
			args: args{
				name: "someName",
			},
			want:    []*entity.SecretCard{},
			wantErr: false,
		},
		{
			name: "Return error",
			args: args{
				name: "someName",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &MockListRepository{}
			if tt.wantErr {
				r.On("ListSecretCardByName", tt.args.name).Return(nil, assert.AnError)
			} else {
				r.On("ListSecretCardByName", tt.args.name).Return(tt.want, nil)
			}
			uc := NewListUC(l, r)
			data, err := uc.ListSecretByName(context.Background(), tt.args.name)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, data)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, data)
			}
		})
	}
}

func TestNewListUC(t *testing.T) {
	l, err := logger.NewLogger(&config.Config{})
	assert.NoError(t, err)
	r := &MockListRepository{}
	uc := NewListUC(l, r)
	assert.NotNil(t, uc)
}
