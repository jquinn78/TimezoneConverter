# TimezoneConverter

Converts the time specified by the user and outputs the time into the specified time zones.

## Usage ##

-csv
    Save data to csv. Append -csv to the end to save the inputted times and converted times to a CSV file.
    	
-time string
    Input the times to convert (format: 02 Jan 06 15:04 -0700). See https://en.wikipedia.org/wiki/List_of_tz_database_time_zones for the offsets.
    	
-timezone string
    Enter the time zones. The timezones must be in the form of Country/City (ex. America/Chicago). See https://en.wikipedia.org/wiki/List_of_tz_database_time_zones for the IANA conversions.

**Linux Usage Examples**
*Convert the times into CST/CDT and EST/EDT*
TimezoneConverter -time="02 Jan 06 15:04 -0700","03 Mar 16 13:00 -0500" -timezone="America/Chicago","America/New_York"

*Convert the system time into CDT/CST and EST/EDT*
TimezoneConverter --timezone="America/Chicago","America/New_York"

*Convert the times into CST/CDT and EST/EDT and save them to a csv file*
TimezoneConverter -time="02 Jan 06 15:04 -0700","03 Mar 16 13:00 -0500" -timezone="America/Chicago","America/New_York" -csv

**Windows Usage Examples**
*Convert the times into CST/CDT and EST/EDT*
TimezoneConverter_win32.exe -time="02 Jan 06 15:04 -0700","03 Mar 16 13:00 -0500" -timezone="America/Chicago","America/New_York"

*Convert the system time into CDT/CST and EST/EDT*
TimezoneConverter_win32.exe --timezone="America/Chicago","America/New_York"

*Convert the times into CST/CDT and EST/EDT and save them to a csv file*
TimezoneConverter_win32.exe -time="02 Jan 06 15:04 -0700","03 Mar 16 13:00 -0500" -timezone="America/Chicago","America/New_York" -csv

To use the tool on Windows, download zoneinfo.zip and set an environment variable called ZONEINFO pointing at the zoneinfo.zip file. To set the environment variable temporarily issue

set ZONEINFO=C:\<directory that contains zoneinfo.zip>\zoneinfo.zip

at the command line.

To set the environment variable permanently issue

set ZONEINFO=c:\<directory that contains zoneinfo.zip>\zoneinfo.zip

at the command line.
