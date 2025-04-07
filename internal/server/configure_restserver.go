package server

import (
	"fmt"
	"net/http"

	"github.com/alireza-mht/tax-calculator/internal/common"
	"github.com/alireza-mht/tax-calculator/internal/log"
	"github.com/alireza-mht/tax-calculator/internal/server/api"
	"github.com/alireza-mht/tax-calculator/internal/tax"
	"github.com/gin-gonic/gin"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) GetTaxCalculatorTaxYearYear(c *gin.Context, year int, salary api.GetTaxCalculatorTaxYearYearParams) {
	log.Debug("Received GET tax year request")

	taxInfo, err := tax.CalculateIncomeTax(year, salary.Salary)
	if err != nil {
		log.Error(fmt.Sprintf("GetTaxCalculatorTaxYearYear returned error: %s", err.Error()))
		switch {
		case common.IsErrBadRequest(err):
			c.JSON(http.StatusBadRequest, infoErr(http.StatusText(http.StatusBadRequest), err.Error()))
		case common.IsErrNotFound(err):
			c.JSON(http.StatusNotFound, infoErr(http.StatusText(http.StatusNotFound), err.Error()))
		default:
			c.JSON(http.StatusInternalServerError, infoErr(http.StatusText(http.StatusInternalServerError), err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, taxInfo)
}

func infoErr(status, message string) api.ErrorResponse {
	return api.ErrorResponse{
		Error:   status,
		Message: message,
	}
}
