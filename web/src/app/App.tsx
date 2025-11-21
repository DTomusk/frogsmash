import { AuthProvider } from "./providers/AuthContext";
import { SnackbarProvider } from "./providers/SnackbarContext";
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
