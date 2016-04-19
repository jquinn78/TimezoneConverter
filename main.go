/*

Name: Timezone Converter
Input: System time or user specifed time
Output: Converted times

TODO: improve the help
      add support for multiple inputs
	  add support to output the results to a csv file
	  add automation of ZONEINFO environment variable
*/

package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

//get time(s) from user
func parsetime(timeInput string, tzInput string, sTime string) {

	test := strings.Split(timeInput, ",")
	test1 := strings.Split(tzInput, ",")

	for _, inputs := range test {

		date, _ := time.Parse(time.RFC822, inputs)

		if inputs == "" {

			fmt.Println("System Time: " + sTime)

		} else {

			fmt.Println("\n" + "Inputed Time:" + date.Format(time.RFC822))
		}

		for _, tz := range test1 {
			newLocation, _ := time.LoadLocation(tz)

			if tz != "" && inputs == "" {

				date, _ := time.Parse(time.RFC822, sTime)
				fmt.Println("Converted Time: " + date.In(newLocation).Format(time.RFC822))

			} else if tz != "" && inputs != "" {
				fmt.Println("Converted Time: " + date.In(newLocation).Format(time.RFC822))
			}

		}

	}

}

func main() {

	t := time.Now()

	tFormat := t.Format(time.RFC822)

	//Define flags
	inputFlag := flag.String("time", "", "Input the times to convert (format: 02 Jan 06 15:04 MST)")
	timezoneFlag := flag.String("timezone", "", "Enter the time zones")

	//Parse the flags
	flag.Parse()

	parsetime(*inputFlag, *timezoneFlag, tFormat)

}
