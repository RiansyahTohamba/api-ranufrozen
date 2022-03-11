# relasi antar kelas
1. relasi struct-> repository
relasi ini untuk penyimpanan data ke persistent

2. repository -> service -> handler
relasi ini untuk menghubungkan router dengan handler terkait

3. handler -> main
relasi ini untuk menghubungkan set up antara handler dan kelas service dan repository terkait


# kode ini  ditest pada active record

test "product price must be positive" do
 	product = Product.new(title: "My Book Title", description: "yyy",image_url: "zzz.jpg")

 	product.price = -1
 	assert product.invalid?
 	assert_equal ["must be greater than or equal to 0.01"],
 	product.errors[:price]

 	product.price = 0
 	assert product.invalid?
 	assert_equal ["must be greater than or equal to 0.01"],
 	product.errors[:price]

 	product.price = 1
 	assert product.valid?
end

the question is:
`apa padanan konsep active_record untuk repo-service pattern dalam app go ini?`
