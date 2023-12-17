package main

import (
	"errors"
	"strconv"

	"co.hei.ricecooker/internal/classes"
	"co.hei.ricecooker/packages/constants"
	"co.hei.ricecooker/packages/utils"
	"co.hei.ricecooker/packages/validator"
)

var rc = classes.NewRiceCooker()

func getTitleByBoolean(condition bool, ifTrue string, ifFalse string) string {
	title := ifFalse
	if condition {
		title = ifTrue
	}
	return title
}

func getPlugTitleByBoolean(condition bool) string {
	return getTitleByBoolean(condition, "UnPlug", "Plug")
}
func getPowerTitleByBoolean(condition bool) string {
	return getTitleByBoolean(condition, "Power off", "Power on")
}

func TogglePlug() {
	title := getPlugTitleByBoolean(rc.GetPlug())
	utils.Show(constants.UiTemplates["title"], []string{title}, 0)
	rc.TogglePlug()
	utils.Show(constants.UiTemplates["simple_result"], []string{"Done !"}, len(title))
}

func TogglePower() error {
	title := getPowerTitleByBoolean(rc.GetPower())
	utils.Show(constants.UiTemplates["title"], []string{title}, 0)
	err := rc.TogglePower()
	if err != nil {
		return err
	}
	utils.Show(constants.UiTemplates["simple_result"], []string{"Done !"}, len(title))
	return nil
}

func AddWater() error {
	title := "Add water cups."
	utils.Show(constants.UiTemplates["title"], []string{title}, 0)
	cups := utils.Input("Number of cups of water you want to add : ")
	err := validator.ValidateNumber(cups)
	if err != nil {
		return err
	}
	cupsNumber := utils.GetNumberFromString(cups)
	err = rc.AddWater(cupsNumber)
	if err != nil {
		return err
	}
	utils.Show(constants.UiTemplates["simple_result"], []string{"Done !"}, len(title))
	return nil
}

func AddRice() error {
	tilte := "Add Rice cups."
	utils.Show(constants.UiTemplates["title"], []string{tilte}, 0)
	cups := utils.Input("Number of cups of rice you want to add : ")
	err := validator.ValidateNumber(cups)
	if err != nil {
		return err
	}
	cupsNumber := utils.GetNumberFromString(cups)
	err = rc.AddRice(cupsNumber)
	if err != nil {
		return err
	}
	utils.Show(constants.UiTemplates["simple_result"], []string{"Done !"}, len(tilte))
	return nil
}

func RemoveWater() error {
	title := "Remove water cups."
	utils.Show(constants.UiTemplates["title"], []string{title}, 0)
	cups := utils.Input("Number of cups of rice you want to remove : ")
	err := validator.ValidateNumber(cups)
	if err != nil {
		return err
	}
	cupsNumber := utils.GetNumberFromString(cups)
	err = rc.RemoveWater(cupsNumber)
	if err != nil {
		return err
	}
	utils.Show(constants.UiTemplates["simple_result"], []string{"Done !"}, len(title))
	return nil
}

func DiscardRice() error {
	title := "Discard rice cups."
	utils.Show(constants.UiTemplates["title"], []string{title}, 0)
	cups := utils.Input("Number of cups of rice you want to discard : ")
	err := validator.ValidateNumber(cups)
	if err != nil {
		return err
	}
	cupsNumber := utils.GetNumberFromString(cups)
	err = rc.DiscardRice(cupsNumber)
	if err != nil {
		return err
	}
	utils.Show(constants.UiTemplates["simple_result"], []string{"Done !"}, len(title))
	return nil
}

func main() {
	for {
		utils.ShowBanner()
		infos := []string{
			strconv.Itoa(rc.GetWaterCupNumber()),
			strconv.Itoa(rc.GetRiceCupNumber()),
			getPlugTitleByBoolean(!rc.GetPlug()),
			getPowerTitleByBoolean(!rc.GetPower())}

		utils.Show(constants.UiTemplates["menu"], infos, 0)
		choice := utils.Input("=> ")
		var err error = nil

		switch choice {
		case "1":
			TogglePlug()
		case "2":
			err = TogglePower()
		case "3":
			err = AddWater()
		case "4":
			err = AddRice()
		case "5":
			err = RemoveWater()
		case "6":
			err = DiscardRice()
		case "7":
			return
		default:
			err = errors.New(constants.Messages["kill:quite"])
		}
		if err != nil {
			utils.Show(constants.UiTemplates["error"], []string{err.Error()}, 0)
		}
		utils.Input("Type enter to continue")
	}
}
