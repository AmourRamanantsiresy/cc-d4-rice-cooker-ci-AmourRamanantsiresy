package utils

import (
	"co.hei.ricecooker/packages/constants"
)

func ShowBanner() {
	clearConsole()
	Show(constants.UiTemplates["banner"], nil, 0)
}
