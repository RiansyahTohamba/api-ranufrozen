migrasi dari datastructure-class ke table-db

jika pustaka yg digunakan adalah gorm, maka terdapat tipe data yang tidak tepat untuk storage production, misalkan
string -> longtext

sebaiknya string menggunakan varchar bukannya longtext

tetapi ada cara untuk mengganti string tadi menjadi longtext.

