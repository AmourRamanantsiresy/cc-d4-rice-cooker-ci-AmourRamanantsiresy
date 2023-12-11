"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const validator_1 = require("../lib/validator");
test('Validate number ok', () => {
    const validNumber = '10';
    expect(() => (0, validator_1.validateNumber)(validNumber)).not.toThrow();
});
test('Validate number ko', () => {
    const notValidNumber = 'not valid';
    expect(() => (0, validator_1.validateNumber)(notValidNumber)).toThrow();
});
//# sourceMappingURL=validate-number.it.test.js.map