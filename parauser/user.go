package parauser

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func (c *Client) LoginUser(ctx context.Context, req *LoginUserRequest) (*User, error) {
	passwordHash := md5.Sum([]byte(req.Password))
	res, err := c.client.R().
		SetFormData(map[string]string{
			"token_type":  "TOKEN",
			"third_party": "SELF",
			"email":       req.Email,
			"password":    hex.EncodeToString(passwordHash[:]),
		}).
		SetResult(&User{}).
		Post("/user/api/login")
	if err != nil {
		return nil, fmt.Errorf("http post: %v", err)
	}
	if res.IsError() {
		svcErr, ok := res.Error().(*ServiceError)
		if !ok {
			return nil, fmt.Errorf("http post response: %s", res.Body())
		}
		return nil, fmt.Errorf("http post response: %s", svcErr.Message)
	}
	user := res.Result().(*User)
	token := res.Header().Get("para_token")
	if token == "" {
		return nil, fmt.Errorf("http post response: empty http header 'para_token'")
	}
	user.Token = token
	return user, nil
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID    string `json:"userId"`
	Email string `json:"email"`
	Token string `json:"-"`
}
