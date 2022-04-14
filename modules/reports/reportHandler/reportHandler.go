package reportHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammadnurbasari/test-be-majoo/middlewares/auth"
	"github.com/muhammadnurbasari/test-be-majoo/models/reportModel"
	"github.com/muhammadnurbasari/test-be-majoo/modules/reports"
)

type reportHandler struct {
	reportUsecase reports.ReportUsecase
}

// NewReportHandler
func NewReportHandler(r *gin.Engine, reportUsecase reports.ReportUsecase) {
	handler := reportHandler{
		reportUsecase: reportUsecase,
	}

	authorized := r.Group("/report")

	{
		authorized.POST("/omzet-merchant", auth.JWTAuth, handler.ReportOmzetMerchant)
		authorized.POST("/omzet-outlet", auth.JWTAuth, handler.ReportOmzetOutlet)
	}

}

// ReportOmzetMerchant godoc
// @Summary get data and pagination omzet merchants
// @Description get data and pagination omzet merchants
// @Tags Reports
// @Accept  json
// @Produce  json
// @Param pageInfo body reportModel.ReqOmzetMerchant true "page info"
// @Success 200 {object} reportDTO.ResInfoOmzetMerchant
// @Router /report/omzet-merchant [post]
func (handler *reportHandler) ReportOmzetMerchant(c *gin.Context) {
	var req reportModel.ReqOmzetMerchant
	errBind := c.BindJSON(&req)

	if errBind != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":   http.StatusBadRequest,
				"messages": "Please Contact Admin !!!, " + errBind.Error(),
			},
		)
		return
	}

	res, err := handler.reportUsecase.GetPaginationOmzetMerchants(&req)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":   http.StatusNotFound,
				"messages": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":   http.StatusOK,
			"messages": "Success",
			"result":   &res,
		},
	)
}

// ReportOmzetOutlet godoc
// @Summary get data and pagination omzet merchants
// @Description get data and pagination omzet merchants
// @Tags Reports
// @Accept  json
// @Produce  json
// @Param pageInfo body reportModel.ReqOmzetMerchant true "page info"
// @Success 200 {object} reportDTO.ResInfoOmzetOutlet
// @Router /report/omzet-outlet [post]
func (handler *reportHandler) ReportOmzetOutlet(c *gin.Context) {
	var req reportModel.ReqOmzetMerchant
	errBind := c.BindJSON(&req)

	if errBind != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":   http.StatusBadRequest,
				"messages": "Please Contact Admin !!!, " + errBind.Error(),
			},
		)
		return
	}

	res, err := handler.reportUsecase.GetPaginationOmzetOutlet(&req)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":   http.StatusNotFound,
				"messages": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":   http.StatusOK,
			"messages": "Success",
			"result":   &res,
		},
	)
}
