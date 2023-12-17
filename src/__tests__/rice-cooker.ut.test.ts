import {RiceCooker} from '../classes/RiceCooker';
import {messages} from '../lib';

describe('Test rice cooker class', () => {
  const rc = new RiceCooker();
  it('Power on not plugged ko', () => {
    expect(() => rc.togglePower()).toThrow(messages['error:power-not-plugged']);
  });

  it('Plug ok', () => {
    expect(() => rc.togglePlug()).not.toThrow();
  });

  it('Power on plugged ok', () => {
    expect(() => rc.togglePower()).not.toThrow();
  });

  it('Test add and remove water', () => {
    const waterToAdd = 10;
    const waterToRemove = 10;
    expect(() => rc.addWater(waterToAdd)).not.toThrow();
    expect(() => rc.addWater(waterToAdd)).toThrow(
      messages['error:max-capacity']
    );
    expect(() => rc.removeWater(waterToRemove)).not.toThrow();
    expect(() => rc.removeWater(waterToRemove)).toThrow(
      messages['error:not-enough-water']
    );
  });

  it('Test add and discard rice', () => {
    const riceToAdd = 10;
    const riceToDiscard = 10;
    expect(() => rc.addRice(riceToAdd)).not.toThrow();
    expect(() => rc.addRice(riceToAdd)).toThrow(messages['error:max-capacity']);
    expect(() => rc.discardRice(riceToDiscard)).not.toThrow();
    expect(() => rc.discardRice(riceToDiscard)).toThrow(
      messages['error:not-enough-rice']
    );
  });

  it('Test add and remove water and rice', () => {
    const waterToAdd = 10;
    const waterToRemove = 5;
    const riceToAdd = 5;
    const riceToDiscard = 10;

    expect(() => rc.addWater(waterToAdd)).not.toThrow();
    expect(() => rc.addRice(riceToAdd)).toThrow(messages['error:max-capacity']);

    expect(() => rc.removeWater(waterToRemove)).not.toThrow();
    expect(() => rc.addRice(riceToAdd)).not.toThrow();

    expect(() => rc.discardRice(riceToDiscard)).toThrow(
      messages['error:not-enough-rice']
    );

    const expectedWater = 5;
    const expectedRice = 5;
    expect(rc.getRiceCupNumber()).toBe(expectedWater);
    expect(rc.getWaterCupNumber()).toBe(expectedRice);
  });
});
