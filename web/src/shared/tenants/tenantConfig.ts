import { bookConfig } from "./book";
import { frogConfig } from "./frog";

export interface TenantConfig {
  tenantKey: string;
  title: string;
  titleImageText1: string;
  titleImageText2: string;
  titleImageText3: string;
}

export const tenantConfigMap: Record<string, TenantConfig> = {
    frog: frogConfig,
    book: bookConfig
};