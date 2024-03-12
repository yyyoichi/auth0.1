package resource

import (
	"context"
	"errors"
	"fmt"
	"time"

	apiv1 "github.com/yyyoichi/OhAuth0.1/api/v1"
)

type (
	Service struct {
		client clientInterface
	}
	clientInterface interface {
		GetUserById(ctx context.Context, id string) (*apiv1.UserProfile, error)
		GetAccessTokenByToken(ctx context.Context, token string) (*apiv1.AccessToken, error)
	}
)

var (
	ErrTokenInadequateSocpe = errors.New("access token has inadequate scope")
	ErrAccessTokenExpired   = errors.New("access token is expired")
)

func (s *Service) VerifyAccessToken(ctx context.Context, accesstoken string) (*apiv1.AccessToken, error) {
	token, err := s.client.GetAccessTokenByToken(ctx, accesstoken)
	if err != nil {
		return nil, fmt.Errorf("cannot get access token: %w", err)
	}
	if time.Now().After(token.Expires.AsTime()) {
		return nil, ErrAccessTokenExpired
	}
	return token, nil
}

// Can be used if the scope has a profile:view
func (s *Service) ViewUserProfile(ctx context.Context, userId string) (*apiv1.UserProfile, error) {
	user, err := s.client.GetUserById(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("cannot get user: %w", err)
	}
	return user, nil
}
