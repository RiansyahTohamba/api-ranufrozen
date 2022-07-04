package main

import "log"

type HashRepo struct {
	rdCl *RedisClient
}

func NewHashRepo(rdCl *RedisClient) *HashRepo {
	return &HashRepo{rdCl}
}

// HSet accepts values in following formats:
//   - HSet("myhash", "key1", "value1", "key2", "value2")
//   - HSet("myhash", []string{"key1", "value1", "key2", "value2"})
//   - HSet("myhash", map[string]interface{}{"key1": "value1", "key2": "value2"})

func (hr *HashRepo) Create() {
	// hr.rdCl.HSet("")
}

func (hr *HashRepo) FindAll(hashkey string) map[string]string {
	res, err := hr.rdCl.HGetAll(ctx, hashkey).Result()
	if err != nil {
		log.Println(err)
	}
	return res
}

// hset cart:c1 id1 "cola:30"
// hset cart:c1 id2 "juswortel:5"

// hset cart:c2 id1 "jusjeruk:15"

// hset cart:c1 id3 "kueberas:25"
