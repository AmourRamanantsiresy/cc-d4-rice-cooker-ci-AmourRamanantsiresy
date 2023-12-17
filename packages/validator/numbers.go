package validator

import (
	"errors"
	"strconv"

	"co.hei.ricecooker/packages/constants"
)

func ValidateNumber(str string) error {
	if _, err := strconv.ParseFloat(str, 64); err != nil {
		return errors.New(constants.Messages["error:bad-input-value"])
	}
	return nil
}
