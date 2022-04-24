package techpalace

import (
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	starLine := ""
	for s := 0; s < numStarsPerLine; s++ {
		starLine += "*"
	}
	return starLine + "\n" + welcomeMsg + "\n" + starLine
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	msgWithoutStars := strings.ReplaceAll(oldMsg, "*", "")
	msgWithoutLineBreak := strings.ReplaceAll(msgWithoutStars, "\n", "")
	newMsg := strings.TrimLeft(msgWithoutLineBreak, " ")
	newMsg = strings.TrimRight(newMsg, " ")
	return newMsg
}
