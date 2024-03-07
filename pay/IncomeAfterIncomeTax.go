package pay

import "math"

const expenseRatio = 0.1
const defaultIncomeTaxRate = 12.8
const socialTaxRate = 9.7

type IncomeDetail struct {
	WorkingDays      int     `json:"workingDays"`
	DailyRate        float64 `json:"dailyRate"`
	TotalGrossIncome float64 `json:"totalGrossIncome"`
	TaxableIncome    int64   `json:"taxableIncome"`
	SocialTaxAmount  int64   `json:"socialTaxAmount"`
	IncomeTaxAmount  int64   `json:"incomeTaxAmount"`
	NetIncome        int64   `json:"netIncome"`
}

func ComputeIncomeDetail(workedDays int, dailyRate float64, taxRate float64) IncomeDetail {
	if taxRate == 0 {
		taxRate = defaultIncomeTaxRate
	}
	grossIncome := ComputeGrossIncome(workedDays, dailyRate)
	return computeIncomeDetail(grossIncome, taxRate)
}

func computeIncomeDetail(grossIncomeByDailyRate GrossIncome, taxRate float64) IncomeDetail {
	grossIncomeAfterExpenses := int64(math.Round(grossIncomeByDailyRate.TotalGrossIncome * (1 - expenseRatio)))
	socialTaxAmount := int64(math.Round(float64(grossIncomeAfterExpenses)*socialTaxRate) / 100)
	incomeTaxAmount := int64(math.Round(float64(grossIncomeAfterExpenses) * defaultIncomeTaxRate / 100))
	netIncome := grossIncomeAfterExpenses - socialTaxAmount - incomeTaxAmount

	return IncomeDetail{grossIncomeByDailyRate.totalDays,
		grossIncomeByDailyRate.dailyPay, grossIncomeByDailyRate.TotalGrossIncome, grossIncomeAfterExpenses,
		socialTaxAmount, incomeTaxAmount, netIncome,
	}
}
