package food_test

import (
	food "api-ranufrozen/food"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("Service", func() {
	// setup ?
	// foodService butuh foodRepository
	// foodRepository := &food.NewRepository(db)
	foodRepository := food.RepositoryMock{Mock: mock.Mock{}}
	foodService := food.NewService(foodRepository)
	foodReq := food.FoodRequest{
		Name:  "Riansyah",
		Price: "-100.0",
	}
	// nah bagaimana cara buat unit testing untuk service class?
	// how to test it?

	// pakai double seharusnya ya?
	Describe("Set Food Price", func() {
		// input nya?
		When("price is minus", func() {
			It("returns a invalid price", func() {
				Expect(foodService.Create(foodReq)).To(BeNumerically(">", 0))
			})
		})
		When("price is zero", func() {
			It("returns a invalid price", func() {
				Expect(foodService.Sum(2, 2)).To(BeNumerically(">", 0))
			})
		})

		It("save price to persistent", func() {
			expected := foodService.Sum(2, 2)
			toBe := BeNumerically(">", 0)

			Expect(expected).To(toBe)
		})

	})
})
