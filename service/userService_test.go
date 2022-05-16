package service

import (
	"github.com/iftdt/server/dto"
	"github.com/iftdt/server/entity"
	"github.com/iftdt/server/repository"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	ADMIN    = "admin"
	PASSWORD = "123456"
)

var testUser entity.User = entity.User{
	Name:     ADMIN,
	Password: PASSWORD,
}

var _ = Describe("User Service", func() {

	var (
		userRepository repository.UserRepository
		userService    UserService
	)

	BeforeEach(func() {
		userRepository = repository.NewUserRepository(dto.DB)
		userService = NewUserService(userRepository)
	})

	Describe("Fetching all existing users", func() {

		Context("If there is a user in the database", func() {

			BeforeEach(func() {
				userService.Create(testUser)
			})

			It("should return at least one element", func() {
				videoList := userService.FindAll()

				Ω(videoList).ShouldNot(BeEmpty())
			})

			It("should map the fields correctly", func() {
				firstUser := userService.FindAll()[0]

				Ω(firstUser.Name).Should(Equal(ADMIN))
			})

			AfterEach(func() {
				// video := userService.FindAll()[0]
				// deviceService.Delete(video)
			})

		})

		// Context("If there are no user in the database", func() {

		// It("should return an empty list", func() {
		// 	users := userService.FindAll()

		// 	Ω(users).Should(BeEmpty())
		// })

		// })
	})
})
