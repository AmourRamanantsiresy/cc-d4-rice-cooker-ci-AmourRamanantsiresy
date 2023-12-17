package constants

var Messages = map[string]string{
	"error:max-capacity":      "Error : Maximum capacity exceeded.",
	"error:not-enough-water":  "Error : You didn't put in enough water.",
	"error:not-enough-rice":   "Error : You didn't put in enough rice.",
	"error:power-not-plugged": "Error : Cannot power on, plug and retry.",
	"error:plug-explosion":    "Error : Cannot power on, plug and retry.",
	"error:bad-input-value":   "Error : Bad input value.",
	"error:bad-choice":        "Error : The specified choice does not exist.",
	"error:no-electricity":    "kill : The electricity was cut off and the appliance stopped.",
	"error:water-leakage":     "kill : There was a water leak, and the appliance stopped.",
	"kill:quite":              "kill",
}
