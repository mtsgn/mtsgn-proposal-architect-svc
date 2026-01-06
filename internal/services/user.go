package services

import (
	"boilerplate-api/internal/common/constant"
	"boilerplate-api/internal/models"
	"boilerplate-api/internal/repositories"
	"boilerplate-api/pkg/redis"
	"context"
	"encoding/json"

	logs "github.com/rs/zerolog/log"
)

type User interface {
	FindById(ctx context.Context, id string) (*models.User, error)
}

type user struct {
	redis    *redis.RedisClient
	userRepo *repositories.UserRepository
}

func NewUserService(redis *redis.RedisClient, userRepo *repositories.UserRepository) User {
	return &user{
		redis:    redis,
		userRepo: userRepo,
	}
}

func (s *user) Create(user *models.User) (*models.User, error) {
	created, err := s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}
	if err := s.redis.Set(context.Background(), constant.UserKey(created.ID.String()), created, constant.DefaultExpiration); err != nil {
		logs.Debug().Err(err).Str("key", constant.UserKey(created.ID.String())).Msg("failed to set user to cache")
	}
	return created, nil
}

func (s *user) Update(user *models.User) (*models.User, error) {
	updated, err := s.userRepo.Update(user)
	if err != nil {
		return nil, err
	}
	if err := s.redis.Set(context.Background(), constant.UserKey(user.ID.String()), updated, constant.DefaultExpiration); err != nil {
		logs.Debug().Err(err).Str("key", constant.UserKey(user.ID.String())).Msg("failed to set user to cache")
		// to prevent old data being cached
		_ = s.redis.Del(context.Background(), constant.UserKey(user.ID.String()))
	}

	return updated, nil
}

func (s *user) FindByPhoneNumber(dialCode, phoneNumber string) (*models.User, error) {
	return s.userRepo.FindByPhoneNumber(dialCode, phoneNumber)
}

func (s *user) FindById(ctx context.Context, id string) (*models.User, error) {
	cacheKey := constant.UserKey(id)

	cacheUser, err := s.redis.GetByBytes(ctx, cacheKey)
	if err != nil {
		logs.Debug().Err(err).Str("key", cacheKey).Msg("failed to get user from cache")
	} else if cacheUser != nil {
		var user models.User
		if err := json.Unmarshal(cacheUser, &user); err != nil {
			logs.Debug().Err(err).Str("key", cacheKey).Msg("failed to unmarshal user from cache")
			_ = s.redis.Del(ctx, cacheKey) // delete invalid cache
		} else {
			return &user, nil
		}
	}

	user, err := s.userRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	if err := s.redis.Set(ctx, cacheKey, user, constant.DefaultExpiration); err != nil {
		logs.Debug().Err(err).Str("key", cacheKey).Msg("failed to set user to cache")
	}

	return user, nil
}
