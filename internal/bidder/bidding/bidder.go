package bidding

import (
	"math/rand"
	"strconv"

	"github.com/adgear/go-commons/pkg/log"
	"github.com/adgear/go-commons/pkg/metric"
	"github.com/prebid/openrtb/v20/adcom1"
	openrtb2 "github.com/prebid/openrtb/v20/openrtb2"
)

// Enforce `Bidder` interface implementation.
var _ Bidder = (*service)(nil)

// service is the struct implementing `Bidder` interface
type service struct {
	logger  log.Service
	metrics metric.Service
}

// Bid function processing bid request.
func (s *service) Bid(req *openrtb2.BidRequest) (resp *openrtb2.BidResponse, err error) {
	impId := "test_impid"
	bidReqid := "test_reqid"

	if len(req.Imp) > 0 {
		impId = req.Imp[0].ID
	}

	if len(req.ID) > 0 {
		bidReqid = req.ID
	}

	resp = &openrtb2.BidResponse{
		ID: req.ID,
		SeatBid: []openrtb2.SeatBid{
			{
				Bid: []openrtb2.Bid{
					{
						ID:             bidReqid,
						ImpID:          impId,
						Price:          10,
						NURL:           "",
						BURL:           "",
						LURL:           "",
						AdM:            "",
						AdID:           "test_adid",
						ADomain:        []string{},
						Bundle:         "",
						IURL:           "",
						CID:            "test_cid",
						CrID:           "test_crid",
						Tactic:         "",
						CatTax:         0,
						Cat:            []string{},
						Attr:           []adcom1.CreativeAttribute{},
						APIs:           []adcom1.APIFramework{},
						API:            0,
						Protocol:       0,
						QAGMediaRating: 0,
						Language:       "",
						LangB:          "",
						DealID:         "",
						W:              0,
						H:              0,
						WRatio:         0,
						HRatio:         0,
						Exp:            0,
						Dur:            0,
						MType:          0,
						SlotInPod:      0,
						Ext:            []byte{},
					},
				},
				Seat:  "",
				Group: 0,
				Ext:   []byte{},
			},
		},
		BidID:      "test_" + strconv.Itoa(rand.Int()),
		Cur:        "USD",
		CustomData: "",
		Ext:        []byte{},
	}
	return
}
