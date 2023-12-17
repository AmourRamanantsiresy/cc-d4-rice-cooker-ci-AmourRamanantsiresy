import {validateNumber} from '../lib/validator';

describe('test number validator', () => {
  it('Validate number ok', () => {
    const validNumber = '10';
    expect(() => validateNumber(validNumber)).not.toThrow();
  });

  it('Validate number ko', () => {
    const notValidNumber = 'not valid';
    expect(() => validateNumber(notValidNumber)).toThrow();
  });
});
