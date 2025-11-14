import { Button, Typography } from "@mui/material";
import FormWrapper from "../FormWrapper";
import { useForm } from "react-hook-form";
import { useLogin, type LoginResponse } from "../../hooks/useLogin";
import AlertSnackbar from "../AlertSnackbar";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import StyledLink from "../StyledLink";
import { useAuth } from "../../contexts/AuthContext";
import EmailField from "../EmailField";
import PasswordField from "../PasswordField";

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