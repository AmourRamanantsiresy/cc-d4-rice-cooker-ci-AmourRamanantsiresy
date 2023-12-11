"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.validateNumber = void 0;
const messages_1 = require("../messages");
const validateNumber = (str) => {
    if (isNaN(+str)) {
        throw new Error(messages_1.messages['error:bad-input-value']);
    }
};
exports.validateNumber = validateNumber;
//# sourceMappingURL=number.js.map