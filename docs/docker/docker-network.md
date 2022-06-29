karena 1 service punya masing-masing 1 container maka kita perlu mencari cara `bagaimana agar antar container bisa saling terhubung?`

misalkan kita akan buat 3 service
1. host app
2. redis 
3. mongo

maka kita perlu menghubungkan host app kita kepada service redis dan service mongo.

docker menyediakan `docker network` untuk menghubungkan antar container tersebut.

kita create network yang berisi container yang akan dihubungkan.

