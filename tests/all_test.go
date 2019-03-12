package tests

import (
	"log"
	"testing"
	"timeServer/utl"
)

func TestAlphaOnly(t *testing.T) {
	vals := []string{"aaa", "123123", "12sfd"}

	for _, v := range vals {
		log.Printf("%t", utl.DigitalOnly(v))
	}
}

func TestCheckServerSettings(t *testing.T) {

	args := map[uint][]string{
		0: {"asd", "asd"},
		1: {"-p", "11037"},
		2: {"asd", "11037"},
		3: {"-p", "as11037d"},
		4: {"-p", "asd"},
	}

	for _, m := range args {
		if _, err := utl.CheckServerSettings(m); err != nil {
			log.Println(m, " - not pass")
		} else {
			log.Println(m, " - pass")
		}
	}

}

func TestRFC868Time(t *testing.T) {
	log.Println(utl.RFC868Time())
}

func TestREFC868TimeToUnix(t *testing.T) {
	log.Println(utl.REFC868TimeToUnix(3761408855))
}

func TestDialToTimeServer(t *testing.T) {
	address := "time.nist.gov:37"
	if str, err := utl.DialToTimeServer(address); err != nil {
		log.Println("No Pass")
	} else {
		log.Printf("Pass, %s", str)
	}
}
