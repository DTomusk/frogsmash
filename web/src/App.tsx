import { Route, Routes } from "react-router-dom";
import Comparison from "./comparison/components/pages/ComparisonPage";
import Template from "./shared/components/templates/Template";
import LoadingPage from "./shared/components/pages/LoadingPage";
import LeaderboardPage from "./comparison/components/pages/LeaderboardPage";
import NotFoundPage from "./shared/components/pages/NotFoundPage";
import UploadPage from "./comparison/components/pages/UploadPage";
import RegistrationPage from "./auth/components/pages/RegistrationPage";
import LoginPage from "./auth/components/pages/LoginPage";
import { AuthProvider } from "./auth/contexts/AuthContext";
import ProtectedRoute from "./auth/components/atoms/ProtectedRoute";
import { SnackbarProvider } from "./shared/contexts/SnackbarContext";
import LandingPage from "./shared/components/pages/LandingPage";
import VerificationPage from "./auth/components/pages/VerificationPage";
import VerificationRequiredPage from "./auth/components/pages/VerificationRequiredPage";

function App() {
  return (
  <AuthProvider>
    <SnackbarProvider>
      <Routes>
        <Route element={<Template />}>
            <Route path='/' element={<LandingPage />} />
            <Route path='/register' element={<RegistrationPage />} />
            <Route path='/login' element={<LoginPage />} />
            <Route path='/verify' element={<VerificationPage />} />
            <Route element={<ProtectedRoute />}>
              <Route path='/smash' element={<Comparison />} />
              <Route path='/loading' element={<LoadingPage />} />
              <Route path='/leaderboard' element={<LeaderboardPage />} />
              <Route path='/verify/required' element={<VerificationRequiredPage />} />
            </Route>
            <Route element={<ProtectedRoute requireVerified={true} />}>
              <Route path='/upload' element={<UploadPage />} />
            </Route>
            <Route path='*' element={<NotFoundPage />} />
        </Route>
      </Routes>
    </SnackbarProvider>
  </AuthProvider>
  );
}

export default App;
