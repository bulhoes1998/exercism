package lasagna

import "fmt"

func PreparationTime(layers []string, prepTime int) int {
	if prepTime == 0 {
		return len(layers) * 2
	}

	return len(layers) * prepTime
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	
	noodles := 0
	sauce := 0.0

	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles += 50
		} else if layers[i] == "sauce" {
			sauce += 0.2
		}
	}

	return noodles, sauce
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList) - 1] = friendsList[len(friendsList) - 1]
	fmt.Println(myList)
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amountsNeeded []float64, portionsNumb int) []float64 {
	scaledQuantities := make([]float64, len(amountsNeeded))

	for i := 0; i < len(amountsNeeded); i++ {
		scaledQuantities[i] = (amountsNeeded[i]/2.0) * float64(portionsNumb)
	}
	return scaledQuantities
}