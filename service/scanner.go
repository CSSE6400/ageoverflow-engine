package service

import (
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"strings"
)

type Scanner struct {
}

func NewScanner() Scanner {
	return Scanner{}
}

func (d Scanner) Spin(cost int) error {
	_, err := bcrypt.GenerateFromPassword([]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ"), cost)
	return err
}

func (d Scanner) ScanEmail(request Request) (*Report, error) {
	seed := request.Metadata
	isMalicious := strings.Split(seed, "|")[0] == "1"
	cost, err := strconv.ParseInt(strings.Split(seed, "|")[1], 10, 64)
	if err != nil {
		return nil, err
	}

	// we are definitely doing a lot of import email scanning here
	err = d.Spin(int(cost))
	if err != nil {
		return nil, err
	}

	report := Report{
		ID:        request.ID,
		Malicious: isMalicious,
	}

	return &report, nil
}
