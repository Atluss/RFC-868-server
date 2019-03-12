package utl

import (
	"log"
	"testing"
)

func TestAlphaOnly(t *testing.T) {
	vals := []string{"aaa", "123123", "12sfd"}

	for _, v := range vals {
		log.Printf("%t", DigitalOnly(v))
	}
}

func TestCheckServerSettings(t *testing.T) {

	args :=  map[uint][]string{
		0 : {"asd", "asd"},
		1 : {"-p", "11037"},
		2 : {"asd", "11037"},
		3 : {"-p", "as11037d"},
		4 : {"-p", "asd"},
	}

	for _, m := range args {
		if _, err := CheckServerSettings(m); err != nil {
			log.Println(m, " - not pass")
		} else {
			log.Println(m, " - pass")
		}
	}

}

func TestRFC868Time(t *testing.T) {
	log.Println(RFC868Time())
}
