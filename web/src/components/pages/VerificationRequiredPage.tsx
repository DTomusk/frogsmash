import { Typography } from "@mui/material";
import ContentWrapper from "../atoms/ContentWrapper";
import ResendVerificationButton from "../organisms/ResendVerificationButton";

function VerificationRequiredPage() {
    return (
        <ContentWrapper>
            <Typography variant="h3" sx={{ mb: 2}}>Verification Required🐸</Typography>
            <Typography variant="subtitle1" sx={{mb: 2}}>To access this feature, please verify your account by clicking the link sent to your registered email address. If you haven't received the email, please check your spam folder or request a new verification email by clicking the magical button below.</Typography>
            <ResendVerificationButton />
        </ContentWrapper>
    );
}

export default VerificationRequiredPage;