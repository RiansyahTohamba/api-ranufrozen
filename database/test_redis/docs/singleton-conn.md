# singleton in connection
Connection untuk redis menggunakan Singleton pattern.

# singleton pattern pada go ditandai dengan:
1. buat type terpisah untuk redis client
type RedisClient struct{ *redis.Client }

2. gunakan sync.Once, method Once.Do
var once sync.Once
once.Do(RedisConn)


We now create a struct that contains a pointer to the redis client. This pointer will have the functions that help us with this service, we also create a constant with the key name for our set in redis.

```go
type RedisClient struct { *redis.Client }
const key = "drivers"
```

For the function to get the Redis client, we are going to use the singleton pattern with the help of the `sync package` and its `Once.Do` functionality.

In software engineering, the singleton pattern is a software design pattern that restricts the instantiation of a class to one object. 

This is useful when exactly one object is needed to coordinate actions across the system. If you want to read more about Singleton Pattern.

But how does once.Do work?

The struct `sync.Once` has an atomic counter and it uses atomic.StoreUint32 to set a value to 1, when the function has been called, and then atomic.LoadUint32 to see if it needs to be called again. 

For this basic implementation GetRedisClient will be called from two endpoints but we only want to `get one instance`.










