package serializer

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSerializer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Serializer Test Suite")
}

type BidReq struct {
	Id   string `json:"id"`
	TMax int64  `json:"tmax"`
}

var _ = Describe("Protobuf and json serialization", func() {
	var (
		s    service
		id   = "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
		tmax = int64(20)
	)
	BeforeEach(func() {
		s = service{}
	})

	jsonStr := func() string {
		return fmt.Sprintf("{\"id\":\"%s\",\"tmax\":%d}", id, tmax)
	}

	Context("Json Marshalling/Unmarshalling", func() {
		It("Converts Struct to Json string", func() {
			bid := BidReq{
				Id:   id,
				TMax: tmax,
			}
			str, err := s.MarshalJson(bid)
			Expect(err).Should(BeNil())
			Expect(string(str)).Should(Equal(jsonStr()))
		})

		It("Converts Json string to Struct", func() {
			bid := &BidReq{}
			err := s.UnmarshalJson([]byte(jsonStr()), bid)
			Expect(err).Should(BeNil())
			Expect(bid.Id).Should(Equal(id))
			Expect(bid.TMax).Should(Equal(tmax))
		})
	})
})
