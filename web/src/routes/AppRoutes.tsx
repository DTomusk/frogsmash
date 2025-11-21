import { Routes, Route } from "react-router-dom";
import Template from "../shared/components/templates/Template";
import LandingPage from "../shared/components/pages/LandingPage";
import RegistrationPage from "../auth/components/pages/RegistrationPage";
import LoginPage from "../auth/components/pages/LoginPage";
import VerificationPage from "../auth/components/pages/VerificationPage";
import ProtectedRoute from "./ProtectedRoute";
import ComparisonPage from "../comparison/components/pages/ComparisonPage";
import LeaderboardPage from "../comparison/components/pages/LeaderboardPage";
import VerificationRequiredPage from "../auth/components/pages/VerificationRequiredPage";
import UploadPage from "../comparison/components/pages/UploadPage";
import NotFoundPage from "../shared/components/pages/NotFoundPage";

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
