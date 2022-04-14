package reportUsecase

import (
	"github.com/muhammadnurbasari/test-be-majoo/models/reportModel"
	"github.com/muhammadnurbasari/test-be-majoo/modules/reports"
)

type reportsUsecase struct {
	reportsRepository reports.ReportRepository
}

//NewReportUsecase - will create new an userUsecase object representation of reports.ReportUsecase interface
func NewReportUsecase(reportsRepository reports.ReportRepository) reports.ReportUsecase {
	return &reportsUsecase{
		reportsRepository: reportsRepository,
	}
}

// GetPaginationOmzetMerchants - usecase get paginations omzet merchants
func (myUC *reportsUsecase) GetPaginationOmzetMerchants(req *reportModel.ReqOmzetMerchant) (*reportModel.ResOmzetMerchant, error) {
	resp, err := myUC.reportsRepository.PaginationOmzetMerchants(req.UserID, req.Start, req.RowPerPage)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetPaginationOmzetOutlet - usecase get paginations omzet outlet
func (myUC *reportsUsecase) GetPaginationOmzetOutlet(req *reportModel.ReqOmzetMerchant) (*reportModel.ResOmzetOutlet, error) {
	resp, err := myUC.reportsRepository.PaginationOmzetOutlet(req.UserID, req.Start, req.RowPerPage)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
