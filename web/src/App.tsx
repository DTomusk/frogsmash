import { AuthProvider } from "./auth/contexts/AuthContext";
import { SnackbarProvider } from "./shared/contexts/SnackbarContext";
import { AppRoutes } from "./routes/AppRoutes";

function App() {
  return (
  <AuthProvider>
    <SnackbarProvider>
      <AppRoutes />
    </SnackbarProvider>
  </AuthProvider>
  );
}

export default App;
