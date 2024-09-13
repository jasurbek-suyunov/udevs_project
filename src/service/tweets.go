package service

import (
	"context"
	"jas/models"
)

func (s *Service) CreateTweet(ctx context.Context, tweet *models.Tweet) (*models.Tweet, error) {
	return s.storage.Tweet().CreateTweet(ctx, tweet)
}
