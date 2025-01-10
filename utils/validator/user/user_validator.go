package validator_util

import (
	"errors"
	"strings"
	"time"
)

func ValidateUsername(username string) error {
	if username == "" {
		return errors.New("username cannot be empty")
	}
	if strings.Contains(username, " ") {
		return errors.New("username cannot contain spaces")
	}
	return nil
}

func ValidateEmail(email string) error {
	if strings.Contains(email, " ") {
		return errors.New("email cannot contain spaces")
	}
	if email == "" {
		return errors.New("email cannot be empty")
	}
	if !strings.Contains(email, "@") {
		return errors.New("email must contain @")
	}
	return nil
}

func ValidatePassword(password string) error {
	if password == "" {
		return errors.New("password cannot be empty")
	}
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	if strings.Contains(password, " ") {
		return errors.New("password cannot contain spaces")
	}
	return nil
}

func ValidateNIK(nik string) error {
	if nik == "" {
		return errors.New("nik cannot be empty")
	}
	if len(nik) != 16 {
		return errors.New("nik must be 16 characters long")
	}
	return nil
}

func ValidateFullname(fullname string) error {
	if fullname == "" {
		return errors.New("fullname cannot be empty")
	}
	return nil
}

func ValidateLegalname(legalname string) error {
	if legalname == "" {
		return errors.New("legalname cannot be empty")
	}
	return nil
}

func ValidateBirthplace(birthplace string) error {
	if birthplace == "" {
		return errors.New("birthplace cannot be empty")
	}
	return nil
}

func ValidateBirthdate(birthdate string) error {
	if birthdate == "" {
		return errors.New("birthdate cannot be empty")
	}

	layout := "02-01-2006"
	_, err := time.Parse(layout, birthdate)
	if err != nil {
		return errors.New("birthdate must be in the format DD-MM-YYYY")
	}

	return nil
}
