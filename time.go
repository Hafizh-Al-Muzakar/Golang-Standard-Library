package main

import (
	"fmt"
	"time"
)

func main() {

	// getting the current time
	now := time.Now()
	fmt.Println("Current Time:", now)

	// geting specific information about the time
	fmt.Println("Year:", now.Year())
	fmt.Println("Month:", now.Month())
	fmt.Println("Days:", now.Day())
	fmt.Println("Hour", now.Hour())
	fmt.Println("Minute:", now.Minute())
	fmt.Println("Second:", now.Second())
	fmt.Println("Nanosecond:", now.Nanosecond())

	//formating time into a string
	fmt.Println("Time in specific format:", now.Format("2006-01-02 15:04:05"))

	//adding duration to time
	future := now.Add(time.Hour * 24)
	fmt.Println("Time 24 hours from now:", future.Format("2006-01-02 15:04:05"))

	//calculating between two time
	duration := future.Sub(now)
	fmt.Println("Duration betweem now and 24 hour from now:", duration)
}
