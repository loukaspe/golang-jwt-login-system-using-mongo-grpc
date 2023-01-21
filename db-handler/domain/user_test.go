package domain

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
	"time"
)

func TestUser_HashAndCheckPassword(t *testing.T) {
	type fields struct {
		Password string
	}
	type args struct {
		providedPassword string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected error
	}{
		{name: "success",
			fields: fields{
				Password: "ntolmadakia",
			},
			args: args{
				providedPassword: "ntolmadakia",
			},
			expected: nil,
		},
		{name: "different password",
			fields: fields{
				Password: "ntolmadakia",
			},
			args: args{
				providedPassword: "ntolmadakiaaaaaa",
			},
			expected: bcrypt.ErrMismatchedHashAndPassword,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &User{
				Password: tt.fields.Password,
			}
			err := user.HashPassword()
			if err != nil {
				t.Errorf("problem with testing, cannot hash code")
			}

			actual := user.CheckPassword(tt.args.providedPassword)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func FuzzUser_HashAndCheckPassword(f *testing.F) {
	testCases := []string{"password", " ", "!12345"}
	for _, tc := range testCases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, orig string) {
		user := &User{
			Password: orig,
		}

		err := user.HashPassword()
		if err != nil {
			t.Errorf("problem with testing, cannot hash code")
		}

		actualError := user.CheckPassword(orig)

		assert.Equal(t, nil, actualError)
	})
}

func TestUser_Validate(t *testing.T) {
	info := "yupii"

	type fields struct {
		Username   string
		Password   string
		info       string
		CreatedAt  time.Time
		ModifiedAt time.Time
	}
	tests := []struct {
		name     string
		fields   fields
		expected error
	}{
		{name: "success",
			fields: fields{
				Username: "username",
				Password: "pass",
				info:     info,
			},
			expected: nil,
		},
		{name: "missing username",
			fields: fields{
				Password: "pass",
				info:     info,
			},
			expected: errors.New("required username"),
		},
		{name: "missing password",
			fields: fields{
				Username: "username",
				info:     info,
			},
			expected: errors.New("required password"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &User{
				Username:   tt.fields.Username,
				Password:   tt.fields.Password,
				Info:       tt.fields.info,
				CreatedAt:  tt.fields.CreatedAt,
				ModifiedAt: tt.fields.ModifiedAt,
			}

			actual := user.Validate()
			assert.Equal(t, tt.expected, actual)
		})
	}
}
