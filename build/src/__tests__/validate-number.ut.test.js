"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const validator_1 = require("../lib/validator");
describe('test number validator', () => {
    it('Validate number ok', () => {
        const validNumber = '10';
        expect(() => (0, validator_1.validateNumber)(validNumber)).not.toThrow();
    });
    it('Validate number ko', () => {
        const notValidNumber = 'not valid';
        expect(() => (0, validator_1.validateNumber)(notValidNumber)).toThrow();
    });
});
//# sourceMappingURL=validate-number.ut.test.js.map