package adapters_test

import (
	"fmt"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vmware-tanzu-labs/git-story/adapters"
)

var _ = Describe("Browser", func() {
	path := os.Getenv("PATH")

	BeforeEach(func() {
		os.Setenv("PATH", fmt.Sprintf("./fixtures:%s", path))
	})
	AfterEach(func() {
		os.Setenv("PATH", path)
	})

	It("should open the given URL", func() {
		cmd, error := adapters.NewBrowser().OpenURL("https://www.url-i-want.com")
		Expect(error).To(BeNil())
		error = cmd.Wait()
		Expect(error).To(BeNil())
		cmd, _ = adapters.NewBrowser().OpenURL("https://www.url-i-dont-want.com")
		error = cmd.Wait()
		Expect(error).NotTo(BeNil())
	})
})
