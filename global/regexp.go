package global

import "regexp"

var (
	// 'CMD:STATS' IN BYTES
	AvailableMessageRegexp = regexp.MustCompile(`(DATA:|CMD:STATS).*?(\\0x00){4}`) //GET MESSAGES WITH PREFIX DATA: OR CMD:STATS AND END WITH \0x00\0x00\0x00\0x00
	AvailableOcurrences    = regexp.MustCompile(`\d{4}_[A-z]{4,6}\b`)              //GET OCCURRENCES WITH 4 DIGITS, UNDERSCORE, 4-6 LETTERS
)
