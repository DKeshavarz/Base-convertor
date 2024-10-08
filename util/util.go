package util

import (
	"errors"
	"regexp"
)

func CheckBase(number string , base int)(error){
	if base <= 1 || base > 36 {
		return errors.New("invalid Base")
	}

	regex := regexp.MustCompile(`^[0-6]+$`)

	if(!regex.MatchString(number)){
		return errors.New("invalid char")
	}

	
	return nil 
}