package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	measurements := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}

	return measurements
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value, exists := units[unit]

	if !exists {
		return false
	} else {
		bill[item] += value
		return true
	}

}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	value1, exists1 := bill[item]
	value2, exists2 := units[unit]
	if !exists1 {
		return false
	} else if !exists2 {
		return false
	} else if value1 - value2 < 0 {
		return false
	} else if value1 - value2 == 0 {
		delete(bill, item)
		return true
	} else {
		bill[item] -= value2
		return true
	}

}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value, exists := bill[item]

	return value, exists
}
