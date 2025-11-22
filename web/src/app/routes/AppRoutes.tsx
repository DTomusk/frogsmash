import { Routes, Route } from "react-router-dom";
import ProtectedRoute from "./ProtectedRoute";
import { LoginPage, RegistrationPage } from "@/features/auth";
import { ComparisonPage, LeaderboardPage, UploadPage } from "@/features/comparison";
import { VerificationPage, VerificationRequiredPage } from "@/features/verification";
import { LandingPage, LoadingPage, NotFoundPage, Template } from "@/shared";

export function AppRoutes() {
  return (
    <Routes>
      <Route element={<Template />}>
        <Route path="/" element={<LandingPage />} />
        <Route path="/register" element={<RegistrationPage />} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/verify" element={<VerificationPage />} />
        <Route path="/loading" element={<LoadingPage />} />

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
