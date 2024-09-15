package service

import (
	"context"
	"github.com/jasurbek-suyunov/udevs_project/models"
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

func (s *Service) GetFollowers(ctx context.Context, userID int) ([]models.Follow, error) {
	// get followers
	result, err := s.storage.User().GetFollowers(ctx, userID)
	// check error
	if err != nil {
		return nil, err
	}
	// return result if no error
	return result, nil
}

func (s *Service) GetFollowing(ctx context.Context, userID int) ([]models.Follow, error) {
	// get following
	result, err := s.storage.User().GetFollowing(ctx, userID)
	// check error
	if err != nil {
		return nil, err
	}
	// return result if no error
	return result, nil
}

func (s *Service) GetFollowingList(ctx context.Context, userID int) ([]models.Follow, error) {
	// get following list
	result, err := s.storage.User().GetFollowingList(ctx, userID)
	// check error
	if err != nil {
		return nil, err
	}
	// return result if no error
	return result, nil
}

func (s *Service) GetFollowersList(ctx context.Context, userID int) ([]models.Follow, error) {
	// get followers list
	result, err := s.storage.User().GetFollowersList(ctx, userID)
	// check error
	if err != nil {
		return nil, err
	}
	// return result if no error
	return result, nil
}

func (s *Service) Search(ctx context.Context, query string) ([]models.SearchResult, error) {
	// search
	result, err := s.storage.User().Search(ctx, query)
	// check error
	if err != nil {
		return nil, err
	}
	// return result if no error
	return result, nil
}

func (s *Service) UploadProfileImage(ctx context.Context, userID string, url string) error {
	// upload profile image
	err := s.storage.User().UploadProfileImage(ctx, userID, url)
	// check error
	if err != nil {
		return err
	}
	// return result if no error
	return nil
}