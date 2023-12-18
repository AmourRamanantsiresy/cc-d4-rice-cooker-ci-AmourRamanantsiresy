package tests

import (
	"testing"

	"co.hei.ricecooker/packages/constants"
	"co.hei.ricecooker/packages/validator"
)

func TestValidateNumberOk(t *testing.T) {
	valid_number := "100"
	err := validator.ValidateNumber(valid_number)

	if err != nil {
		t.Errorf("ValidateNumber(%s) shouldn't return error.", valid_number)
	}
}

func TestValidateNumberKo(t *testing.T) {
	not_valid_number := "1not00"
	err := validator.ValidateNumber(not_valid_number)
	expectedErrorMessage := constants.Messages["error:bad-input-value"]
	if err == nil {
		t.Errorf("ValidateNumber(%s) should return error.", not_valid_number)
	} else if err.Error() != expectedErrorMessage {
		t.Errorf("Bad error value, expected '%s' but got '%s'.", expectedErrorMessage, err.Error())
	}
}
