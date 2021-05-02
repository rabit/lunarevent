package main

import (
	"fmt"
	"github.com/Lofanmi/chinese-calendar-golang/calendar"
	"github.com/jinzhu/now"
	"time"
)

func main() {
	var year int = 2021
	var month int
	var day int
	var trigger int = 24
	fmt.Println("BEGIN:VCALENDAR")
	fmt.Println("PRODID:-//Google Inc//Google Calendar 70.9054//EN")
	fmt.Println("VERSION:2.0")
	fmt.Println("CALSCALE:GREGORIAN")
	fmt.Println("METHOD:PUBLISH")
	fmt.Println("X-WR-CALNAME:農曆拜拜")
	fmt.Println("X-WR-TIMEZONE:Asia/Taipei")
	currTime := time.Now()
	for month = 1; month <= 12; month++ {
		t1 := time.Date(year, time.Month(month), 1, 12, 0, 0, 0, time.Now().Location())
		endOfMonth := now.With(t1).EndOfMonth()
		for day = 1; day <= endOfMonth.Day(); day++ {
			t2 := time.Date(year, time.Month(month), day, 12, 0, 0, 0, time.Now().Location())
			c := calendar.ByTimestamp(t2.Unix())
			if c.Lunar.GetDay() == 1 || c.Lunar.GetDay() == 15 {
				fmt.Println("\nBEGIN:VEVENT")
				//fmt.Printf("solar: %s\n", t2.Format("2006/01/02"))
				//fmt.Printf("lunar: %4d/%2d/%2d\n", c.Lunar.GetYear(), c.Lunar.GetMonth(), c.Lunar.GetDay())
				fmt.Printf("DTSTART:%sT040000Z\n", t2.Format("20060102"))
				fmt.Printf("DTEND:%sT050000Z\n", t2.Format("20060102"))
				fmt.Printf("DTSTAMP:%s\n", currTime.Format("20060102T150405Z"))
				fmt.Printf("CREATED:%s\n", currTime.Format("20060102T150405Z"))
				fmt.Printf("DESCRIPTION:\n")
				fmt.Printf("LAST-MODIFIED:%s\n", currTime.Format("20060102T150405Z"))
				fmt.Printf("LOCATION:\n")
				fmt.Printf("SEQUENCE:0\n")
				fmt.Printf("STATUS:CONFIRMED\n")
				if c.Lunar.GetDay() == 1 {
					fmt.Printf("SUMMARY:初一拜拜\n")
				} else {
					fmt.Printf("SUMMARY:十五拜拜\n")
				}
				fmt.Printf("TRANSP:OPAQUE\n")
				fmt.Printf("BEGIN:VALARM\n")
				fmt.Printf("ACTION:DISPLAY\n")
				fmt.Printf("DESCRIPTION:This is a lunar event reminder\n")
				fmt.Printf("TRIGGER:-P0DT%dH0M0S\n", trigger)
				fmt.Printf("END:VALARM\n")
				fmt.Printf("END:VEVENT\n")
			}
		}
	}
	fmt.Println("END:VCALENDAR")
}
