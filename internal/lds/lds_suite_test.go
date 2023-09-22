package lds

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLds(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lds Suite")
}
