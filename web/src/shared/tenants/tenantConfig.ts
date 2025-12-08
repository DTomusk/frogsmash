import { bookConfig } from "./book";
import { frogConfig } from "./frog";

export interface TenantConfig {
  title: string;
}

export const tenantConfigMap: Record<string, TenantConfig> = {
    frog: frogConfig,
    book: bookConfig
};