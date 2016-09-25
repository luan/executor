package guidgen_test

import (
	. "code.cloudfoundry.org/executor/guidgen"
	"code.cloudfoundry.org/lager/lagertest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GUID Generator", func() {
	Context("when a prefix is provided", func() {
		It("prepends the prefix to the guid", func() {
			logger := lagertest.NewTestLogger("test")
			Expect(DefaultGenerator.Guid(logger, "prefix")).To(HavePrefix("prefix"))
			Expect(DefaultGenerator.Guid(logger, "someting")).To(HavePrefix("someting"))
		})
	})
})
