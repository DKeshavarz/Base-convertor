package util

import (
	"errors"
	"unicode"
)

func CheckBase(number string , base int)(error){
	if base <= 1 || base > 36 {
		return errors.New("invalid Base")
	}

	number_base := 1
	for _ ,digit := range number{
		if '0' <= digit && digit <= '9' {
			number_base = max(int(digit) - '0' , number_base)
		}else if 'a' <= unicode.ToLower(digit) && unicode.ToLower(digit) <= 'z'{
			number_base = max(int(unicode.ToLower(digit)) - 'a' + 10 , number_base)
		}else{
			return errors.New("invalid Number")
		}
	}
	
	return nil 
}

func max(a , b int)int{
	if(a > b){
		return a
	}
	return b
}