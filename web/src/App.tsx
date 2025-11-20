import { Route, Routes } from "react-router-dom";
import Comparison from "./components/pages/ComparisonPage";
import Template from "./components/templates/Template";
import LoadingPage from "./components/pages/LoadingPage";
import LeaderboardPage from "./components/pages/LeaderboardPage";
import NotFoundPage from "./components/pages/NotFoundPage";
import UploadPage from "./components/pages/UploadPage";
import RegistrationPage from "./components/pages/RegistrationPage";
import LoginPage from "./components/pages/LoginPage";
import { AuthProvider } from "./contexts/AuthContext";
import ProtectedRoute from "./components/atoms/ProtectedRoute";
import { SnackbarProvider } from "./contexts/SnackbarContext";
import LandingPage from "./components/pages/LandingPage";
import VerificationPage from "./components/pages/VerificationPage";
import VerificationRequiredPage from "./components/pages/VerificationRequiredPage";

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
