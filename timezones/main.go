package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var timezones map[string]int

func init() {
	timezones = map[string]int{
		"UTC":  0,
		"GMT":  0,
		"BST":  60,
		"IST":  60,
		"WET":  0,
		"WEST": 60,
		"CET":  60,
		"CEST": 60 * 2,
		"EET":  60 * 2,
		"EEST": 60 * 3,
		"MSK":  60 * 3,
		"MSD":  60 * 4,
		"AST":  -60 * 4,
		"ADT":  -60 * 3,
		"NST":  -60*3 - 30,
		"NDT":  -60*2 - 30,
		"EST":  -60 * 5,
		"EDT":  -60 * 4,
		"CST":  -60 * 6,
		"CDT":  -60 * 5,
		"MST":  -60 * 7,
		"MDT":  -60 * 6,
		"PST":  -60 * 8,
		"PDT":  -60 * 7,
		"HST":  -60 * 10,
		"AKST": -60 * 9,
		"AKDT": -60 * 8,
		"AEST": 60 * 10,
		"AEDT": 60 * 11,
		"ACST": 60*9 + 30,
		"ACDT": 60*10 + 30,
		"AWST": 60 * 8,
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	times := strToInt(readLine(reader))
	for i := 0; i < times; i++ {
		minutes, from, to := readTask(reader)
		minutes = minutes - timezones[from] + timezones[to]
		for minutes < 0 {
			minutes += 24 * 60
		}
		minutes %= 24 * 60
		fmt.Println(minutesToTime(minutes))
	}
}

func minutesToTime(minutes int) string {
	if minutes == 0 {
		return "midnight"
	}

	if minutes == 12*60 {
		return "noon"
	}

	suffix := "a.m."
	if minutes > 12*60 {
		suffix = "p.m."
	}

	hours, mins := minutes/60, minutes%60
	if hours == 0 {
		hours = 12
	}
	if hours > 12 {
		hours -= 12
	}

	return fmt.Sprintf("%v:%02d %v", hours, mins, suffix)
}

func readTask(reader *bufio.Reader) (minutes int, from, to string) {
	parts := strings.Split(readLine(reader), " ")
	from, to = parts[len(parts)-2], parts[len(parts)-1]
	time := parts[0]
	switch time {
	case "noon":
		minutes = 12 * 60
	case "midnight":
		// nothing to do
	default:
		timeParts := strings.Split(time, ":")
		hours := strToInt(timeParts[0])
		minutes = strToInt(timeParts[1])
		switch day := parts[1]; day {
		case "a.m.":
			if hours == 12 {
				hours = 0
			}
		case "p.m.":
			if hours < 12 {
				hours += 12
			}
		}

		minutes += hours * 60
	}

	return minutes, from, to
}

func readLine(reader *bufio.Reader) string {
	str, err := reader.ReadString('\n')
	must(err)

	return str[:len(str)-1]
}

func strToInt(v string) int {
	vi, err := strconv.Atoi(v)
	must(err)

	return vi
}

func must(err error) {
	if err != nil {
		panic(err.Error())
	}
}
