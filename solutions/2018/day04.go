package solutions

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pradeepg26/advent/util"
)

type DutyRoster struct {
	date   string
	guard  string
	awake  []bool
	asleep []bool
}

func Merge(log []DutyRoster) DutyRoster {
	awake := make([]bool, 1440)
	asleep := make([]bool, 1440)
	var date string
	var guard string
	for _, dr := range log {
		for i := 0; i < 1440; i++ {
			awake[i] = awake[i] || dr.awake[i]
			asleep[i] = asleep[i] || dr.asleep[i]
		}
		if dr.guard != "" {
			guard = dr.guard
		}
		date = dr.date
	}
	return DutyRoster{
		date:   date,
		guard:  guard,
		awake:  awake,
		asleep: asleep,
	}
}

// [1518-09-24 00:59] wakes up
// [1518-05-28 00:02] Guard #2719 begins shift
// [1518-03-31 00:42] falls asleep
// var re *regexp.Regexp = regexp.MustCompile(`\[(?P<date>\d{4}-\d{2}-\d{2})\] (?P<hour>\d{2}):(?P<minute>\d{2})] (?P<message>.*)`)
var re *regexp.Regexp = regexp.MustCompile(`\[(?P<date>\d{4}-\d{2}-\d{2}) (?P<hour>\d{2}):(?P<minute>\d{2})\] (?P<message>.*)`)

func ParseLine(s string) DutyRoster {
	extr := re.FindStringSubmatch(s)
	date := extr[1]
	message := extr[4]
	minute := util.ParseInt(extr[2])*24 + util.ParseInt(extr[3])
	dr := DutyRoster{
		date:   date,
		guard:  "",
		awake:  make([]bool, 1440),
		asleep: make([]bool, 1440),
	}
	if strings.Contains(message, "wakes up") {
		dr.awake[minute] = true
	} else if strings.Contains(message, "falls asleep") {
		dr.asleep[minute] = true
	} else {
		message = strings.TrimPrefix(message, "Guard #")
		message = strings.TrimSuffix(message, " begins shift")
		fmt.Printf("Trimmed down to '%s'\n", message)
		dr.guard = message
		dr.awake[0] = true
	}
	return dr
}

func LoadInput() []DutyRoster {
	log := make(map[string][]DutyRoster)
	for _, line := range util.LoadInputAsLines() {
		dr := ParseLine(line)
		log[dr.date] = append(log[dr.date], dr)
	}
	output := make([]DutyRoster, 0, 30)
	for _, val := range log {
		merged := Merge(val)
		status := make([]bool, 1440)
		isCurrentlyAsleep := false
		for i := 0; i < 1440; i++ {
			if merged.awake[i] {
				isCurrentlyAsleep = false
			}
			if merged.asleep[i] {
				isCurrentlyAsleep = true
			}
			status[i] = isCurrentlyAsleep
		}
		output = append(output, DutyRoster{
			date:   merged.date,
			guard:  merged.guard,
			asleep: status,
		})
	}
	return output
}

func Count(bitmap []bool) int {
	var count int
	for _, b := range bitmap {
		if b {
			count++
		}
	}
	return count
}

func main1804() {
	data := LoadInput()
	for _, dr := range data {
		fmt.Printf("Guard %s slept for %d minutes on %s\n", dr.guard, Count(dr.asleep), dr.date)
	}
}
