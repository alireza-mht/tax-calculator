package tax

import (
	"testing"

	"github.com/alireza-mht/tax-calculator/internal/common"
	"github.com/alireza-mht/tax-calculator/internal/server/api"
	"github.com/stretchr/testify/assert"
)

func TestComputeTaxBreakdown(t *testing.T) {
	tests := []struct {
		name         string
		taxInfo      *taxInfo
		salary       float32
		wantTax      api.IncomeTax
		wantInternal bool // Expect InternalError
		wantBadReq   bool // Expect BadRequestError
	}{
		{
			name: "Single bracket",
			taxInfo: &taxInfo{
				TaxBrackets: []taxBracketInfo{
					{Max: 0, Min: 0, Rate: 0.1},
				},
			},
			salary:       10000,
			wantTax:      api.IncomeTax{},
			wantInternal: true,
		},
		{
			name: "Multiple brackets",
			taxInfo: &taxInfo{
				TaxBrackets: []taxBracketInfo{
					{Max: 10000, Min: 0, Rate: 0.1},
					{Max: 30000, Min: 10000, Rate: 0.2},
					{Max: 0, Min: 30000, Rate: 0.3},
				},
			},
			salary: 50000,
			wantTax: api.IncomeTax{
				TotalTax: 11000, // 1000 (10k @ 10%) + 4000 (20k @ 20%) + 6000 (20k @ 30%)
				TaxPerBand: []api.TaxBrackets{
					{Max: 10000, Min: 0, Rate: 0.1, Tax: 1000},
					{Max: 30000, Min: 10000, Rate: 0.2, Tax: 4000},
					{Max: 50000, Min: 30000, Rate: 0.3, Tax: 6000},
				},
				EffectiveRate: 0.22, // 11000 / 50000
			},
			wantInternal: false,
		},
		{
			name:         "Nil taxInfo",
			taxInfo:      nil,
			salary:       5000,
			wantTax:      api.IncomeTax{},
			wantInternal: true,
		},
		{
			name: "Negative Max value",
			taxInfo: &taxInfo{
				TaxBrackets: []taxBracketInfo{
					{Max: -10000, Min: 0, Rate: 0.1},
				},
			},
			salary:       5000,
			wantTax:      api.IncomeTax{},
			wantInternal: true,
		},
		{
			name: "Negative Min value",
			taxInfo: &taxInfo{
				TaxBrackets: []taxBracketInfo{
					{Max: 10000, Min: -5000, Rate: 0.1},
				},
			},
			salary:       5000,
			wantTax:      api.IncomeTax{},
			wantInternal: true,
		},
		{
			name: "Negative Rate value",
			taxInfo: &taxInfo{
				TaxBrackets: []taxBracketInfo{
					{Max: 10000, Min: 0, Rate: -0.1},
				},
			},
			salary:       5000,
			wantTax:      api.IncomeTax{},
			wantInternal: true,
		},
		{
			name: "Max equals Min",
			taxInfo: &taxInfo{
				TaxBrackets: []taxBracketInfo{
					{Max: 10000, Min: 10000, Rate: 0.1},
				},
			},
			salary:       15000,
			wantTax:      api.IncomeTax{},
			wantInternal: true,
		},
		{
			name: "Max less than Min",
			taxInfo: &taxInfo{
				TaxBrackets: []taxBracketInfo{
					{Max: 5000, Min: 10000, Rate: 0.1},
				},
			},
			salary:       15000,
			wantTax:      api.IncomeTax{},
			wantInternal: true,
		},
		{
			name: "Empty brackets",
			taxInfo: &taxInfo{
				TaxBrackets: []taxBracketInfo{},
			},
			salary: 10000,
			wantTax: api.IncomeTax{
				TotalTax:      0,
				TaxPerBand:    []api.TaxBrackets{},
				EffectiveRate: 0,
			},
			wantInternal: false,
		},
		{
			name: "Zero salary",
			taxInfo: &taxInfo{
				TaxBrackets: []taxBracketInfo{
					{Max: 10000, Min: 0, Rate: 0.1},
				},
			},
			salary: 0,
			wantTax: api.IncomeTax{
				TotalTax:      0,
				TaxPerBand:    []api.TaxBrackets{},
				EffectiveRate: 0,
			},
			wantInternal: false,
		},
		{
			name: "Salary below bracket minimum",
			taxInfo: &taxInfo{
				TaxBrackets: []taxBracketInfo{
					{Max: 20000, Min: 10000, Rate: 0.2},
				},
			},
			salary: 5000,
			wantTax: api.IncomeTax{
				TotalTax:      0,
				TaxPerBand:    []api.TaxBrackets{},
				EffectiveRate: 0,
			},
			wantInternal: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ComputeTaxBreakdown(tt.taxInfo, tt.salary)

			if tt.wantInternal || tt.wantBadReq {
				assert.Error(t, err)
				assert.Equal(t, tt.wantInternal, common.IsErrInternal(err), "InternalError mismatch")
				assert.Equal(t, tt.wantTax, got, "IncomeTax should match expected on error")
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.wantTax.TotalTax, got.TotalTax, "TotalTax mismatch")
			assert.Equal(t, tt.wantTax.EffectiveRate, got.EffectiveRate, "EffectiveRate mismatch")
			assert.Equal(t, len(tt.wantTax.TaxPerBand), len(got.TaxPerBand), "TaxPerBand length mismatch")

			for i := range tt.wantTax.TaxPerBand {
				assert.Equal(t, tt.wantTax.TaxPerBand[i].Max, got.TaxPerBand[i].Max, "Max mismatch")
				assert.Equal(t, tt.wantTax.TaxPerBand[i].Min, got.TaxPerBand[i].Min, "Min mismatch")
				assert.Equal(t, tt.wantTax.TaxPerBand[i].Rate, got.TaxPerBand[i].Rate, "Rate mismatch")
				assert.Equal(t, tt.wantTax.TaxPerBand[i].Tax, got.TaxPerBand[i].Tax, "Tax mismatch")
			}
		})
	}
}
