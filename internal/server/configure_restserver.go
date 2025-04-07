package server

import (
	"github.com/alireza-mht/tax-calculator/internal/log"
	"github.com/alireza-mht/tax-calculator/internal/server/api"
	"github.com/gin-gonic/gin"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) GetTaxCalculatorTaxYearYear(c *gin.Context, year int, salary api.GetTaxCalculatorTaxYearYearParams) {
	log.Debug("Received GET tax year request")
}
