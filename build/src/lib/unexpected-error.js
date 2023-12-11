"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.suddenShutdown = exports.WaterLeakage = void 0;
const messages_1 = require("./messages");
const getRandom = () => Math.round(Math.random() * 50);
const shouldThrow = () => getRandom() === getRandom();
const WaterLeakage = () => {
    if (shouldThrow())
        throw new Error(messages_1.messages['error:water-leakage']);
};
exports.WaterLeakage = WaterLeakage;
const suddenShutdown = () => {
    if (shouldThrow())
        throw new Error(messages_1.messages['error:no-electricity']);
};
exports.suddenShutdown = suddenShutdown;
//# sourceMappingURL=unexpected-error.js.map