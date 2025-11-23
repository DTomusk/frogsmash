import { Button, Typography } from "@mui/material";
import { ContentWrapper, LoadingSpinner } from "@/shared";
import { useNavigate, useSearchParams } from "react-router-dom";
import { useEffect, useState } from "react";
import ResendVerificationButton from "../organisms/ResendVerificationButton";
import { useVerifyCode } from "../../hooks/useVerify";
import { AlreadyVerifiedCode, VerifiedCode } from "../../dtos/verificationResponse";

// This page should just show components based on the user state and the presence of the code param
// 1. if code, verify
// 1.a. success
// 1.a.i user not logged in -> show success and prompt to log in
// 1.a.ii user logged in to same account as code -> show success and prompt to continue to app
// 1.a.iii user logged in to different account -> show success and prompt to log out and log in with correct account
// 1.b. error -> show error and prompt to resend verification email
// 1.b.i user not logged in -> show error and prompt to resend verification email
// 1.b.ii user logged in and not verified -> show error and prompt to resend verification email
// 1.b.iii user logged in and verified -> show already verified message and prompt to continue to app
// 2. if no code
// 2.a user not logged in -> show email entry to resend verification email
// 2.b user logged in and not verified -> show button to resend verification email
// 2.c user logged in and verified -> show already verified message and prompt to continue to app
function VerificationPage() {
    const [params] = useSearchParams();
    const code = params.get("code");
    const navigate = useNavigate();

    const { mutate: verifyCode, isPending: isVerifying } = useVerifyCode();

    const [status, setStatus] = useState<"pending" | "success" | "error">(code ? "pending" : "error");

    useEffect(() => {
        if (!code) {
            return;
        }

        console.log("Verifying code:", code);
        
        verifyCode(code, {
            onSuccess: (response) => {
                console.log("Verification response code:", response.code);
                if (response.code === VerifiedCode || response.code === AlreadyVerifiedCode) {
                    
                }
                setStatus("success");

            },
            onError: () => {
                setStatus("error");
            }
        });
    }, [code]);

    return (<>
        {status === "pending" && (<><ContentWrapper>
            <Typography variant="h3" sx={{ mb: 2}}>Verify your account🐸</Typography>
            <Typography variant="subtitle1" sx={{mb: 2}}>We've sent a verification email to your registered email address. It might take a few minutes to show up. Please check your inbox and click on the verification link to activate your account. If you still haven't received an email, please check your spam folder or request a new verification email by clicking the magical button below.</Typography>
            {isVerifying ? <LoadingSpinner /> : <ResendVerificationButton />}
        </ContentWrapper></>)}

        {status === "success" && (<><ContentWrapper>
            <Typography variant="h3" sx={{ mb: 2}}>Verification Successful🐸</Typography>
            <Typography variant="subtitle1" sx={{mb: 2}}>Thank you for verifying your account! You can now access all the features of our platform. Welcome aboard!</Typography>
            <Button variant="contained" 
                    size="large" 
                    color="primary"
                    sx={{ mt: 4, px: 2 }}
                    onClick={() => navigate("/smash")}>
                    <Typography variant="h5">Start smashing!</Typography>
                </Button>
        </ContentWrapper></>)}

        {status === "error" && (<><ContentWrapper>
            <Typography variant="h3" sx={{ mb: 2}}>Verification Failed🐸</Typography>
            <Typography variant="subtitle1" sx={{mb: 2}}>The verification link is invalid or has expired. Please request a new verification email by clicking the magical button below.</Typography>
            <ResendVerificationButton />
        </ContentWrapper></>)}
    </>
    );
}

export default VerificationPage;