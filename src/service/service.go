package service

import "jas/src/storage"

type Service struct {
	storage storage.StorageI
	cache  storage.CacheStorageI
}

func NewService(repo storage.StorageI,cache storage.CacheStorageI) *Service {
	return &Service{
		storage: repo,
		cache: cache,
	}
}