"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.show = void 0;
const fs = require("fs");
const show = (path, toReplace, toReplaceLength) => {
    const filePath = `./assets/${path}`;
    let fileContent = fs.readFileSync(filePath, 'utf-8');
    const replacePattern = 'xxx';
    const repeatPattern = '═';
    if (toReplace) {
        toReplace.forEach(value => {
            fileContent = fileContent.replace(replacePattern, value);
        });
        fileContent = fileContent.replace(replacePattern, repeatPattern.repeat(toReplaceLength || toReplace[0].length));
    }
    console.log(fileContent);
};
exports.show = show;
//# sourceMappingURL=show.js.map