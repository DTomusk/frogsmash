import { tenantConfigMap } from "@/shared/tenants";
import type { TenantConfig } from "@/shared/tenants/tenantConfig";
import { createContext, useContext, useMemo, type ReactNode } from "react";

const TenantContext = createContext<TenantConfig | null>(null);

export const useTenant = () => {
  const ctx = useContext(TenantContext);
  if (!ctx) {
    throw new Error("useTenant must be used within <TenantProvider>");
  }
  return ctx;
};

export function TenantProvider({ children }: { children: ReactNode }) {
    const tenant = import.meta.env.VITE_TENANT ?? "frog";

    const config = useMemo<TenantConfig>(() => {
        return tenantConfigMap[tenant] ?? tenantConfigMap["frog"];
    }, [tenant]);

    return (
        <TenantContext.Provider value={config}>
            {children}
        </TenantContext.Provider>
    );
}