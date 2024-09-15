package service

import (
	"context"
	"github.com/jasurbek-suyunov/udevs_project/models"
)

func (s *Service) FollowUser(ctx context.Context, follower int, followed int) error {
	return s.storage.User().FollowUser(ctx, follower, followed)
}

func (s *Service) UnFollowUser(ctx context.Context, follower int, followed int) error {
	return s.storage.User().UnFollowUser(ctx, follower, followed)
}

func (s *Service) IsFollowing(ctx context.Context, follower int, followed int) (bool, error) {
	return s.storage.User().IsFollowing(ctx, follower, followed)
}

func (s *Service) GetFollowers(ctx context.Context, userID int) ([]models.Follow, error) {
	return s.storage.User().GetFollowers(ctx, userID)
}

func (s *Service) GetFollowing(ctx context.Context, userID int) ([]models.Follow, error) {
	return s.storage.User().GetFollowing(ctx, userID)
}

func (s *Service) GetFollowingList(ctx context.Context, userID int) ([]models.Follow, error) {
	return s.storage.User().GetFollowingList(ctx, userID)
}

func (s *Service) GetFollowersList(ctx context.Context, userID int) ([]models.Follow, error) {
	return s.storage.User().GetFollowersList(ctx, userID)
}

func (s *Service) Search(ctx context.Context, query string) ([]models.SearchResult, error) {
	return s.storage.User().Search(ctx, query)
}

func (s *Service) UploadProfileImage(ctx context.Context, userID string, url string) error {
	return s.storage.User().UploadProfileImage(ctx, userID, url)
}