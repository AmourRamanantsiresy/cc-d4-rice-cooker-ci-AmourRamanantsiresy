package org.riceCooker.ut;

import org.junit.jupiter.api.Test;
import org.riceCooker.utils.Validator;

import static org.junit.jupiter.api.Assertions.*;
import static org.riceCooker.utils.Constants.getMessage;

class ValidatorTest {

    @Test
    void validate_number_ok() {
        String valid_number = "10";
        assertDoesNotThrow(() -> Validator.validateNumber(valid_number));
    }


    @Test
    void validate_number_ko() {
        String not_valid_number = "1ab0";
        assertThrows(RuntimeException.class, () -> Validator.validateNumber(not_valid_number));
        try {
            Validator.validateNumber(not_valid_number);
        } catch (RuntimeException re) {
            assertEquals(re.getMessage(), getMessage("error:bad-input-value"));
        }
    }
}