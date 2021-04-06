package stats

import (
	"github.com/SonnLarissa/bank/v2/pkg/types"
)

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

//FilterByCategory возвращяет платежи в указаной категории
func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment {
	var filtered []types.Payment
	for _, payment := range payments {
		if payment.Category == category {
			filtered = append(filtered, payment)
		}
	}
	return filtered
}

//CategoriesTotal возвращает сумму платежей по каждой категории
func CategoriesTotal(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}
	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
	}
	return categories
}

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	r := make(map[types.Category]types.Money)
	filter := make(map[types.Category][]types.Payment)
	for _, payment := range payments {
		if _, res := filter[payment.Category]; res {
			filter[payment.Category] = append(filter[payment.Category], payment)
		} else {
			filter[payment.Category] = []types.Payment{payment}
		}
	}
	f := func(payments []types.Payment) (types.Money, int) {
		count := 0
		m := types.Money(0)
		for _, payment := range payments {
			count++
			m += payment.Amount
		}
		return m, count
	}

	for k, v := range filter {
		m, c := f(v)
		r[k] = m / types.Money(c)
	}
	return r
}

func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money {
	res := make(map[types.Category]types.Money)
	if len(first) >= 0 && len(second) >= 0 {
		for ind, current := range second {
			res[ind] = current - first[ind]
		}
	}
	for k, v := range first {
		if _, ok := second[k]; !ok {
			res[k] = -v
		}
	}
	return res
}
