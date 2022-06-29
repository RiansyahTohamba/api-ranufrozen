# seed_food.rb
require "faker"
require "mongo"

def get_seed_food
	foods = []
	total_record = 30
	total_record.times do |n|
		food = {
			name: Faker::Food.dish,
			photoPath: "breakfast_item.jpg",
			rating: rand(1..5),
			price: rand(30000..55000),
			stock: rand(0..20),
			isSuperSeller: rand(0..1),
			category: rand(1..5),
			qtSold: rand(10..100),
			desc: "Rasanya bikin nyam nyam"
		}
		foods.push(food)	
	end
	# p foods.size
	# p foods
	foods
end

def insert_to_mongo(foods)
	client = Mongo::Client.new([ '127.0.0.1:27017' ], :database => 'ranufrozen')
	collection = client[:products]
	collection.insert_many(foods)
	p "succesfully inserted"	
	p collection.count
end

def test_query()

end

def insert_to_mysql()
	require "mysql2"    # if needed

	db_host  = "localhost"
	db_user  = "root"
	db_pass  = "Allah1takbir!"
	db_name = "api-ranufrozen"

	client = Mysql2::Client.new(:host => db_host, :username => db_user, :password => db_pass, :database => db_name)
	# cdr_result = client.query("SELECT * from your_db_table_name'")
	query_insert = "INSERT INTO foods (name, photo_path, rating, price, stock, is_super_seller, category, quantity_sold, description, discount) "
	query_insert += "VALUES "	

	total_record = 30
	total_record.times do |n|
		name = Faker::Food.dish
		photo_path= "breakfast_item.jpg"
		rating= rand(1..5)
		price= rand(30000..55000)
		stock= rand(0..20)
		is_super_seller= rand(0..1)
		category= rand(1..5)
		quantity_sold= rand(10..100)
		description= "Rasanya bikin nyam nyam"
		discount=0.1

		query = "('#{name}','#{photo_path}',#{rating},#{price},#{stock},#{is_super_seller},#{category},#{quantity_sold},'#{description}',#{discount})"

		if n != (total_record-1)			
			query_insert += "#{query},"
				
		else
			query_insert += "#{query};"
		end
	end
	# p query_insert
	
	client.query(query_insert)
	# p collection.count
	# p "succesfully inserted"	
end

# foods = get_seed_food
# insert_to_mongo(foods)
insert_to_mysql
