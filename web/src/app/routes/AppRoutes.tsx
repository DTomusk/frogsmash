import { Routes, Route } from "react-router-dom";
import Template from "../../shared/components/templates/Template";
import LandingPage from "../../shared/components/pages/LandingPage";
import RegistrationPage from "../../features/auth/components/pages/RegistrationPage";
import LoginPage from "../../features/auth/components/pages/LoginPage";
import VerificationPage from "../../features/auth/components/pages/VerificationPage";
import ProtectedRoute from "./ProtectedRoute";
import ComparisonPage from "../../features/comparison/components/pages/ComparisonPage";
import LeaderboardPage from "../../features/comparison/components/pages/LeaderboardPage";
import VerificationRequiredPage from "../../features/auth/components/pages/VerificationRequiredPage";
import UploadPage from "../../features/comparison/components/pages/UploadPage";
import NotFoundPage from "../../shared/components/pages/NotFoundPage";

export function AppRoutes() {
  return (
    <Routes>
      <Route element={<Template />}>
        <Route path="/" element={<LandingPage />} />
        <Route path="/register" element={<RegistrationPage />} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/verify" element={<VerificationPage />} />

        <Route element={<ProtectedRoute />}>
          <Route path="/smash" element={<ComparisonPage />} />
          <Route path="/leaderboard" element={<LeaderboardPage />} />
          <Route path="/verify/required" element={<VerificationRequiredPage />} />
        </Route>

        <Route element={<ProtectedRoute requireVerified />}>
          <Route path="/upload" element={<UploadPage />} />
        </Route>

        <Route path="*" element={<NotFoundPage />} />
      </Route>
    </Routes>
  );
}
