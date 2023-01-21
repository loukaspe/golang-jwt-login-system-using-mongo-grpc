package app

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/loukaspe/auth/mongo-handler/domain"
	"github.com/loukaspe/auth/mongo-handler/mocks/mock_domain"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

var result *domain.User

func BenchmarkUserService_Login(b *testing.B) {
	var actual *domain.User
	ctrl := gomock.NewController(b)

	const username = "username"
	const correctPassword = "password"
	encryptedPasswordBytes, _ := bcrypt.GenerateFromPassword([]byte(correctPassword), 14)
	encryptedPassword := string(encryptedPasswordBytes)

	mockUserDb := mock_domain.NewMockUserDBInterface(ctrl)

	mockUserDbResponse := &domain.User{
		Username: "username",
		Password: encryptedPassword,
	}

	var mockUserDbResponseError error

	mockUserDb.EXPECT().
		GetUser(bson.M{"username": username}).
		Return(mockUserDbResponse, mockUserDbResponseError)

	userService := UserService{
		userDb: mockUserDb,
	}

	for n := 0; n < b.N; n++ {
		actual, _ = userService.Login(username, correctPassword)
	}

	result = actual
}

func TestUserService_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	info := "yupiii"

	type args struct {
		user *domain.User
	}
	tests := []struct {
		name                        string
		args                        args
		mockUserDbResponse          error
		expectedErrorFromValidation bool
		expected                    error
	}{
		{name: "success",
			args: args{
				user: &domain.User{
					Username: "username",
					Password: "test",
					Info:     info,
				},
			},
			expectedErrorFromValidation: false,
			expected:                    nil,
		},
		{name: "user validation-missing username error",
			args: args{
				user: &domain.User{
					Password: "test",
					Info:     info,
				},
			},
			expectedErrorFromValidation: true,
			expected:                    errors.New("required username"),
		},
		{name: "db error",
			args: args{
				user: &domain.User{
					Username: "username",
					Password: "test",
					Info:     info,
				},
			},
			expectedErrorFromValidation: false,
			mockUserDbResponse:          errors.New("erroraraaaaa"),
			expected:                    errors.New("erroraraaaaa"),
		},
	}

	mockUserDb := mock_domain.NewMockUserDBInterface(ctrl)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.expectedErrorFromValidation {
				mockUserDb.EXPECT().CreateUser(tt.args.user).Return(tt.mockUserDbResponse)
			}

			userService := UserService{
				userDb: mockUserDb,
			}

			actual := userService.CreateUser(tt.args.user)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestUserService_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)

	type args struct {
		id string
	}
	tests := []struct {
		name               string
		args               args
		mockUserDbResponse error
		expected           error
	}{
		{name: "success",
			args: args{
				id: "aa",
			},
			expected: nil,
		},
		{name: "db error",
			args: args{
				id: "",
			},
			mockUserDbResponse: errors.New("erroraraaaaa"),
			expected:           errors.New("erroraraaaaa"),
		},
	}

	mockUserDb := mock_domain.NewMockUserDBInterface(ctrl)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserDb.EXPECT().DeleteUser(bson.M{"_id": tt.args.id}).Return(tt.mockUserDbResponse)

			userService := UserService{
				userDb: mockUserDb,
			}

			actual := userService.DeleteUser(tt.args.id)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestUserService_GetUser(t *testing.T) {
	ctrl := gomock.NewController(t)

	type args struct {
		id string
	}
	tests := []struct {
		name                    string
		args                    args
		mockUserDbResponse      *domain.User
		mockUserDbResponseError error
		expected                *domain.User
		expectedError           error
	}{
		{name: "success",
			args: args{
				id: "aa",
			},
			mockUserDbResponse: &domain.User{
				Username: "oh captain my captain",
			},
			mockUserDbResponseError: nil,
			expected: &domain.User{
				Username: "oh captain my captain",
			},
			expectedError: nil,
		},
		{name: "db error",
			args: args{
				id: "aa",
			},
			mockUserDbResponse:      &domain.User{},
			mockUserDbResponseError: errors.New("abouuuu"),
			expected:                &domain.User{},
			expectedError:           errors.New("abouuuu"),
		},
	}

	mockUserDb := mock_domain.NewMockUserDBInterface(ctrl)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserDb.EXPECT().
				GetUser(bson.M{"_id": tt.args.id}).
				Return(tt.mockUserDbResponse, tt.mockUserDbResponseError)

			userService := UserService{
				userDb: mockUserDb,
			}

			actual, actualError := userService.GetUser(tt.args.id)

			assert.Equal(t, tt.expected, actual)
			assert.Equal(t, tt.expectedError, actualError)
		})
	}
}

func TestUserService_Login(t *testing.T) {
	ctrl := gomock.NewController(t)

	const correctPassword = "password"
	const wrongPassword = "wrongPassword"

	encryptedPasswordBytes, _ := bcrypt.GenerateFromPassword([]byte(correctPassword), 14)
	encryptedPassword := string(encryptedPasswordBytes)

	type args struct {
		username string
		password string
	}
	tests := []struct {
		name                    string
		args                    args
		mockUserDbResponse      *domain.User
		mockUserDbResponseError error
		expected                *domain.User
		expectedError           error
	}{
		{name: "success",
			args: args{
				username: "username",
				password: correctPassword,
			},
			mockUserDbResponse: &domain.User{
				Username: "username",
				Password: encryptedPassword,
			},
			mockUserDbResponseError: nil,
			expected: &domain.User{
				Username: "username",
				Password: encryptedPassword,
			},
		},
		{name: "no user found",
			args: args{
				username: "username",
				password: correctPassword,
			},
			mockUserDbResponse:      &domain.User{},
			mockUserDbResponseError: errors.New("no user found"),
			expected:                &domain.User{},
			expectedError:           errors.New("no user found"),
		},
		{name: "wrong password",
			args: args{
				username: "username",
				password: wrongPassword,
			},
			mockUserDbResponse: &domain.User{
				Username: "username",
				Password: encryptedPassword,
			},
			mockUserDbResponseError: nil,
			expected:                &domain.User{},
			expectedError:           bcrypt.ErrMismatchedHashAndPassword,
		},
	}
	mockUserDb := mock_domain.NewMockUserDBInterface(ctrl)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserDb.EXPECT().
				GetUser(bson.M{"username": tt.args.username}).
				Return(tt.mockUserDbResponse, tt.mockUserDbResponseError)

			userService := UserService{
				userDb: mockUserDb,
			}

			actual, actualError := userService.Login(tt.args.username, tt.args.password)

			assert.Equal(t, tt.expected, actual)
			assert.Equal(t, tt.expectedError, actualError)
		})
	}
}

func TestUserService_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)

	userId, _ := primitive.ObjectIDFromHex("507f1f77bcf86cd799439011")
	info := "yupiii"

	type args struct {
		user *domain.User
	}
	tests := []struct {
		name                        string
		args                        args
		mockUserDbResponse          error
		expectedErrorFromValidation bool
		expected                    error
	}{
		{name: "success",
			args: args{
				user: &domain.User{
					ID:       userId,
					Username: "username",
					Password: "test",
					Info:     info,
				},
			},
			expectedErrorFromValidation: false,
			expected:                    nil,
		},
		{name: "user validation-missing username error",
			args: args{
				user: &domain.User{
					ID:       userId,
					Password: "test",
					Info:     info,
				},
			},
			expectedErrorFromValidation: true,
			expected:                    errors.New("required username"),
		},
		{name: "db error",
			args: args{
				user: &domain.User{
					ID:       userId,
					Username: "username",
					Password: "test",
					Info:     info,
				},
			},
			expectedErrorFromValidation: false,
			mockUserDbResponse:          errors.New("erroraraaaaa"),
			expected:                    errors.New("erroraraaaaa"),
		},
	}

	mockUserDb := mock_domain.NewMockUserDBInterface(ctrl)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.expectedErrorFromValidation {
				mockUserDb.EXPECT().CreateUser(tt.args.user).Return(tt.mockUserDbResponse)
			}

			userService := UserService{
				userDb: mockUserDb,
			}

			actual := userService.CreateUser(tt.args.user)

			assert.Equal(t, tt.expected, actual)
		})
	}
}
