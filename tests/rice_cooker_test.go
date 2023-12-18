package tests

import (
	"testing"

	"co.hei.ricecooker/internal/classes"
	"co.hei.ricecooker/packages/constants"
)

var rc = classes.NewRiceCooker()

func TestPowerOnNotPluggedKo(t *testing.T) {
	err := rc.TogglePower()
	expectedErrorMessage := constants.Messages["error:power-not-plugged"]
	if err == nil {
		t.Error("Should return error 'cause rice cooker is not already plugged.")
	} else if err.Error() != expectedErrorMessage {
		t.Errorf("Expected error %s but got %s", expectedErrorMessage, err.Error())
	}
}

func TestPlugOk(t *testing.T) {
	err := rc.TogglePlug()
	if err != nil {
		t.Errorf("Shouldn't return error but got %s", err.Error())
	}

	isPlug := rc.GetPlug()
	if !isPlug {
		t.Error("Plug variable should be true but got false")
	}
}

func TestPowerOnPluggedOk(t *testing.T) {
	err := rc.TogglePower()
	if err != nil {
		t.Errorf("Shouldn't return error but got %s", err.Error())
	}
}

func TestAddWaterOk(t *testing.T) {
	validWaterCup := 5
	err := rc.AddWater(validWaterCup)
	if err != nil {
		t.Errorf("Shouldn't return error but got %s", err.Error())
	}
	currentWaterCup := rc.GetWaterCupNumber()
	if validWaterCup != currentWaterCup {
		t.Errorf("rc.GetWaterCupNumber() return %d, expected %d", currentWaterCup, validWaterCup)
	}
}

func TestAddWaterKo(t *testing.T) {
	notValidaWaterCup := 15
	err := rc.AddWater(notValidaWaterCup)
	expectedError := constants.Messages["error:max-capacity"]

	if err == nil {
		t.Errorf("Should return error")
	} else if expectedError != err.Error() {
		t.Errorf("Expected error %s but got %s", expectedError, err.Error())
	}
}

func TestAddRiceOk(t *testing.T) {
	validRiceCup := 5
	err := rc.AddRice(validRiceCup)
	if err != nil {
		t.Errorf("Shouldn't return error but got %s", err.Error())
	}
	currentRiceCup := rc.GetRiceCupNumber()
	if validRiceCup != currentRiceCup {
		t.Errorf("rc.GetRiceCupNumber() return %d, expected %d", currentRiceCup, validRiceCup)
	}
}

func TestAddRiceKo(t *testing.T) {
	notValidaRiceCup := 15
	err := rc.AddRice(notValidaRiceCup)
	expectedError := constants.Messages["error:max-capacity"]

	if err == nil {
		t.Errorf("Should return error")
	} else if expectedError != err.Error() {
		t.Errorf("Expected error %s but got %s", expectedError, err.Error())
	}
}

func TestRemoveWaterOk(t *testing.T) {
	validWaterCup := 1
	err := rc.RemoveWater(validWaterCup)
	if err != nil {
		t.Errorf("Shouldn't return error but got %s", err.Error())
	}
}
func TestRemoveWaterKo(t *testing.T) {
	notValidWaterCup := 10
	err := rc.RemoveWater(notValidWaterCup)
	expectedError := constants.Messages["error:not-enough-water"]

	if err == nil {
		t.Errorf("Should return error")
	} else if expectedError != err.Error() {
		t.Errorf("Expected error %s but got %s", expectedError, err.Error())
	}
}

func TestDiscardRiceOk(t *testing.T) {
	validRiceCup := 1
	err := rc.DiscardRice(validRiceCup)
	if err != nil {
		t.Errorf("Shouldn't return error but got %s", err.Error())
	}
}
func TestDiscardRiceKo(t *testing.T) {
	notValidRiceCup := 10
	err := rc.DiscardRice(notValidRiceCup)
	expectedError := constants.Messages["error:not-enough-rice"]

	if err == nil {
		t.Errorf("Should return error")
	} else if expectedError != err.Error() {
		t.Errorf("Expected error %s but got %s", expectedError, err.Error())
	}
}
