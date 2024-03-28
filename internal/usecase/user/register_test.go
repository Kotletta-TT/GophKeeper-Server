package user

import (
	"GophKeeper-Server/config"
	"GophKeeper-Server/internal/entity"
	customErrs "GophKeeper-Server/internal/errors"
	"GophKeeper-Server/logger"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRegisterRepository struct {
	mock.Mock
}

func (m *MockRegisterRepository) Register(login, password string) error {
	args := m.Called(login, password)
	return args.Error(0)
}

func (m *MockRegisterRepository) GetUser(login string) (*entity.User, error) {
	args := m.Called(login)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

func MockHashFunc(password string) (string, error) {
	return password, nil
}

func TestRegisterUC_Register(t *testing.T) {
	l, err := logger.NewLogger(&config.Config{})
	if err != nil {
		assert.Error(t, err)
	}
	type fields struct {
		l logger.Logger
		h HashFunc
	}
	type args struct {
		login      string
		password   string
		returnUser *entity.User
		dbErr      error
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		codeErr int
	}{
		{
			name: "Success create",
			fields: fields{
				l: l,
				h: MockHashFunc,
			},
			args: args{
				login:      "login",
				password:   "password",
				returnUser: nil,
				dbErr:      nil,
			},
			wantErr: false,
			codeErr: 0,
		},
		{
			name: "Error already exists",
			fields: fields{
				l: l,
				h: MockHashFunc,
			},
			args: args{
				login:      "login",
				password:   "password",
				returnUser: &entity.User{},
				dbErr:      nil,
			},
			wantErr: true,
			codeErr: 409,
		},
		{
			name: "Error database internal",
			fields: fields{
				l: l,
				h: MockHashFunc,
			},
			args: args{
				login:      "login",
				password:   "password",
				returnUser: nil,
				dbErr:      customErrs.ErrDatabaseInternal(fmt.Errorf("error")),
			},
			wantErr: true,
			codeErr: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := new(MockRegisterRepository)
			uc := NewRegisterUC(tt.fields.l, r, tt.fields.h)
			r.On("GetUser", tt.args.login).Return(tt.args.returnUser, tt.args.dbErr)
			r.On("Register", tt.args.login, tt.args.password).Return(tt.args.dbErr)

			err := uc.Register(tt.args.login, tt.args.password)
			if tt.wantErr {
				var ce *customErrs.CustomErrors
				if errors.As(err, &ce) {
					t.Log(ce.Code)
					assert.Equal(t, tt.codeErr, ce.Code)
				} else {
					assert.Error(t, err)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
