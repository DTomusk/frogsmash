import { Button, Typography } from "@mui/material";
import FormWrapper from "../atoms/FormWrapper";
import { useForm } from "react-hook-form";
import { useRegister } from "../../hooks/useRegister";
import StyledLink from "../atoms/StyledLink";
import EmailField from "../atoms/EmailField";
import PasswordField from "../atoms/PasswordField";
import PasswordStrength from "../atoms/PasswordStrength";
import { checkPasswordStrength } from "../../utils/PasswordStrength";
import { useLogin, type LoginResponse } from "../../hooks/useLogin";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../../contexts/AuthContext";
import { useSnackbar } from "../../contexts/SnackbarContext";

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
    const { mutate, isPending } = useRegister();
    const { mutate: login } = useLogin();
    const { login: authLogin } = useAuth();
    const navigate = useNavigate();
    const { showSnackbar } = useSnackbar();

    const password = watch("password") || "";
    const strength = checkPasswordStrength(password);
    const passwordValid = Object.values(strength).every(Boolean);

    const onSubmit = (data: RegistrationData) => {
        mutate(data, {
            onSuccess: () => {
              // TODO: add snackbar provider to show snackbar independent of route
                showSnackbar({ message: "Registration successful!", severity: "success",  });
                login({ email: data.email, password: data.password }, {
                    onSuccess: (response: LoginResponse) => {
                        authLogin(response.jwt);
                        navigate("/");
                    },
                    onError: (err: any) => {
                        showSnackbar({ message: err.message || "Login failed", severity: "error" });
                    }
                });
            },
            onError: (err: any) => {
              showSnackbar({ message: err.message || "Registration failed", severity: "error" });
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
    </>
  );
}

export default RegistrationPage;