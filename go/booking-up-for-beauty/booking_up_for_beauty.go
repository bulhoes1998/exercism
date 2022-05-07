package booking

import "time"

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	const layout = "1/2/2006 15:04:05"

	tm, _ := time.Parse(layout, date)
	return tm
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	dateF, _ := time.Parse(layout, date)

	tNow := time.Now()

	return tNow.After(dateF)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	dateF, _ := time.Parse(layout, date)

	if dateF.Hour() >= 12 && dateF.Hour() < 18 {
		return true
	}

	return false
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	dateF, _ := time.Parse(layout, date)

	descriptionDate := dateF.Format("Monday, January 2, 2006, at 15:04")

	description := "You have an appointment on " + descriptionDate + "."

	return description

}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	layout := "2006-01-02 15:04:05"
	anniversary := "2020-09-15 00:00:00"

	anniversaryDate, _ := time.Parse(layout, anniversary)

	anniversaryDate = anniversaryDate.AddDate(time.Now().Year() - anniversaryDate.Year(), 0, 0)

	return anniversaryDate
}
