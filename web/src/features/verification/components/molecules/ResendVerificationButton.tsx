import Button from "@mui/material/Button";
import { useResendVerification } from "../../hooks/useVerify";
import { Alert, Typography } from "@mui/material";
import { useState } from "react";

function ResendVerificationButton() {
    const [state, setState] = useState<"idle" | "success" | "error">("idle");
    const { mutate, isPending } = useResendVerification();

    const handleClick = () => {
        mutate(undefined, {
            onSuccess: () => {
                setState("success");
            },
            onError: () => {
                setState("error");
            }
        });  
    };

    return (<>
        {state === "idle" && (
        <Button variant="contained" color="primary" size="large" onClick={handleClick} disabled={isPending} loading={isPending}>
            <Typography variant="h5">Resend verification email</Typography>
        </Button>)}

        {state === "success" && (
            <Alert severity="success" variant="filled" sx={{ p: 2 }}>Verification email sent! Please check your inbox.</Alert>
        )}
        {state === "error" && (
            <Alert severity="error" variant="filled" sx={{ p: 2 }}>Failed to send verification email. Please try again later.</Alert>
        )}
    </>
    )
};

export default ResendVerificationButton;