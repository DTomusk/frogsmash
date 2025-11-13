import { Button, TextField, Typography } from "@mui/material";
import FormWrapper from "../FormWrapper";
import { useForm } from "react-hook-form";
import { useRegister } from "../../hooks/useRegister";
import AlertSnackbar from "../AlertSnackbar";
import { useState } from "react";
import StyledLink from "../StyledLink";

interface RegistrationData {
    email: string;
    username: string;
    password: string;
}

function RegistrationPage() {
    const {
        register,
        handleSubmit,
        formState: { errors },
        reset,
    } = useForm<RegistrationData>();
    const { mutate, data, isPending } = useRegister();
    const [errorMessage, setErrorMessage] = useState("");
    const [openError, setOpenError] = useState(false);
    const [openSuccess, setOpenSuccess] = useState(false);

    const onSubmit = (data: RegistrationData) => {
        mutate(data, {
            onSuccess: () => {
                setOpenSuccess(true);
                reset();
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
        label="Username"
        variant="outlined"
        required
        fullWidth
        sx={{ mb: 2 }}
        {...register("username", { 
            required: "Username is required",
            minLength: { value: 3, message: "Username must be at least 3 characters" },
            pattern: { value: /^[a-zA-Z0-9_]+$/, message: "Username can only contain letters, numbers, and underscores" },
        })}
        error={!!errors.username}
        helperText={errors.username ? errors.username.message?.toString() : ""}
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