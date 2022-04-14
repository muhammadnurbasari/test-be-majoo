package reports

import "github.com/muhammadnurbasari/test-be-majoo/models/reportModel"

type ReportUsecase interface {
	GetPaginationOmzetMerchants(req *reportModel.ReqOmzetMerchant) (*reportModel.ResOmzetMerchant, error)
	GetPaginationOmzetOutlet(req *reportModel.ReqOmzetMerchant) (*reportModel.ResOmzetOutlet, error)
}
