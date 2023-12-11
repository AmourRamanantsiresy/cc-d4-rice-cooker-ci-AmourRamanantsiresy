export declare class RiceCooker {
    private maxCupNumber;
    private waterCupNumber;
    private riceCupNumber;
    private plug;
    private power;
    constructor();
    addWater(cups: number): void;
    addRice(cups: number): void;
    removeWater(cups: number): void;
    discardRice(cups: number): void;
    togglePlug(): void;
    togglePower(): void;
    getWaterCupNumber: () => number;
    getRiceCupNumber: () => number;
    getPlug: () => boolean;
    getPower: () => boolean;
}
