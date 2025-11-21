import { Navigate, Outlet } from "react-router-dom";
import { useAuth } from "../auth/contexts/AuthContext";

function ProtectedRoute({requireVerified = false}: {requireVerified?: boolean}) {
    const { token, user } = useAuth();

    if (!token) {
        return <Navigate to="/login" replace />;
    }

    // TODO: add verification page to resend verification email
    if (requireVerified && !user?.isVerified) {
        return <Navigate to="/verify/required" replace />;
    }

    return <Outlet />;
}

export default ProtectedRoute;