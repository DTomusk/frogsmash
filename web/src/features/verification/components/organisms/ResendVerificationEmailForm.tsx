import { EmailField, FormWrapper } from "@/shared";
import { Alert, Button, Typography } from "@mui/material";
import { useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import { useResendVerificationWithEmail } from "../../hooks/useVerify";

interface ResendVerificationEmailFormProps {
    title: string;
    message: string;
}

export default function ResendVerificationEmailForm({ title, message }: ResendVerificationEmailFormProps) {
    const {
        register,
        handleSubmit,
        watch,
        formState: { errors },
    } = useForm<{ email: string }>();

    const [state, setState] = useState<"idle" | "success" | "error">("idle");
    const { mutate, isPending } = useResendVerificationWithEmail();

    const email = watch("email");

    useEffect(() => {
        if (state !== "idle") {
            setState("idle");
        }
    }, [email]);

    const onSubmit = (data: { email: string }) => {
        // Handle resend verification email logic here
        console.log("Resend verification email to:", data.email);
        mutate(data.email, {
            onSuccess: () => {
                setState("success");
            },
            onError: () => {
                setState("error");
            }
        })
    }

    return (
        <FormWrapper onSubmit={handleSubmit(onSubmit)}>
            <Typography variant='h3'>{title}</Typography>
            <Typography variant='body1' sx={{ mb: 2}}>{message}</Typography>
            
            <EmailField
                registration={register("email", {
                    required: "Email is required",
                    pattern: { value: /\S+@\S+\.\S+/, message: "Please enter a valid email address" },
                })}
                fieldError={errors.email}
            />
            {state === "idle" && (
            <Button type="submit" variant="contained" color="primary" fullWidth loading={isPending} disabled={isPending ||Object.keys(errors).length > 0}>
                <Typography variant="h6">Resend email</Typography>
            </Button>)}

            {state === "success" && (
                <Alert severity="success" variant="filled" sx={{ p: 2 }}>Verification email sent! Please check your inbox.</Alert>
            )}
            {state === "error" && (
                <Alert severity="error" variant="filled" sx={{ p: 2 }}>Failed to send verification email. Please try again later.</Alert>
            )}
        </FormWrapper>
    );
}