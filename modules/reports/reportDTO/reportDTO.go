package reportDTO

import (
	"github.com/muhammadnurbasari/test-be-majoo/models/reportModel"
)

type (
	// ResInfoOmzetMerchant
	ResInfoOmzetMerchant struct {
		Status    int                          `json:"status" example:"200"`
		Messagges string                       `json:"messages" example:"Success"`
		Result    reportModel.ResOmzetMerchant `json:"result"`
	}

	// ResInfoOmzetOutlet
	ResInfoOmzetOutlet struct {
		Status    int                        `json:"status" example:"200"`
		Messagges string                     `json:"messages" example:"Success"`
		Result    reportModel.ResOmzetOutlet `json:"result"`
	}
)
