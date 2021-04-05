package stats

import (
	"fmt"
	"github.com/SonnLarissa/bank/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			Amount: 10_000_00,
		},
		{
			Amount: 20_000_00,
		},
		{
			Amount: 10_000_00,
		},
		{
			Amount: 10_000_00,
		}}
	fmt.Println(Avg(payments))
	//Output 1250000
}
func ExampleTotalInCategoryAvg() {
	payments := []types.Payment{
		{
			Amount:   2_000_00,
			Category: "salom",
		},
		{
			Amount:   1_000_00,
			Category: "alek",
		},
		{
			Amount:   5_000_00,
			Category: "salom",
		},
		{
			Amount:   1_000_00,
			Category: "paka",
		},
	}
	fmt.Println(TotalInCategory(payments, "salom"))
	//Output: 700000
}