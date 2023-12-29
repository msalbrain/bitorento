package decode

import (
	

	"fmt"
	"strconv"
	"unicode"
	
)

func decodeString(bencodedString string) (interface{}, int, error) {
	var firstColonIndex int

	for i := 0; i < len(bencodedString); i++ {
		if bencodedString[i] == ':' {
			firstColonIndex = i
			break
		}
	}

	lengthStr := bencodedString[:firstColonIndex]
	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		return "", length, err
	}

	return bencodedString[firstColonIndex+1 : firstColonIndex+1+length], firstColonIndex + 1 + length, nil
}

func decodeInteger(bencodedString string) (interface{}, int, error) {
	var integerBytes []byte
	for _, char := range bencodedString {
		if len(integerBytes) == 0 && char == 'i' {
			continue
		} else if char != 'e' {
			// integerString += char
			integerBytes = append(integerBytes, byte(char))
		} else {
			break
		}
	}

	int_length := len(string(integerBytes)) + 2
	j, err := strconv.Atoi(string(integerBytes))
	if err != nil {
		return nil, int_length, err
	}
	return j, int_length, nil
}

func decodeList(bencodedString string) (interface{}, int, error) {

	bListInternal := bencodedString[1:]
	initLenght := len(bListInternal)

	var bList []interface{}

	for {
		next, l, err := DecodeBencode(bListInternal)

		if err != nil {
			panic(err)
		}
		bListInternal = bListInternal[l:]
		if bListInternal[0:1] == "" || bListInternal[0:1] == "e" {

			bListInternal = bListInternal[1:]
			bList = append(bList, next)
			break
		}

		bList = append(bList, next)

	}

	return bList, initLenght - len(bListInternal) + 1, nil
}


func decodeDict(bencodedString string) (interface{}, int, error) {
	
	bListInternal := bencodedString[1:]
	initLenght := len(bListInternal)

	bDict := make(map[string]interface{})

	for {
		
		key, keyL, err := DecodeBencode(bListInternal)
		
		if err != nil {
			panic(err)
		}
		bListInternal = bListInternal[keyL:]
		
		value, valueL, err := DecodeBencode(bListInternal)
		if err != nil {

			panic(err)
		}
		bListInternal = bListInternal[valueL:]


		if bListInternal[0:1] == "" || bListInternal[0:1] == "e" {

			bListInternal = bListInternal[1:]
			bDict[key.(string)] = value
			break
		}
		
		bDict[key.(string)] = value
	}

	return bDict, initLenght - len(bListInternal) + 1, nil
}

func DecodeBencode(bencodedString string) (interface{}, int, error) {

	if unicode.IsDigit(rune(bencodedString[0])) {

		str, l, err := decodeString(bencodedString)
		return str, l, err

	} else if bencodedString[0] == 'i' {
		inter, l, err := decodeInteger(bencodedString)
		return inter, l, err

	} else if bencodedString[0] == 'l' {
		list, l, err := decodeList(bencodedString)
		return list, l, err

	} else if bencodedString[0] == 'd' {
		dict, l, err := decodeDict(bencodedString)
		return dict, l, err
	} else {
		return "", 0, fmt.Errorf("Only strings and intergers are supported at the moment")
	}
}
