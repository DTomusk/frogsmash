import { AuthProvider } from "./providers/AuthContext";
import { SnackbarProvider } from "./providers/SnackbarContext";
import { useTenant } from "./providers/TenantProvider";
import { AppRoutes } from "./routes/AppRoutes";
import { tenantFavicons } from "./tenants/TenantFavicon";


function App() {
  const config = useTenant();
  
  const favicon = tenantFavicons[config.tenantKey];
  
  if (favicon) {
    const link = document.querySelector("link[rel~='icon']") as HTMLLinkElement;
  
    if (link) {
      link.href = favicon;
    } else {
      const newLink = document.createElement("link");
      newLink.rel = "icon";
      newLink.href = favicon;
      document.head.appendChild(newLink);
    }
  }

  document.title = config.appTitle;

  return (
  <AuthProvider>
    <SnackbarProvider>
      <AppRoutes />
    </SnackbarProvider>
  </AuthProvider>
  );
}

export default App;
