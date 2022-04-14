package reportRoutes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/muhammadnurbasari/test-be-majoo/modules/reports/reportHandler"
	"github.com/muhammadnurbasari/test-be-majoo/modules/reports/reportRepository"
	"github.com/muhammadnurbasari/test-be-majoo/modules/reports/reportUsecase"
)

func ReportRoutes(r *gin.Engine, mySQL *sql.DB) {
	reportRepo := reportRepository.NewReportRepository(mySQL)
	reportUC := reportUsecase.NewReportUsecase(reportRepo)
	reportHandler.NewReportHandler(r, reportUC)
}
