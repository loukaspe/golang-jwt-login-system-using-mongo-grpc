package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthMech_CreateTokenAndGetClaimsFromToken(t *testing.T) {
	type args struct {
		subject  string
		userInfo interface{}
	}
	tests := []struct {
		name                 string
		args                 args
		expectedClaim        string
		expectError          bool
		expectedErrorMessage string
	}{
		{name: "success",
			args: args{
				subject:  "subject",
				userInfo: "allTheSingleLadiesAllTheSingleLadies",
			},
			expectedClaim:        "allTheSingleLadiesAllTheSingleLadies",
			expectError:          false,
			expectedErrorMessage: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &AuthMech{}
			token, err := j.CreateToken(tt.args.subject, tt.args.userInfo)
			if tt.expectError {
				assert.Equal(t, tt.expectedErrorMessage, err.Error())
			}

			claims, err := j.GetClaimsFromToken(token)
			actualClaim, ok := claims["UserInfo"].(string)
			if !ok {
				t.Error("no user info in claims")
			}

			assert.Equal(t, tt.expectedClaim, actualClaim)
		})
	}
}
