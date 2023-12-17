package org.riceCooker.ut;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.riceCooker.classes.RiceCooker;

import static org.junit.jupiter.api.Assertions.*;
import static org.riceCooker.utils.Constants.getMessage;

class RiceCookerTest {
    private RiceCooker rc = new RiceCooker();

    @BeforeEach
    void newRiceCooker() {
        rc = new RiceCooker();
    }

    @Test
    void power_on_not_plugged_ko() {
        assertThrows(RuntimeException.class, () -> rc.togglePower());
        try {
            rc.togglePower();
        } catch (RuntimeException re) {
            assertEquals(re.getMessage(), getMessage("error:power-not-plugged"));
        }
    }

    @Test
    void plug_ok() {
        assertDoesNotThrow(() -> rc.togglePlug());
    }

    @Test
    void power_on_plugged_ok() {
        rc.togglePlug();
        assertTrue(rc.getPlug());
        assertDoesNotThrow(() -> rc.togglePower());
    }

    @Test
    void add_water_ok() {
        final int waterToAdd = 10;
        assertDoesNotThrow(() -> rc.addWater(waterToAdd));
    }

    @Test
    void add_water_ko() {
        final int waterToAdd = 11;
        assertThrows(RuntimeException.class, () -> rc.addWater(waterToAdd));
        try {
            rc.addWater(waterToAdd);
        } catch (RuntimeException re) {
            assertEquals(re.getMessage(), getMessage("error:max-capacity"));
        }
    }

    @Test
    void remove_water_ok() {
        final int waterToAdd = 6;
        final int waterToRemove = 5;
        final int expectedRest = 1;

        rc.addWater(waterToAdd);
        assertDoesNotThrow(() -> rc.removeWater(waterToRemove));
        assertEquals(expectedRest, rc.getWaterCupNumber());
    }


    @Test
    void remove_water_ko() {
        final int waterToRemove = 1;
        assertThrows(RuntimeException.class, () -> rc.removeWater(waterToRemove));
        try {
            rc.addWater(waterToRemove);
        } catch (RuntimeException re) {
            assertEquals(re.getMessage(), getMessage("error:not-enough-water"));
        }
    }
    @Test
    void add_rice_ok() {
        final int riceToAdd = 10;
        assertDoesNotThrow(() -> rc.addRice(riceToAdd));
    }

    @Test
    void add_rice_ko() {
        final int riceToAdd = 11;
        assertThrows(RuntimeException.class, () -> rc.addRice(riceToAdd));
        try {
            rc.addRice(riceToAdd);
        } catch (RuntimeException re) {
            assertEquals(re.getMessage(), getMessage("error:max-capacity"));
        }
    }

    @Test
    void remove_rice_ok() {
        final int riceToAdd = 6;
        final int riceToRemove = 5;
        final int expectedRest = 1;

        rc.addRice(riceToAdd);
        assertDoesNotThrow(() -> rc.discardRice(riceToRemove));
        assertEquals(expectedRest, rc.getRiceCupNumber());
    }


    @Test
    void remove_rice_ko() {
        final int riceToRemove = 1;
        assertThrows(RuntimeException.class, () -> rc.discardRice(riceToRemove));
        try {
            rc.discardRice(riceToRemove);
        } catch (RuntimeException re) {
            assertEquals(re.getMessage(), getMessage("error:not-enough-rice"));
        }
    }
}
