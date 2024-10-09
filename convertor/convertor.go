package convertor

import (
	"base-convertor/util"
	"errors"
	"fmt"
	"unicode"
)

func ToDecimal(number string , base int)(int,error){
	intNumber := 0
	for _ ,digit := range number{
		tmp , err := digitBase(digit)
		if err != nil {
			return -1 , errors.New("invlaid number")
		}
		intNumber = intNumber * base + tmp
	}
	return intNumber,nil
}



func CheckBase(number string , base int)(bool , error){
	if base <= 1 || base > 36 {
		return false , errors.New("invalid Base")
	}

	numBase := 1
	for _ ,digit := range number{
		tmp , err := digitBase(digit)
		if err != nil {
			return false , errors.New("invlaid number")
		}
		numBase = util.Max(numBase,tmp)
	}

	fmt.Println("base is -> " , numBase)
	
	return numBase < base , nil
}

func digitBase(digit rune)(int,error){
	if '0' <= digit && digit <= '9' {
		return int(digit) - '0' , nil 
	}else if 'a' <= unicode.ToLower(digit) && unicode.ToLower(digit) <= 'z'{
		return int(unicode.ToLower(digit)) - 'a' + 10 , nil
	}
	return -1 , errors.New("invalid digit")
}