package utl

import (
	"fmt"
	"strings"
	"time"
)

const alpha = "abcdefghijklmnopqrstuvwxyz"

// DigitalOnly check symbols for a letters
func DigitalOnly(s string) bool {
	for _, char := range s {
		if strings.Contains(alpha, strings.ToLower(string(char))) {
			return false
		}
	}
	return true
}

func CheckServerSettings(args []string) (string, error){

	if len(args) != 2 {
		return "", fmt.Errorf("error: please input server port, example: 'server -p 11037'")
	}

	if args[0] != "-p" {
		return "", fmt.Errorf("error: please input correct argument, example: 'server -p 11037'")
	}

	if !DigitalOnly(args[1]) {
		return "", fmt.Errorf("error: please input correct argument port number, example: 'server -p 11037'")
	}

	return args[1], nil
}

func RFC868Time() uint32 {
	then := time.Date(1900, time.January, 1, 0, 0, 1, 0, time.UTC)
	now := time.Now()
	diff := now.Sub(then)
	return uint32(diff.Seconds()) & 0xFFFFFFFF
}