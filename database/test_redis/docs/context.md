context tergantung interface apa yang akan kita gunakan:
jika pakai gin.RestAPI maka gunakan gin.Context
jika pakai CLI maka gunakan context.Background
https://github.com/go-redis/redis/issues/1594
