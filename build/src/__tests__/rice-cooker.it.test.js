"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const RiceCooker_1 = require("../classes/RiceCooker");
const lib_1 = require("../lib");
const rc = new RiceCooker_1.RiceCooker();
test('Power on not plugged ko', () => {
    expect(() => rc.togglePower()).toThrow(lib_1.messages['error:power-not-plugged']);
});
test('Plug ok', () => {
    expect(() => rc.togglePlug()).not.toThrow();
});
test('Power on plugged ok', () => {
    expect(() => rc.togglePower()).not.toThrow();
});
test('Test add and remove water', () => {
    const waterToAdd = 10;
    const waterToRemove = 10;
    expect(() => rc.addWater(waterToAdd)).not.toThrow();
    expect(() => rc.addWater(waterToAdd)).toThrow(lib_1.messages['error:max-capacity']);
    expect(() => rc.removeWater(waterToRemove)).not.toThrow();
    expect(() => rc.removeWater(waterToRemove)).toThrow(lib_1.messages['error:not-enough-water']);
});
test('Test add and discard rice', () => {
    const riceToAdd = 10;
    const riceToDiscard = 10;
    expect(() => rc.addRice(riceToAdd)).not.toThrow();
    expect(() => rc.addRice(riceToAdd)).toThrow(lib_1.messages['error:max-capacity']);
    expect(() => rc.discardRice(riceToDiscard)).not.toThrow();
    expect(() => rc.discardRice(riceToDiscard)).toThrow(lib_1.messages['error:not-enough-rice']);
});
test('Test add and remove water and rice', () => {
    const waterToAdd = 10;
    const waterToRemove = 5;
    const riceToAdd = 5;
    const riceToDiscard = 10;
    expect(() => rc.addWater(waterToAdd)).not.toThrow();
    expect(() => rc.addRice(riceToAdd)).toThrow(lib_1.messages['error:max-capacity']);
    expect(() => rc.removeWater(waterToRemove)).not.toThrow();
    expect(() => rc.addRice(riceToAdd)).not.toThrow();
    expect(() => rc.discardRice(riceToDiscard)).toThrow(lib_1.messages['error:not-enough-rice']);
    const expectedWater = 5;
    const expectedRice = 5;
    expect(rc.getRiceCupNumber()).toBe(expectedWater);
    expect(rc.getWaterCupNumber()).toBe(expectedRice);
});
//# sourceMappingURL=rice-cooker.it.test.js.map