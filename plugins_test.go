package connect_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	. "github.com/go-kafka/connect"
)

var _ = Describe("Plugins Tests", func() {
	BeforeEach(func() {
		server = ghttp.NewServer()
		client = NewClient(server.URL())
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("ListPlugins", func() {
		var statusCode int
		resultPlugins := []*Plugin{
			&Plugin{
				Class:   "test-class",
				Type:    "source",
				Version: "5.3.0",
			},
		}

		BeforeEach(func() {
			statusCode = http.StatusOK

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/connector-plugins"),
					ghttp.VerifyHeader(jsonAcceptHeader),
					ghttp.RespondWithJSONEncodedPtr(&statusCode, &resultPlugins),
				),
			)
		})

		It("returns list of connector plugins", func() {
			plugins, _, err := client.ListPlugins()
			Expect(err).NotTo(HaveOccurred())
			Expect(plugins).To(Equal(resultPlugins))
		})
	})
})
