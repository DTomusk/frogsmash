import { Button, Typography } from "@mui/material";
import FormWrapper from "../atoms/FormWrapper";
import { useForm } from "react-hook-form";
import { useRegister } from "../../hooks/useRegister";
import AlertSnackbar from "../molecules/AlertSnackbar";
import { useState } from "react";
import StyledLink from "../atoms/StyledLink";
import EmailField from "../atoms/EmailField";
import PasswordField from "../atoms/PasswordField";
import PasswordStrength from "../atoms/PasswordStrength";
import { checkPasswordStrength } from "../../utils/PasswordStrength";
import { useLogin, type LoginResponse } from "../../hooks/useLogin";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../../contexts/AuthContext";

interface RegistrationData {
    email: string;
    password: string;
}

function RegistrationPage() {
    const {
        register,
        handleSubmit,
        formState: { errors },
        watch,
    } = useForm<RegistrationData>();
    const { mutate, data, isPending } = useRegister();
    const { mutate: login } = useLogin();
    const { login: authLogin } = useAuth();
    const navigate = useNavigate();

    const [errorMessage, setErrorMessage] = useState("");
    const [openError, setOpenError] = useState(false);
    const [openSuccess, setOpenSuccess] = useState(false);

    const password = watch("password") || "";
    const strength = checkPasswordStrength(password);
    const passwordValid = Object.values(strength).every(Boolean);

    const onSubmit = (data: RegistrationData) => {
        mutate(data, {
            onSuccess: () => {
              // TODO: add snackbar provider to show snackbar independent of route
                setOpenSuccess(true);
                login({ email: data.email, password: data.password }, {
                    onSuccess: (response: LoginResponse) => {
                        authLogin(response.jwt);
                        navigate("/");
                    },
                    onError: (err: any) => {
                        setErrorMessage(err.message || "Login after registration failed");
                        setOpenError(true);
                    }
                });
            },
            onError: (err: any) => {
              setErrorMessage(err.message || "Registration failed");
              setOpenError(true);
            },
        });
    }
  
  return (
    <>
    <FormWrapper onSubmit={handleSubmit(onSubmit)}>
      <Typography variant="h3">Register</Typography>
      <Typography variant="body1" sx={{ mb: 2 }}>Already have an account? Click <StyledLink to="/login" text="here" /> to login.</Typography>
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
            validate: (value) => {
              const result = checkPasswordStrength(value);
              const allPassed = Object.values(result).every(Boolean);
              return allPassed || "Password does not meet requirements";
            }
        })}
        fieldError={errors.password}
      />
      <PasswordStrength password={watch("password") || ""} />
      <Button type="submit" variant="contained" color="primary" fullWidth loading={isPending} disabled={isPending || !passwordValid}>
        Register
      </Button>
    </FormWrapper>
    <AlertSnackbar
      open={openError}
      message={errorMessage}
      severity="error"
      onClose={() => setOpenError(false)}
    />
    <AlertSnackbar
      open={openSuccess}
      message={data?.message || "Registration successful!"}
      severity="success"
      onClose={() => setOpenSuccess(false)}
    />
    </>
  );
}

export default RegistrationPage;