package tax

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/alireza-mht/tax-calculator/internal/common"
	"github.com/alireza-mht/tax-calculator/internal/server/api"
)

// taxInfo represents tax bracket information for a specific year
type taxInfo struct {
	TaxBrackets []taxBracketInfo `json:"tax_brackets"`
}

// taxBracketInfo defines the tax details for a specific income bracket
type taxBracketInfo struct {
	Max  float32 `json:"max,omitempty"`
	Min  float32 `json:"min"`
	Rate float32 `json:"rate"`
}

// CalculateIncomeTax computes the income tax for a given year and salary
func CalculateIncomeTax(year int, salary float32) (api.IncomeTax, error) {
	var incomeTax api.IncomeTax
	taxInfo, err := FetchTaxYearInfo(year)
	if err != nil {
		return incomeTax, &common.InternalError{Details: fmt.Sprintf("failed to fetch the tax year information: %s", err)}
	}

	return ComputeTaxBreakdown(taxInfo, salary)
}

// FetchTaxYearInfo retrieves tax bracket information for a specific year from an external service
func FetchTaxYearInfo(year int) (*taxInfo, error) {
	externalServiceUrl := "localhost:5001/tax-calculator/tax-year/" + strconv.Itoa(year)
	resp, err := common.HttpRequestWithResponse(externalServiceUrl, http.MethodGet, "", "")
	if err != nil {
		return nil, &common.InternalError{Details: fmt.Sprintf("failed to get the tax info from %s: %s", externalServiceUrl, err)}
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, &common.InternalError{Details: fmt.Sprintf("failed to read response body: %s", err)}
	}

	// Check the error code
	if resp.StatusCode != http.StatusOK {
		return nil, &common.InternalError{Details: fmt.Sprintf("Failed to get the response from remote tax info container. Request returned status code %d, with body: %s", resp.StatusCode, string(body))}
	}

	var taxInfo taxInfo
	if err := json.Unmarshal(body, &taxInfo); err != nil {
		return nil, &common.InternalError{Details: fmt.Sprintf("failed to unmarshal response: %s", err)}
	}
	return &taxInfo, nil
}

// ComputeTaxBreakdown calculates the detailed tax breakdown based on salary and tax brackets
func ComputeTaxBreakdown(taxInfo *taxInfo, salary float32) (api.IncomeTax, error) {
	var incomeTax api.IncomeTax
	if taxInfo == nil {
		return incomeTax, &common.InternalError{Details: "failed to compute the tax breakdown because taxInfo is nil"}
	}
	for _, bracket := range taxInfo.TaxBrackets {
		// Validate the bracket
		if bracket.Max < 0 || bracket.Min < 0 || bracket.Rate < 0 || bracket.Max == bracket.Min {
			return incomeTax, &common.InternalError{Details: "failed to compute the tax breakdown because the parameters are invalid"}
		}

		// Ensure Max >= Min when Max is non-zero
		if bracket.Max != 0 && bracket.Max < bracket.Min {
			return incomeTax, &common.InternalError{Details: "failed to compute the tax breakdown because max is less than min"}
		}

		// Skip brackets where salary is below the minimum
		if salary <= bracket.Min {
			continue
		}

		// Calculate upper and min
		upper := salary
		if bracket.Max != 0 && bracket.Max < salary {
			upper = bracket.Max
		}

		bracketAmount := upper - bracket.Min
		bracketTax := bracketAmount * bracket.Rate
		incomeTax.TotalTax += bracketTax

		bracketInfo := api.TaxBrackets{
			Max:  upper,
			Min:  bracket.Min,
			Rate: bracket.Rate,
			Tax:  bracketTax,
		}
		incomeTax.TaxPerBand = append(incomeTax.TaxPerBand, bracketInfo)
	}
	if salary != 0 {
		incomeTax.EffectiveRate = incomeTax.TotalTax / salary
	} else {
		incomeTax.EffectiveRate = 0
	}

	return incomeTax, nil
}
