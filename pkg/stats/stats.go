package stats

import "github.com/SonnLarissa/bank/pkg/bank/types"

// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	result := types.Money(0)
	for _, item := range payments {
		result += item.Amount
	}
	return result / types.Money(len(payments))
}
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)
	for _, p := range payments {
		if p.Category != category {
			continue
		}
		sum += p.Amount
	}
	return sum
}