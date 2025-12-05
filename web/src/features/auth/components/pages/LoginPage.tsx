import { Button, Typography } from "@mui/material";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import { useLogin } from "../../hooks/useLogin";
import { useAuth, useSnackbar } from "@/app/providers";
import { FormWrapper, EmailField, StyledLink } from "@/shared";
import PasswordField from "../atoms/PasswordField";
import type { LoginResponse } from "../../dtos/loginResponse";
import GoogleSignInButton from "../atoms/GoogleSignInButton";
import { useGoogleLogin } from "../../hooks/useGoogleLogin";

interface LoginData {
    email: string;
    password: string;
}

function LoginPage() {
    const {
        register,
        handleSubmit,
        formState: { errors },
    } = useForm<LoginData>();

    const { mutate: login, isPending } = useLogin();
    const { mutate: googleLogin } = useGoogleLogin();
    const { login: authLogin } = useAuth();
    const { showSnackbar } = useSnackbar();
    const navigate = useNavigate();

    const onSubmit = (data: LoginData) => {
        login(data, {
            onSuccess: (response: LoginResponse) => {
                authLogin(response.jwt, response.user);
                showSnackbar({ message: "Login successful, welcome back!🎉", severity: "success" });
                navigate("/smash");
            },
            onError: (err: any) => {
                showSnackbar({ message: err.message || "Login failed", severity: "error" });
            }
        });
    };

    function handleLogin(credential: string) {
        console.log("Google ID Token:", credential);

        googleLogin({ idToken: credential }, {
            onSuccess: (response: LoginResponse) => {
                authLogin(response.jwt, response.user);
                showSnackbar({ message: "Login successful, welcome back!🎉", severity: "success" });
                navigate("/smash");
            },
            onError: (err: any) => {
                showSnackbar({ message: err.message || "Google Login failed", severity: "error" });
            }
        });
    }

    return <>
        <FormWrapper onSubmit={handleSubmit(onSubmit)}>
            <Typography variant="h3">Login</Typography>
            <Typography variant="body1" sx={{ mb: 2 }}>Don't have an account? Click <StyledLink to="/register" text="here" /> to register.</Typography>
            <EmailField
                registration={register("email", {
                    required: "Email is required",
                    pattern: { value: /\S+@\S+\.\S+/, message: "Invalid email format" },
                })}
                fieldError={errors.email}
            />
            <PasswordField
                registration={register("password", { 
                    required: "Password is required",
                })}
                fieldError={errors.password}
            />
            <Button type="submit" variant="contained" color="primary" fullWidth loading={isPending} disabled={isPending || Object.keys(errors).length > 0}>
                <Typography variant="h6">Login</Typography>
            </Button>
            <Typography variant="body1" sx={{ my: 2 }}>Or, alternatively, log in with Google:</Typography>
            <GoogleSignInButton onLogin={handleLogin} />
        </FormWrapper>
    </>;
}

export default LoginPage;