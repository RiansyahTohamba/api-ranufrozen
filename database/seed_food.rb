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


foods = get_seed_food
insert_to_mongo(foods)
