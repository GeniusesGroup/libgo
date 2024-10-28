//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package utc

// Check given day exist in desire Weekdays.
func (wd Weekdays) String() (day string) {
	switch wd {
	case Weekdays_Monday:
		return "Monday"
	case Weekdays_Tuesday:
		return "Tuesday"
	case Weekdays_Wednesday:
		return "Wednesday"
	case Weekdays_Thursday:
		return "Thursday"
	case Weekdays_Friday:
		return "Friday"
	case Weekdays_Saturday:
		return "Saturday"
	case Weekdays_Sunday:
		return "Sunday"
	}
	return
}
