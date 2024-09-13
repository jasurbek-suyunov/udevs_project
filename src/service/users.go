package service

import (
	"context"
	"jas/models"
)

func (s *Service) FollowUser(ctx context.Context, follower int, followed int) error {
	// follow user
	err := s.storage.User().FollowUser(ctx, follower, followed)
	// check error
	if err != nil {
		return err
	}
	// return result if no error
	return nil
}

func (s *Service) UnFollowUser(ctx context.Context, follower int, followed int) error {
	// unfollow user
	err := s.storage.User().UnFollowUser(ctx, follower, followed)
	// check error
	if err != nil {
		return err
	}
	// return result if no error
	return nil
}

func (s *Service) IsFollowing(ctx context.Context, follower int, followed int) (bool, error) {
	// check if user is following
	result, err := s.storage.User().IsFollowing(ctx, follower, followed)
	// check error
	if err != nil {
		return false, err
	}
	// return result if no error
	return result, nil
}

func (s *Service) GetFollowers(ctx context.Context, userID int) ([]models.User, error) {
	// get followers
	result, err := s.storage.User().GetFollowers(ctx, userID)
	// check error
	if err != nil {
		return nil, err
	}
	// return result if no error
	return result, nil
}

func (s *Service) GetFollowing(ctx context.Context, userID int) ([]models.User, error) {
	// get following
	result, err := s.storage.User().GetFollowing(ctx, userID)
	// check error
	if err != nil {
		return nil, err
	}
	// return result if no error
	return result, nil
}