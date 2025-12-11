import { Routes, Route } from "react-router-dom";
import ProtectedRoute from "./ProtectedRoute";
import { LoginPage, RegistrationPage } from "@/features/auth";
import { ComparisonPage, LeaderboardPage, UploadPage } from "@/features/comparison";
import { VerificationPage, VerificationRequiredPage } from "@/features/verification";
import { LandingPage, LoadingPage, NotFoundPage, Template } from "@/shared";
import { useTenant } from "../providers/TenantProvider";
import PrivacyPolicyPage from "@/shared/components/pages/PrivacyPolicyPage";

export function AppRoutes() {
  const config = useTenant();
  return (
    <Routes>
      <Route element={<Template />}>
        <Route path="/" element={<LandingPage />} />
        <Route path="/register" element={<RegistrationPage />} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/verify" element={<VerificationPage />} />
        <Route path="/loading" element={<LoadingPage />} />
        <Route path="/privacy" element={<PrivacyPolicyPage />} />

        <Route element={<ProtectedRoute />}>
          <Route path="/smash" element={<ComparisonPage />} />
          <Route path="/leaderboard" element={<LeaderboardPage />} />
          <Route path="/verify/required" element={<VerificationRequiredPage />} />
        </Route>

        <Route element={<ProtectedRoute requireVerified />}>
          {config.tenantKey === "frog" && <Route path="/upload" element={<UploadPage />} />}
        </Route>

        <Route path="*" element={<NotFoundPage />} />
      </Route>
    </Routes>
  );
}
