/*

Name: Timezone Converter
Input: Current time or user specifed time
Output: Converted times

Going to have to make something to convert the time zones to short form (sucks),

*/

package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {

	t := time.Now()

	tFormat := t.Format(time.RFC822)

	//Define flags
	nowFlag := flag.String("now", tFormat, "The current time (based on system time)")
	inputFlag := flag.String("time", "", "Input the times to convert (format: 02 Jan 06 15:04 MST)")
	timezoneFlag := flag.String("timezone", "America/Chicago", "Enter the time zones")

	//Parse the flags
	flag.Parse()

	newLocation, _ := time.LoadLocation(*timezoneFlag) //user inputted timezone

	date, _ := time.Parse(time.RFC822, *inputFlag)

	fmt.Println(*nowFlag)                              //currentTime
	fmt.Println(t.In(newLocation).Format(time.RFC822)) //formatted local time in specifed location

	fmt.Println(*inputFlag)                               //time inputted by user
	fmt.Println(date.In(newLocation).Format(time.RFC822)) //time in new location and formatted

}
