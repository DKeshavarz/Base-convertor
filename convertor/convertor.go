package convertor

import (
	"base-convertor/util"
	"errors"
	"strconv"
	"unicode"
)

func ToDecimal(number string , base int)(int,error){
	intNumber := 0
	for _ ,digit := range number{
		tmp , err := digitToNumber(digit)
		if err != nil {
			return -1 , errors.New("invlaid number")
		}
		intNumber = intNumber * base + tmp
	}
	return intNumber,nil
}

func FromDecimal(number , newBase int)(string,error){
	str := ""
	for number > 0 {
		tmp := number % newBase
		newDig , err := numberToDigit(tmp)
		if err != nil {
			return "" , err
		}
		str = newDig + str
		number = number / newBase
	}
	return str , nil
}


func CheckBase(number string , base int)(bool , error){
	if base <= 1 || base > 36 {
		return false , errors.New("invalid Base")
	}

	numBase := 1
	for _ ,digit := range number{
		tmp , err := digitToNumber(digit)
		if err != nil {
			return false , errors.New("invlaid number")
		}
		numBase = util.Max(numBase,tmp)
	}
	
	return numBase < base , nil
}

func digitToNumber(digit rune)(int,error){
	if '0' <= digit && digit <= '9' {
		return int(digit) - '0' , nil 
	}else if 'a' <= unicode.ToLower(digit) && unicode.ToLower(digit) <= 'z'{
		return int(unicode.ToLower(digit)) - 'a' + 10 , nil
	}
	return -1 , errors.New("invalid digit")
}

func numberToDigit(number int)(string,error){
	if 0 <= number && number <= 9 {
		str := strconv.Itoa(number)
		return str , nil 
	} 
	if 10 <= number && number <= 28 {
		str := []rune{'A'}
		str[0] += rune(number) - 10

		return string(str) , nil
	}

	return "" , errors.New("invliad number")
}