package main

import (
	"fmt"
	"time"
)

func main() {
	// See the example for time.Format for a thorough description of how
	// to define the layout string to parse a time.Time value; Parse and
	// Format use the same model to describe their input and output.

	// longForm shows by example how the reference time would be represented in
	// the desired layout.
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ := time.Parse(longForm, "Feb 3, 2013 at 7:54pm (PST)")
	fmt.Println(t)

	// shortForm is another way the reference time would be represented
	// in the desired layout; it has no time zone present.
	// Note: without explicit zone, returns time in UTC.
	const shortForm = "2006-Jan-02"
	t, _ = time.Parse(shortForm, "2013-Feb-03")
	fmt.Println(t)

	// asctime <---> the number of seconds elapsed since January 1, 1970 UTC
	const testStr = "Tue Nov  3 09:49:37 2015"
	t, _ = time.Parse(time.ANSIC, testStr)
	secs := t.Unix()
	fmt.Println(t)
	fmt.Println(secs)

	ansicStr := time.Unix(secs, 0).UTC().Format(time.ANSIC)
	fmt.Println(ansicStr)
	fmt.Println(testStr)

	if testStr == ansicStr {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
