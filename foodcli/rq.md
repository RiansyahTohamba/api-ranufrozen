
https://dev.to/franciscomendes10866/caching-in-golang-using-redis-418n

read data food in redis not in DB.

update data in redis if DB is up to date.
how to ensure user always get fresh data?
1. baca buku redis
2. lihat SF
https://stackoverflow.com/questions/49299958/how-would-redis-get-to-know-if-it-has-to-return-cached-data-or-fresh-data-from-d?noredirect=1&lq=1

# RETRIEVE data
```go
// apakah key-nya terdapat pada redis?
val := rcl.get(foodkey).result()

if (val != nil){
    // ctx.json(val)
    fmt.Printlin(val)
}else{
    foods := repo.findAll()
    foodjson := json.newencoder(foods)
    rcl.set(foodkey,foodjson)
    // ctx.json(foodjson)
    fmt.Printlin(foodjson)
}


```
# CREATE,UPDATE,DELETE
jika terjadi perubahan create, update, delete
```go

```
