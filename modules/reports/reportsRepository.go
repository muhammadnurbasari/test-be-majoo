package reports

import "github.com/muhammadnurbasari/test-be-majoo/models/reportModel"

type ReportRepository interface {
	PaginationOmzetMerchants(userID int64, start, rowPerPage int) (*reportModel.ResOmzetMerchant, error)
	PaginationOmzetOutlet(userID int64, start, rowPerPage int) (*reportModel.ResOmzetOutlet, error)
}
