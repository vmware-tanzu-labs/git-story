package adapters_test

import (
	"fmt"
	"os"

	"github.com/git-story-branch/git-story-branch/adapters"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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
