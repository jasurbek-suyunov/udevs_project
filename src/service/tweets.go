package service

import (
	"context"
	"errors"
	"github.com/jasurbek-suyunov/udevs_project/models"
)

func (s *Service) CreateTweet(ctx context.Context, tweet *models.Tweet) (*models.Tweet, error) {
	return s.storage.Tweet().CreateTweet(ctx, tweet)
}

func (s *Service) UpdateTweet(ctx context.Context, tweet *models.Tweet) (*models.Tweet, error) {
	canModify, err := s.storage.Tweet().CanModifyTweet(ctx, tweet.UserID, tweet.ID)
	if err != nil {
		return nil, err
	}
	if !canModify {
		return nil, errors.New("you can't modify this tweet")
	}
	return s.storage.Tweet().UpdateTweet(ctx, tweet)
}

func (s *Service) DeleteTweet(ctx context.Context, userID, tweetID string) error {
	canModify, err := s.storage.Tweet().CanModifyTweet(ctx, userID, tweetID)
	if err != nil {
		return err
	}

	if !canModify {
		return errors.New("you can't delete this tweet")
	}

	return s.storage.Tweet().DeleteTweet(ctx, tweetID)
}

func (s *Service) GetTweetByID(ctx context.Context, id string) (*models.Tweet, error) {
	return s.storage.Tweet().GetTweetByID(ctx, id)
}

func (s *Service) GetTweets(ctx context.Context, userID string) ([]models.Tweet, error) {
	return s.storage.Tweet().GetTweets(ctx, userID)
}

func (s *Service) GetTweetsByUserID(ctx context.Context, userID string) ([]models.Tweet, error) {
	return s.storage.Tweet().GetTweetsByUserID(ctx, userID)
}

func(s *Service) LikeTweet(ctx context.Context, userID ,tweetID string) error {
	return s.storage.Tweet().LikeTweet(ctx,userID,tweetID)
}

func(s *Service) UnLikeTweet(ctx context.Context, userID ,tweetID string) error {
	return s.storage.Tweet().UnlikeTweet(ctx,userID,tweetID)
}

func(s *Service) RetweetTweet(ctx context.Context, userID ,tweetID string) error {
	return s.storage.Tweet().RetweetTweet(ctx,userID,tweetID)
}

func(s *Service) UnRetweetTweet(ctx context.Context, userID ,tweetID string) error {
	return s.storage.Tweet().UnRetweetTweet(ctx,userID,tweetID)
}