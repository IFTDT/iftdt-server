package service

import (
	"github.com/iftdt/server/entity"
	"github.com/iftdt/server/repository"
)

type DeviceService interface {
	Create(entity.Device) entity.Device
	Update(entity.Device) entity.Device
	Delete(entity.Device) entity.Device
	FindOne(id string) entity.Device
	FindAll() []entity.Device
}

type deviceService struct {
	deviceRepository repository.DeviceRepository
}

func NewDeviceService(repo repository.DeviceRepository) DeviceService {
	return &deviceService{
		deviceRepository: repo,
	}
}

func (service *deviceService) Create(device entity.Device) entity.Device {
	return service.deviceRepository.Create(device)
}

func (service *deviceService) Update(device entity.Device) entity.Device {
	service.deviceRepository.Update(device)
	return device
}

func (service *deviceService) Delete(device entity.Device) entity.Device {
	service.deviceRepository.Delete(device)
	return device
}

func (service *deviceService) FindOne(id string) entity.Device {
	return service.deviceRepository.FindOne(id)
}

func (service *deviceService) FindAll() []entity.Device {
	return service.deviceRepository.FindAll()
}
