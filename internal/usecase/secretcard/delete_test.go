package secretcard

import (
	"GophKeeper-Server/config"
	"GophKeeper-Server/logger"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDeleteRepository struct {
	mock.Mock
}

func (m *MockDeleteRepository) DeleteSecretCard(secretId uuid.UUID) error {
	args := m.Called(secretId)
	return args.Error(0)
}

func TestDeleteUC_DeleteSecret(t *testing.T) {
	l, err := logger.NewLogger(&config.Config{})
	assert.NoError(t, err)
	type args struct {
		secretId uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Normal delete",
			args: args{
				secretId: uuid.New(),
			},
			wantErr: false,
		},
		{
			name: "Delete error",
			args: args{
				secretId: uuid.New(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &MockDeleteRepository{}
			if tt.wantErr {
				r.On("DeleteSecretCard", tt.args.secretId).Return(assert.AnError)
			} else {
				r.On("DeleteSecretCard", tt.args.secretId).Return(nil)
			}
			uc := NewDeleteUC(l, r)
			err := uc.DeleteSecret(tt.args.secretId)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
