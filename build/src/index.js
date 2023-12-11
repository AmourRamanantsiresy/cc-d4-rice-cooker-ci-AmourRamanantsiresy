"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const RiceCooker_1 = require("./classes/RiceCooker");
const lib_1 = require("./lib");
const validator_1 = require("./lib/validator");
const plugHandler = (rc) => {
    const title = !rc.getPlug() ? 'Plug' : 'Unplug';
    (0, lib_1.show)(lib_1.templates.title, [title]);
    rc.togglePlug();
    (0, lib_1.show)(lib_1.templates.simple_result, ['Done'], title.length);
};
const powerHandler = (rc) => {
    const title = !rc.getPower() ? 'Power on' : 'Power off';
    (0, lib_1.show)(lib_1.templates.title, [title]);
    rc.togglePower();
    (0, lib_1.show)(lib_1.templates.simple_result, ['Done !'], title.length);
};
const addWaterHandler = async (rc) => {
    const title = 'Add water cups.';
    (0, lib_1.show)(lib_1.templates.title, [title]);
    const cups = await (0, lib_1.input)('Number of cups of water you want to add : ');
    (0, validator_1.validateNumber)(cups);
    rc.addWater(+cups);
    (0, lib_1.show)(lib_1.templates.simple_result, ['Done !'], title.length);
};
const addRiceHandler = async (rc) => {
    const title = 'Add Rice cups.';
    (0, lib_1.show)(lib_1.templates.title, [title]);
    const cups = await (0, lib_1.input)('Number of cups of rice you want to add : ');
    (0, validator_1.validateNumber)(cups);
    rc.addRice(+cups);
    (0, lib_1.show)(lib_1.templates.simple_result, ['Done !'], title.length);
};
const removeWaterHandler = async (rc) => {
    const title = 'Remove water cups.';
    (0, lib_1.show)(lib_1.templates.title, [title]);
    const cups = await (0, lib_1.input)('Number of cups of water you want to remove : ');
    (0, validator_1.validateNumber)(cups);
    rc.removeWater(+cups);
    (0, lib_1.show)(lib_1.templates.simple_result, ['Done !'], title.length);
};
const discardRiceHandler = async (rc) => {
    const title = 'Discard Rice cups.';
    (0, lib_1.show)(lib_1.templates.title, [title]);
    const cups = await (0, lib_1.input)('Number of cups of rice you want to discard : ');
    (0, validator_1.validateNumber)(cups);
    rc.discardRice(+cups);
    (0, lib_1.show)(lib_1.templates.simple_result, ['Done !'], title.length);
};
const handler = async (rc) => {
    (0, lib_1.suddenShutdown)();
    (0, lib_1.show)(lib_1.templates.menu, [
        rc.getWaterCupNumber().toString(),
        rc.getRiceCupNumber().toString(),
        rc.getPlug() ? 'Plugged' : 'Unplugged',
        rc.getPower() ? 'On' : 'Off',
    ]);
    const choice = await (0, lib_1.input)('=> ');
    if (choice === '1')
        plugHandler(rc);
    else if (choice === '2')
        powerHandler(rc);
    else if (choice === '3')
        await addWaterHandler(rc);
    else if (choice === '4')
        await addRiceHandler(rc);
    else if (choice === '5')
        await removeWaterHandler(rc);
    else if (choice === '6')
        await discardRiceHandler(rc);
    else if (choice === '7')
        throw Error(lib_1.messages['kill:quite']);
    else
        throw new Error(lib_1.messages['error:bad-choice']);
};
const main = async () => {
    let run = true;
    const riceCooker = new RiceCooker_1.RiceCooker();
    while (run) {
        (0, lib_1.banner)();
        try {
            await handler(riceCooker);
        }
        catch (err) {
            const error = err;
            if (error.message.includes('kill')) {
                run = false;
            }
            else {
                (0, lib_1.show)(lib_1.templates.error, [error.message]);
            }
        }
        await (0, lib_1.input)('Type enter to continue.');
    }
};
main();
//# sourceMappingURL=index.js.map