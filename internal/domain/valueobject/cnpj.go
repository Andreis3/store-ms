package valueobject

import (
	error2 "github.com/andreis3/stores-ms/internal/domain/notification"
	"regexp"
	"slices"
	"strconv"
)

var blackListCNPJ = []string{
	"00000000000000",
	"11111111111111",
	"22222222222222",
	"33333333333333",
	"44444444444444",
	"55555555555555",
	"66666666666666",
	"77777777777777",
	"88888888888888",
	"99999999999999",
}

type CNPJ struct {
	CNPJ string
}

func NewCNPJ(cnpj string) *CNPJ {
	return &CNPJ{CNPJ: cnpj}
}

func (c *CNPJ) Validate(ctx *error2.NotificationError) {
	regex := regexp.MustCompile("[^0-9]")
	cnpj := regex.ReplaceAllString(c.CNPJ, "")
	if cnpj == "" {
		ctx.AddNotification(`cnpj: is required`)
	}
	if cnpj != "" && len(cnpj) < 14 {
		ctx.AddNotification(`cnpj: is invalid, must have 14 characters`)
	}
	if cnpj != "" && slices.Contains(blackListCNPJ, cnpj) {
		ctx.AddNotification(`cnpj: is invalid, must be a valid CNPJ number`)
	}
	if cnpj != "" && !validateCNPJ(cnpj) && len(cnpj) == 14 {
		ctx.AddNotification(`cnpj: is invalid, must be a valid CNPJ number calculated with the module 11 algorithm`)
	}
}
func validateCNPJ(cnpj string) bool {
	size := len(cnpj) - 2
	numbers := cnpj[:size]
	digits := cnpj[size:]
	sum := 0
	pos := size - 7
	for i := size; i >= 1; i-- {
		convertNumber, _ := strconv.Atoi(string(numbers[size-i]))
		sum += convertNumber * pos
		pos--
		if pos < 2 {
			pos = 9
		}
	}
	result := 0
	if rest := sum % 11; rest < 2 {
		result = 0
	} else {
		result = 11 - (sum % 11)
	}
	if strconv.Itoa(result) != string(digits[0]) {
		return false
	}
	size++
	numbers = cnpj[:size]
	sum = 0
	pos = size - 7
	for i := size; i >= 1; i-- {
		convertNumber, _ := strconv.Atoi(string(numbers[size-i]))
		sum += convertNumber * pos
		pos--
		if pos < 2 {
			pos = 9
		}
	}
	if rest := sum % 11; rest < 2 {
		result = 0
	} else {
		result = 11 - (sum % 11)
	}
	if strconv.Itoa(result) != string(digits[1]) {
		return false
	}
	return true
}
