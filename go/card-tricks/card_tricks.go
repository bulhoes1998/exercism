package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2,6,9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if index < 0 || index > len(slice) - 1 {
		return -1
	}

	return slice[index]
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index < 0 || index > len(slice) - 1 {
		slice = append(slice, value)
	} else {
		slice[index] = value
	}

	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, value ...int) []int {
	if len(value) > 0 {
		return append(value, slice...)
	}

	return slice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(s []int, index int) []int {
	if index < 0 || index > len(s) - 1 {
		return s
	}

	return append(s[:index], s[index + 1:]...)
}
