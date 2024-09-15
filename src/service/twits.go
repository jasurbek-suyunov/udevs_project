package service

import (
	"context"
	"errors"
	"github.com/jasurbek-suyunov/udevs_project/models"
)

func (s *Service) CreateTwit(ctx context.Context, twit *models.Twit) (*models.Twit, error) {
	return s.storage.Twit().CreateTwit(ctx, twit)
}

func (s *Service) UpdateTwit(ctx context.Context, twit *models.Twit) (*models.Twit, error) {
	canModify, err := s.storage.Twit().CanModifyTwit(ctx, twit.UserID, twit.ID)
	if err != nil {
		return nil, err
	}
	if !canModify {
		return nil, errors.New("you can't modify this twit")
	}
	return s.storage.Twit().UpdateTwit(ctx, twit)
}

func (s *Service) DeleteTwit(ctx context.Context, userID, twitID string) error {
	canModify, err := s.storage.Twit().CanModifyTwit(ctx, userID, twitID)
	if err != nil {
		return err
	}

	if !canModify {
		return errors.New("you can't delete this twit")
	}

	return s.storage.Twit().DeleteTwit(ctx, twitID)
}

func (s *Service) GetTwitByID(ctx context.Context, id string) (*models.Twit, error) {
	return s.storage.Twit().GetTwitByID(ctx, id)
}

func (s *Service) GetTwits(ctx context.Context, userID string) ([]models.Twit, error) {
	return s.storage.Twit().GetTwits(ctx, userID)
}

func (s *Service) GetTwitsByUserID(ctx context.Context, userID string) ([]models.Twit, error) {
	return s.storage.Twit().GetTwitsByUserID(ctx, userID)
}

func(s *Service) LikeTwit(ctx context.Context, userID ,twitID string) error {
	return s.storage.Twit().LikeTwit(ctx,userID,twitID)
}

func(s *Service) UnLikeTwit(ctx context.Context, userID ,twitID string) error {
	return s.storage.Twit().UnlikeTwit(ctx,userID,twitID)
}

func(s *Service) RetwitTwit(ctx context.Context, userID ,twitID string) error {
	return s.storage.Twit().RetwitTwit(ctx,userID,twitID)
}

func(s *Service) UnRetwitTwit(ctx context.Context, userID ,twitID string) error {
	return s.storage.Twit().UnRetwitTwit(ctx,userID,twitID)
}