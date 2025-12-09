import { frogThemes } from "./frogTheme";
import { bookThemes } from "./bookTheme"; // optional

export const tenantThemeMap: Record<
  string,
  { light: any; dark: any }
> = {
  frog: frogThemes,
  book: bookThemes,
};
