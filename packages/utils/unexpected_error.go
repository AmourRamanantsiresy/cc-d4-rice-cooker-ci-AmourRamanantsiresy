package utils

import (
	"errors"
	"math/rand"

	"co.hei.ricecooker/packages/constants"
)

func GetRandom() int {
	return rand.Intn(51)
}

func shouldThrow() bool {
	return GetRandom() == GetRandom()
}

func WaterLeakage() error {
	if shouldThrow() {
		return errors.New(constants.Messages["error:water-leakage"])
	}
	return nil
}

func SuddenShutdown() error {
	if shouldThrow() {
		return errors.New(constants.Messages["error:no-electricity"])
	}
	return nil
}
