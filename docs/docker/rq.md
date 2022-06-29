1. WORKDIR /app

WORKDIR = `mkdir` then `cd`

command diatas bermakna
```sh
mkdir /app
cd /app
```

2. beberapa command umum
WORKDIR
ADD
CMD
RUN
COPY
ENTRYPOINT

3. container vs images ?

images: 
    - template yg digunakan untuk membuat container.
    - sudah di build
    - An image is an inert, immutable, file that's essentially a snapshot of a container. 
    - snapshot dari container
container:
    - guess OS, sama seperti VM    
    - host OS adalah OS yang sedang kita gunakan, Guess OS adalah tamu dari host
    - didalam container banyak images
    - Images are created with the build command, and they'll produce a container when started with run.

kalau sudah dibuild, namanya image. 
begitu sudah jalan, jadi container.


1 image bisa dibuat beberapa container.
contoh image nodejs bisa dibuat 3 container.
1. container untuk service authenticate
2. container untuk service order
3. container untuk service product

container bisa dalam state inactive maupun active.

# What is a Docker container?
Once you have a container image, you can start a container based on it. 
Docker containers are the way to execute that package of instructions in a runtime environment. Start a container using the docker run command.
https://www.techtarget.com/searchitoperations/answer/What-is-a-Docker-container-vs-an-image


4. docker compose ?
Docker Compose. Best practices dalam Docker adalah menjalankan satu proses dalam satu container. 

Misalnya, bila kita ingin menjalankan Wordpress, maka kita butuh
    a. satu container untuk webserver (webserver.Dockerfile), 
    b. satu container untuk database server (db.Dockerfile), 
    c. satu container untuk menyimpan data dari database server (data.Dockerfile), dan 
    d. satu container lagi untuk menyimpan file dari webserver (static.Dockerfile). 
    
Untuk memudahkan pengelolaannya, Docker menyediakan aplikasi yang bernama Docker Compose.


5. bagaimana caranya expose port pada docker?
pembahasannya ada di docker-network