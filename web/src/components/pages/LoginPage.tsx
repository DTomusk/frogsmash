import { Button, TextField, Typography } from "@mui/material";
import FormWrapper from "../FormWrapper";
import { useForm } from "react-hook-form";
import { useLogin, type LoginResponse } from "../../hooks/useLogin";
import AlertSnackbar from "../AlertSnackbar";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import StyledLink from "../StyledLink";
import { useAuth } from "../../contexts/AuthContext";

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
    const [errorMessage, setErrorMessage] = useState("");
    const [openError, setOpenError] = useState(false);

    const { mutate: login, isPending } = useLogin();

    const { login: authLogin } = useAuth();

    const navigate = useNavigate();

    const onSubmit = (data: LoginData) => {
        login(data, {
            onSuccess: (response: LoginResponse) => {
                // TODO: add user data to api response and pass it here
                authLogin(response.jwt);
                navigate("/");
            },
            onError: (err: any) => {
                setErrorMessage(err.message || "Login failed");
                setOpenError(true);
            }
        });
    };

    return <>
        <FormWrapper onSubmit={handleSubmit(onSubmit)}>
            <Typography variant="h3">Login</Typography>
            <Typography variant="body1" sx={{ mb: 2 }}>Don't have an account? Click <StyledLink to="/register" text="here" /> to register.</Typography>
            <TextField
                label="Email"
                variant="outlined"
                required
                fullWidth
                sx={{ mb: 2 }}
                {...register("email", {
                required: "Email is required",
                pattern: { value: /\S+@\S+\.\S+/, message: "Invalid email format" },
                })}
                error={!!errors.email}
                helperText={errors.email ? errors.email.message?.toString() : ""}
            />
            <TextField
                label="Password"
                type="password"
                variant="outlined"
                required
                fullWidth
                sx={{ mb: 2 }}
                {...register("password", { 
                    required: "Password is required",
                    minLength: { value: 8, message: "Password must be at least 8 characters" },
                    pattern: {
                        value: /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]).{8,}$/,
                        message: "Password must include upper, lower, number, and special character",
                    }
                })}
                error={!!errors.password}
                helperText={errors.password ? errors.password.message?.toString() : ""}
            />
            <Button type="submit" variant="contained" color="primary" fullWidth loading={isPending} disabled={isPending || Object.keys(errors).length > 0}>
                Login
            </Button>
        </FormWrapper>
        <AlertSnackbar
            open={openError}
            message={errorMessage}
            severity="error"
            onClose={() => setOpenError(false)}
        />
    </>;
}

export default LoginPage;