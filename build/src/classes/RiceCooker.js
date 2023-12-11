"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.RiceCooker = void 0;
const lib_1 = require("../lib");
class RiceCooker {
    constructor() {
        this.maxCupNumber = 10;
        this.waterCupNumber = 0;
        this.riceCupNumber = 0;
        this.plug = false;
        this.power = false;
        this.getWaterCupNumber = () => this.waterCupNumber;
        this.getRiceCupNumber = () => this.riceCupNumber;
        this.getPlug = () => this.plug;
        this.getPower = () => this.power;
    }
    addWater(cups) {
        (0, lib_1.WaterLeakage)();
        if (this.waterCupNumber + this.riceCupNumber + cups > this.maxCupNumber) {
            throw new Error(lib_1.messages['error:max-capacity']);
        }
        this.waterCupNumber += cups;
    }
    addRice(cups) {
        if (this.waterCupNumber + this.riceCupNumber + cups > this.maxCupNumber) {
            throw new Error(lib_1.messages['error:max-capacity']);
        }
        this.riceCupNumber += cups;
    }
    removeWater(cups) {
        if (this.waterCupNumber < cups) {
            throw new Error(lib_1.messages['error:not-enough-water']);
        }
        this.waterCupNumber -= cups;
    }
    discardRice(cups) {
        if (this.riceCupNumber < cups) {
            throw new Error(lib_1.messages['error:not-enough-rice']);
        }
        this.riceCupNumber -= cups;
    }
    togglePlug() {
        const magicNumber = Math.round(Math.random() * 100);
        if (magicNumber > 48 && magicNumber < 52) {
            throw new Error(lib_1.messages['error:plug-explosion']);
        }
        this.plug = !this.power;
    }
    togglePower() {
        if (!this.plug) {
            throw new Error(lib_1.messages['error:power-not-plugged']);
        }
        this.power = !this.power;
    }
}
exports.RiceCooker = RiceCooker;
//# sourceMappingURL=RiceCooker.js.map