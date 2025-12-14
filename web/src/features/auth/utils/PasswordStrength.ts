export const passwordRules = {
  minLength: (pw: string) => pw.length >= 12,
  lowerCase: (pw: string) => /[a-z]/.test(pw),
  upperCase: (pw: string) => /[A-Z]/.test(pw),
  number: (pw: string) => /\d/.test(pw),
  specialChar: (pw: string) => /[!@#$%^&*(),.?":{}|<>]/.test(pw),
};

export type PasswordStrengthResult = {
  minLength: boolean;
  lowerCase: boolean;
  upperCase: boolean;
  number: boolean;
  specialChar: boolean;
};

export const checkPasswordStrength = (pw: string): PasswordStrengthResult => ({
  minLength: passwordRules.minLength(pw),
  lowerCase: passwordRules.lowerCase(pw),
  upperCase: passwordRules.upperCase(pw),
  number: passwordRules.number(pw),
  specialChar: passwordRules.specialChar(pw),
});
