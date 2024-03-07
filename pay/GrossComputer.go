package pay

var defaultDailyRate = 500.0

type GrossIncome struct {
	totalDays        int
	dailyPay         float64
	TotalGrossIncome float64
}

func ComputeGrossIncome(totalDays int, dailyRate float64) GrossIncome {
	if dailyRate == 0 {
		dailyRate = defaultDailyRate
	}
	income := float64(totalDays) * dailyRate
	return GrossIncome{totalDays: totalDays, dailyPay: dailyRate, TotalGrossIncome: float64(income)}
}
