package reuse

import (
	"bufio"
	"regexp"
	"strings"
)

type IPFreq struct {
	IP string
	Frequency int
}

func ParseLogForIP(l string) []IPFreq {

	re, _ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)

	matches := re.FindAllString(l, -1)

	for _, m := range matches {
		s := strings.Split(m, `.`)
	}

// https://stackoverflow.com/questions/33162449/iterate-over-multiline-string-in-go
	scanner := bufio.NewScanner(strings.NewReader(l))
	for scanner.Scan() {

	}
}