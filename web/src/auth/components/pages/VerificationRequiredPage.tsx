import { Typography } from "@mui/material";
import ContentWrapper from "../../../shared/components/atoms/ContentWrapper";
import ResendVerificationButton from "../organisms/ResendVerificationButton";
import { useAuth } from "../../contexts/AuthContext";
import { useEffect } from "react";
import { useNavigate } from "react-router-dom";

function VerificationRequiredPage() {
    const { user } = useAuth();
    const navigate = useNavigate(); 

    useEffect(() => {
        console.log("User in VerificationRequiredPage:", user);
        if (user && user.isVerified) {
            navigate("/smash");
        }
    }, [user, navigate]);

    return (
        <ContentWrapper>
            <Typography variant="h3" sx={{ mb: 2}}>Verification Required🐸</Typography>
            <Typography variant="subtitle1" sx={{mb: 2}}>To access this feature, please verify your account by clicking the link sent to your registered email address. If you haven't received the email, please check your spam folder or request a new verification email by clicking the magical button below.</Typography>
            <ResendVerificationButton />
        </ContentWrapper>
    );
}

export default VerificationRequiredPage;