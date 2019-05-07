package v1

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"strings"
	"time"
)

const (
	alpha    = "abcdefghijklmnopqrstuvwxyz"
	ConnHost = "localhost" // connection host
	ConnType = "tcp"       // network type
)

var then1900 = time.Date(1900, time.January, 1, 0, 0, 1, 0, time.UTC)
var then1970 = time.Date(1970, time.January, 1, 0, 0, 1, 0, time.UTC)

// diffSec it's 2 208 988 800 but calculate for clarity
var diffSec = uint32(then1970.Sub(then1900).Seconds())

// DigitalOnly check symbols for a letters
func DigitalOnly(s string) bool {
	for _, char := range s {
		if strings.Contains(alpha, strings.ToLower(string(char))) {
			return false
		}
	}
	return true
}

// CheckServerSettings when run server, example: client -p 11037
func CheckServerSettings(args []string) (string, error) {

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

// CheckClientSettings when run client example: client localhost 11037
func CheckClientSettings(args []string) (string, string, error) {
	if len(args) != 2 {
		return "", "", fmt.Errorf("error: please input server address and port, example: 'client localhost 11037'")
	}

	if args[0] == "" {
		return "", "", fmt.Errorf("error: please input correct address argument, example: 'client localhost 11037'")
	}

	if !DigitalOnly(args[1]) {
		return "", "", fmt.Errorf("error: please input correct argument port number, client: 'server localhost 11037'")
	}

	return args[0], args[1], nil
}

// RFC868Time seconds from 1 jun 1970
func RFC868Time() uint32 {
	now := time.Now()
	diff := now.Sub(then1900)
	return uint32(diff.Seconds())
}

// REFC868TimeToUnix fix time from 1 jun 1970 to Unix timestamp
func REFC868TimeToUnix(secondsLeft uint32) uint32 {
	return secondsLeft - diffSec
}

// DialToTimeServer make a dial to time server
func DialToTimeServer(address string) (string, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return "", fmt.Errorf("error to connect: %s", err)
	}

	if _, err := fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n"); err != nil {
		return "", fmt.Errorf("error: %s", err)
	}

	buf := make([]byte, 4)
	var status []byte
	if status, _, err = bufio.NewReader(conn).ReadLine(); err != nil {
		return "", err
	}

	copy(buf, status)
	str := fmt.Sprintf("respond: %d", REFC868TimeToUnix(binary.BigEndian.Uint32(buf)))
	return str, nil
}
