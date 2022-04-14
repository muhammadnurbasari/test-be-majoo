package reportModel

import "time"

type (
	// DataOmzetMerchants- for mapping response data slice merchants omzet
	DataOmzetMerchants struct {
		UserID       int64      `json:"userid"`
		Name         string     `json:"name"`
		MerchantName string     `json:"merchant_name"`
		Tanggal      *time.Time `json:"tanggal"`
		Omzet        float32    `json:"omzet"`
	}

	// ResOmzetMerchant - object for response omzet merchant
	ResOmzetMerchant struct {
		Start      int                  `json:"start"`
		RowPerPage int                  `json:"row_per_page"`
		Data       []DataOmzetMerchants `json:"data"`
	}

	// ReqOmzetMerchant - object for request get pagination omzet merchants
	ReqOmzetMerchant struct {
		UserID     int64 `json:"userid"`
		Start      int   `json:"start"`
		RowPerPage int   `json:"row_per_page"`
	}
)

type (
	// DataOmzetOutlets- for mapping response data slice outlets omzet
	DataOmzetOutlets struct {
		UserID       int64      `json:"userid"`
		Name         string     `json:"name"`
		MerchantName string     `json:"merchant_name"`
		OutletName   string     `json:"outlet_name"`
		Tanggal      *time.Time `json:"tanggal"`
		Omzet        float32    `json:"omzet"`
	}

	// ResOmzetOutlet - object for response outlet merchant
	ResOmzetOutlet struct {
		Start      int                `json:"start"`
		RowPerPage int                `json:"row_per_page"`
		Data       []DataOmzetOutlets `json:"data"`
	}
)
