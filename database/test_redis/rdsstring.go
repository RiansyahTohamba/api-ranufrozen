package main

import (
	"log"
)

type StringRepo struct {
	rCl *RedisClient
}

func NewStringRepo(rCl *RedisClient) *StringRepo {
	return &StringRepo{rCl}
}

func (sd *StringRepo) Create(key, val string) {
	err := sd.rCl.Set(ctx, key, val, 0).Err()
	if err != nil {
		log.Println(err)
	}
}

func (sd *StringRepo) Show(key string) string {
	val, err := sd.rCl.Get(ctx, key).Result()
	if err != nil {
		log.Println(err)
	}
	return val
}

func (sd *StringRepo) Update() {

}

func (sd *StringRepo) Delete() {

}
