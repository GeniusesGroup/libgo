//go:build lang_per

/* For license and copyright information please see the LEGAL file in the code repository */

package utc

// Check given day exist in desire Weekdays.
func (wd Weekdays) String() (day string) {
	switch wd {
	case Weekdays_Monday:
		return "دوشنبه"
	case Weekdays_Tuesday:
		return "سه شنبه"
	case Weekdays_Wednesday:
		return "چهارشنبه"
	case Weekdays_Thursday:
		return "پنچ شنبه"
	case Weekdays_Friday:
		return "جمعه"
	case Weekdays_Saturday:
		return "شنبه"
	case Weekdays_Sunday:
		return "یکشنبه"
	}
	return
}
