package repository

import (
	"context"
	"encoding/json"
	"github.com/fernandomajeric/ml-challenge/app/model"
	"github.com/fernandomajeric/ml-challenge/config"
	"github.com/go-redis/redis/v8"
	"strconv"
)

type StatisticRepositoryInterface interface {
	Increment(statist model.StatisticItem) error
	GetScores() map[string]model.StatisticCore
}

type StatisticRepository struct{}

func (StatisticRepository) GetScores() map[string]model.StatisticCore {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Configuration.RedisCache.Addrs,
		Password: config.Configuration.RedisCache.Password, // no password set
		DB:       config.Configuration.RedisCache.DB,       // use default DB
	})

	var keys []string
	var item model.StatisticItem
	var list = map[string]model.StatisticCore{}
	var core model.StatisticCore
	keys, _ = rdb.ZScan(ctx, "Statistic", 0, "*", 0).Val()
	for i, k := range keys {
		if i%2 == 0 {
			json.Unmarshal([]byte(k), &item)
			core = model.StatisticCore{
				Country:  item.CountryName,
				Distance: item.Distance,
				HitCount: 0,
			}
			list[item.CountryName] = core
		} else {
			core.HitCount, _ = strconv.ParseInt(k, 10, 64)
			list[item.CountryName] = core
		}
	}

	return list
}

func (StatisticRepository) Increment(statist model.StatisticItem) error {
	var ctx = context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Configuration.RedisCache.Addrs,
		Password: config.Configuration.RedisCache.Password, // no password set
		DB:       config.Configuration.RedisCache.DB,       // use default DB
	})
	parsed, _ := json.Marshal(statist)
	_, err := rdb.ZIncrBy(ctx, "Statistic", 1.0, string(parsed)).Result()

	if err != nil {
		return err
	}
	return nil
}
