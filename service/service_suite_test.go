package service

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDeviceService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Device Service Test Suite")
}
