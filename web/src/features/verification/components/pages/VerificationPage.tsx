import { Button, Typography } from "@mui/material";
import { ContentWrapper, LoadingSpinner } from "@/shared";
import { useNavigate, useSearchParams } from "react-router-dom";
import { useEffect, useState } from "react";
import ResendVerificationButton from "../organisms/ResendVerificationButton";
import { useVerifyCode } from "../../hooks/useVerify";

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
            onSuccess: () => {
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