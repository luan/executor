package guidgen_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGuidgen(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Guidgen Suite")
}
