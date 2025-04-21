package types

import (
	"errors"
	"regexp"
)

type (
	Wallet string
)

func (w Wallet) String() string {
	return string(w)
}

func (w Wallet) Validate() error {
	regexPattern := `^0x[a-fA-F0-9]{40}$`

	regex := regexp.MustCompile(regexPattern)
	if regex.MatchString(w.String()) {
		return nil
	}
	return errors.New("Validation error: invalid wallet address")
}
