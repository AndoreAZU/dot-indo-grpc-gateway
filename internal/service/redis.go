package service

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v9"
)

type Redis interface {
	SetKey(ctx context.Context, k string, v string, duration time.Duration) error
	GetKeyString(ctx context.Context, k string) (v string, err error)
	GetKeyHash(ctx context.Context, k, f string) (interface{}, error)
	GetKeyToStruct(ctx context.Context, k string, dst any) (err error)
	RemoveKey(ctx context.Context, k string) (err error)
}

type instance struct {
	Client *redis.Client
}

func (x *instance) SetKey(ctx context.Context, k string, v string, duration time.Duration) error {
	if x.Client == nil {
		return fmt.Errorf("redis_not_connect")
	}
	return x.Client.Set(ctx, k, v, duration).Err()
}

func (x *instance) GetKeyString(ctx context.Context, k string) (v string, err error) {
	if x.Client == nil {
		return "", fmt.Errorf("redis_not_connect")
	}
	t, err := x.Client.Get(ctx, k).Result()
	if err != nil {
		return
	}

	return t, nil
}

func (x *instance) GetKeyHash(ctx context.Context, k, f string) (interface{}, error) {
	if x.Client == nil {
		return nil, fmt.Errorf("redis_not_connect")
	}
	r, err := x.Client.HMGet(ctx, k, f).Result()
	return r[0], err
}

func (x *instance) GetKeyToStruct(ctx context.Context, k string, dst any) (err error) {
	t, err := x.Client.Get(ctx, k).Result()
	if x.Client == nil {
		return fmt.Errorf("redis_not_connect")
	}
	if err != nil {
		return
	}
	return json.Unmarshal([]byte(t), &dst)
}

func (x *instance) RemoveKey(ctx context.Context, k string) (err error) {
	return x.Client.Del(ctx, k).Err()
}

func NewRedis() Redis {
	redis_cache := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	_, err := redis_cache.Ping(context.Background()).Result()
	if err != nil {
		return &instance{}
	} else {
		return &instance{
			Client: redis_cache,
		}
	}
}
