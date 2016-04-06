/*

Name: Timezone Converter
Input: System time or user specifed time
Output: Converted times

TODO: improve the help,
      add support for multiple inputs
	  add support to output the results to a csv file
	  add automation of ZONEINFO environment variable
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
	inputFlag := flag.String("time", tFormat, "Input the times to convert (format: 02 Jan 06 15:04 MST)")
	timezoneFlag := flag.String("timezone", "", "Enter the time zones")

	//Parse the flags
	flag.Parse()

	newLocation, _ := time.LoadLocation(*timezoneFlag) //user inputted timezone

	date, _ := time.Parse(time.RFC822, *inputFlag)

	//Return current day and time and exit if no input is specified
	if *inputFlag == "" {
		fmt.Println("System time: " + tFormat) //currentTime
		return
	}

	if *timezoneFlag == "" { //return current day and time and exit if timezoneFlag is not specified
		fmt.Println("System time: " + tFormat)
		return
	} else if *inputFlag == "" && *timezoneFlag != "" { //return the converted system time if the time is not specifed, but the timezone is specified
		fmt.Println("System time: " + date.In(newLocation).Format(time.RFC822)) //time in new location and formatted
	} else if *inputFlag != "" && *timezoneFlag != "" { //return the inputted day and time and the converted day and time if both flags contain input
		fmt.Println("Inputed time: " + *inputFlag)                                 //time inputted by user
		fmt.Println("Converted time: " + date.In(newLocation).Format(time.RFC822)) //time in new location and formatted

	}

}
