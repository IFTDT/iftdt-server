package service

import (
	"github.com/iftdt/server/dto"
	"github.com/iftdt/server/entity"
	"github.com/iftdt/server/repository"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	DEVICETOKEN = "1231321231"
)

var testVideo entity.Device = entity.Device{
	DeviceToken: DEVICETOKEN,
	UserId:      3,
}

var _ = Describe("Device Service", func() {

	var (
		deviceRepository repository.DeviceRepository
		deviceService    DeviceService
	)

	BeforeEach(func() {
		deviceRepository = repository.NewDeviceRepository(dto.DB)
		deviceService = NewDeviceService(deviceRepository)
	})

	Describe("Fetching all existing devices", func() {

		Context("If there is a device in the database", func() {

			BeforeEach(func() {
				deviceService.Create(testVideo)
			})

			It("should return at least one element", func() {
				videoList := deviceService.FindAll()

				Ω(videoList).ShouldNot(BeEmpty())
			})

			It("should map the fields correctly", func() {
				firstDevice := deviceService.FindAll()[0]

				Ω(firstDevice.DeviceToken).Should(Equal(DEVICETOKEN))
			})

			AfterEach(func() {
				video := deviceService.FindAll()[0]
				deviceService.Delete(video)
			})

		})

		Context("If there are no device in the database", func() {

			It("should return an empty list", func() {
				devices := deviceService.FindAll()

				Ω(devices).Should(BeEmpty())
			})

		})
	})

})
