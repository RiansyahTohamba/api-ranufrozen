# when Create() has writed in Service
cannot use foodRepository (variable of type food.RepositoryMock) as food.Repository value in argument to food.NewService: missing method Create (Create has pointer receiver)

cannot use foodRepository 

mock.Mock bisa digunakan untuk mengganti foodRepository. 
so problemnya bukan tentang tipe mock, tetapi masalahnya pada `missing method Create() yang punya receiver` alias create punya FoodRepository.

Create has pointer receiver.
nah ini apa?

coba kita bandingkan Create() pada 
1. food.RepositoryMock
2. food.Repository


## pada RepositoryMock
`
func (repo *RepositoryMock) Create(food Food) (Food, error) {
	nugget := Food{
		Name: "nugget",
	}
	return nugget, nil
}
`

## pada Repository
`
func (r *repository) Create(food Food) (Food, error) {
	err := r.db.Create(&food).Error
	return food, err
}
`


receiver pada `struct` repository ini yang bikin error. coba cermati lagi tutorial milik PZN.

= apakah PZN membuat CategoryRepository?
> ya dibuat.

= apakah CategoryRepository punya interface?
> ya punya

= apakah CategoryRepository punya struct?
> belum dibuat jadi method-struct, masih function dari interface.



= apakah pada CategoryRepository terdapat method FindById() ?
> 

= apakah pada method FindById() terdapat receiver? 


= apakah receiver nya dari struct categoryRepository?


# skip pertanyaan diatas
coba kita lihat contoh kode TestService yang diberikan oleh PZN.







# what problems?
cannot use foodRepository (variable of type food.RepositoryMock) as
food.Repository value in argument to food.NewService

food.RepositoryMock sebagai nilai food.Repository pada argument foodNewService
missing method Create. missing di foodRepositoryMock?

kalau gitu, coba kita hilangkan Create pada interface Food Repository
yang terjadi, menyasar findAll
lalu findAll saya hapus, dan keluarlah hasil berikut:
wrong type for method FindById (have func(id int) *api-ranufrozen/food.Food, want func(id int) (api-ranufrozen/food.Food, error))

FindById sudah dideteksi tapi salah type nya
saat ini
func(id int) Create *api-ranufrozen/food.Food

seharusnya
func(id int) Create (api-ranufrozen/food.Food, error)

bukan pointer dan tak ada error
int

jadi kenapa repositoryMock dianggap tidak memiliki method tersebut.
padahal sdh sy buatkan?
masalahnya sudah fix