import { Navigate, Outlet } from "react-router-dom";
import { useAuth } from "../providers/AuthContext";
import { useCurrentUser } from "@/features/auth/hooks/useCurrentUser";

function ProtectedRoute({ requireVerified = false }: { requireVerified?: boolean }) {
  const { token } = useAuth();
  const { data: user, isPending } = useCurrentUser();

  if (!token) {
    return <Navigate to="/login" replace />;
  }

  if (isPending) {
    return null;
  }

  if (requireVerified && !user?.isVerified) {
    return <Navigate to="/verify/required" replace />;
  }

  return <Outlet />;
}

export default ProtectedRoute;
