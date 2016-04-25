/*

Name: Timezone Converter
Input: System time or user specifed time
Output: Converted times

*/

package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

//save the times and converstions to a csv file
func createCSV(timeInput string, tzInput string, sTime string) {

	file, _ := os.Create("results.csv")
	defer file.Close()

	writer := csv.NewWriter(file)

	headers := []string{"Time", "Converted Time"}

	writer.Write(headers)

	splitTime := strings.Split(timeInput, ",")
	splitTimezone := strings.Split(tzInput, ",")

	//Iterate through the user inputed time(s) or system time and
	//save them and the converted times to results.csv.
	for _, times := range splitTime {

		date, _ := time.Parse(time.RFC822Z, times)

		for _, timezones := range splitTimezone {

			location, _ := time.LoadLocation(timezones)

			if times != "" && timezones != "" {

				writer.Write([]string{times, date.In(location).Format(time.RFC822)})
			} else if times == "" && timezones != "" {

				date, _ := time.Parse(time.RFC822Z, sTime)
				writer.Write([]string{sTime, date.In(location).Format(time.RFC822)})
			} else if times == "" && timezones == "" {
				writer.Write([]string{sTime})
			} else if times != "" && timezones == "" {
				writer.Write([]string{times, "No time zone specified"})
			}
		}

	}

	defer writer.Flush()

}

//get time(s) from user
func parsetime(timeInput string, tzInput string, sTime string) {

	test := strings.Split(timeInput, ",")
	test1 := strings.Split(tzInput, ",")

	for _, inputs := range test {

		date, _ := time.Parse(time.RFC822Z, inputs)

		if inputs == "" {

			fmt.Println("System Time: " + sTime)

		} else {

			fmt.Println("\n" + "Inputted Time:" + date.Format(time.RFC822Z))

		}

		for _, tz := range test1 {
			newLocation, _ := time.LoadLocation(tz)

			if tz != "" && inputs == "" {

				date, _ := time.Parse(time.RFC822Z, sTime)

				fmt.Println("Converted Time: " + date.In(newLocation).Format(time.RFC822))

			} else if tz != "" && inputs != "" {
				fmt.Println("Converted Time: " + date.In(newLocation).Format(time.RFC822))

			}

		}

	}

}

func main() {

	t := time.Now()

	tFormat := t.Format(time.RFC822Z)

	//Define flags
	inputFlag := flag.String("time", "", "Input the times to convert (format: 02 Jan 06 15:04 -0700). See https://en.wikipedia.org/wiki/List_of_tz_database_time_zones for the offsets.")
	timezoneFlag := flag.String("timezone", "", "Enter the time zones. The timezones must be in the form of Country/City (ex. America/Chicago). See https://en.wikipedia.org/wiki/List_of_tz_database_time_zones for the IANA conversions.")
	csvFlag := flag.Bool("csv", false, "Save data to csv. Append -csv to the end to save the inputted times and converted times to a CSV file.")

	//Parse the flags
	flag.Parse()

	parsetime(*inputFlag, *timezoneFlag, tFormat)

	if *csvFlag == true {
		createCSV(*inputFlag, *timezoneFlag, tFormat)
	}

}
