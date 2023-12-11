"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.banner = void 0;
const show_1 = require("./show");
const templates_1 = require("./templates");
/* eslint-disable node/no-unsupported-features/node-builtins */
const banner = () => {
    console.clear();
    (0, show_1.show)(templates_1.templates.banner);
};
exports.banner = banner;
//# sourceMappingURL=banner.js.map