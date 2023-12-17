package classes

import (
	"errors"

	"co.hei.ricecooker/packages/constants"
	"co.hei.ricecooker/packages/utils"
)

type RiceCooker struct {
	maxCupNumber   int
	waterCupNumber int
	riceCupNumber  int
	plug           bool
	power          bool
}

func NewRiceCooker() *RiceCooker {
	return &RiceCooker{
		maxCupNumber:   10,
		waterCupNumber: 0,
		riceCupNumber:  0,
		plug:           false,
		power:          false,
	}
}

func (rc *RiceCooker) AddWater(cups int) error {
	if err := utils.WaterLeakage(); err != nil {
		return err
	}
	if rc.waterCupNumber+rc.riceCupNumber+cups > rc.maxCupNumber {
		return errors.New(constants.Messages["error:max-capacity"])
	}
	rc.waterCupNumber += cups
	return nil
}

func (rc *RiceCooker) AddRice(cups int) error {
	if rc.waterCupNumber+rc.riceCupNumber+cups > rc.maxCupNumber {
		return errors.New(constants.Messages["error:max-capacity"])
	}
	rc.riceCupNumber += cups
	return nil
}

func (rc *RiceCooker) RemoveWater(cups int) error {
	if rc.waterCupNumber < cups {
		return errors.New(constants.Messages["error:not-enough-water"])
	}
	rc.waterCupNumber -= cups
	return nil
}

func (rc *RiceCooker) DiscardRice(cups int) error {
	if rc.riceCupNumber < cups {
		return errors.New(constants.Messages["error:not-enough-rice"])
	}
	rc.riceCupNumber -= cups
	return nil
}

func (rc *RiceCooker) TogglePlug() error {
	magicNumber := utils.GetRandom()
	if magicNumber > 48 && magicNumber < 52 {
		return errors.New(constants.Messages["error:plug-explosion"])
	}
	rc.plug = !rc.plug
	return nil
}

func (rc *RiceCooker) TogglePower() error {
	if !rc.plug {
		return errors.New(constants.Messages["error:power-not-plugged"])
	}
	rc.power = !rc.power
	return nil
}

func (rc *RiceCooker) GetWaterCupNumber() int {
	return rc.waterCupNumber
}

func (rc *RiceCooker) GetRiceCupNumber() int {
	return rc.riceCupNumber
}

func (rc *RiceCooker) GetPlug() bool {
	return rc.plug
}

func (rc *RiceCooker) GetPower() bool {
	return rc.power
}
