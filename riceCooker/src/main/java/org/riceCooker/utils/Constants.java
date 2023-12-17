package org.riceCooker.utils;

import java.util.HashMap;
import java.util.Map;


/**
 * Constants string for the cli label
 */
public class Constants {

    private static final Map<String, String> MESSAGES = new HashMap<>();
    private static final Map<String, String> UI = new HashMap<>();

    static {
        MESSAGES.put("error:max-capacity", "Error : Maximum capacity "
                + "exceeded.");
        MESSAGES.put("error:not-enough-water", "Error : You didn't put"
                + " in enough water.");
        MESSAGES.put("error:not-enough-rice", "Error : You didn't put"
                + " in enough rice.");
        MESSAGES.put("error:power-not-plugged", "Error : Cannot power on,"
                +
                " plug and retry.");
        MESSAGES.put("error:plug-explosion", "Error : There was a big error"
                + " so the rice cooker expose.");
        MESSAGES.put("error:bad-input-value", "Error : Bad input value.");
        MESSAGES.put("error:bad-choice", "Error : The specified choice do "
                + "not exist.");
        MESSAGES.put("error:no-electricity", "kill : The electricity was cut"
                + " off and the appliance stopped.");
        MESSAGES.put("error:water-leakage", "kill : There was a water leak,"
                + " and the appliance stopped.");
        MESSAGES.put("kill:quite", "Rice-cooker was killed.");
        UI.put("banner", "rice-cooker-logo.txt");
        UI.put("error", "error.txt");
        UI.put("menu", "menu.txt");
        UI.put("result", "result.txt");
        UI.put("simpleResult", "simple-result.txt");
        UI.put("title", "title.txt");
    }

    public static String getMessage(final String key) {
        return MESSAGES.get(key);
    }

    public static String getUi(final String key) {
        return UI.get(key);
    }
}
