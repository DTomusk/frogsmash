import { Visibility, VisibilityOff } from "@mui/icons-material";
import { IconButton, InputAdornment, TextField, type TextFieldProps } from "@mui/material";
import { useState } from "react";
import type { FieldError, UseFormRegisterReturn } from "react-hook-form";

interface PasswordFieldProps extends Omit<TextFieldProps, "name"> {
  registration: UseFormRegisterReturn;
  fieldError?: FieldError;
}

function PasswordField({ registration, fieldError, ...props }: PasswordFieldProps) {
    const [showPassword, setShowPassword] = useState(false);

    const handleClickShowPassword = () => {
        setShowPassword((show) => !show);
    }
    return (
    <TextField
        label="Password"
        variant="outlined"
        required
        fullWidth
        type={showPassword ? "text" : "password"}
        sx={{ mb: 2 }}
        {...registration}
        error={!!fieldError}
        helperText={fieldError ? fieldError.message?.toString() : ""}
        {...props}
        slotProps={{
            input: {
                endAdornment: <InputAdornment position="end">
                    <IconButton
                    aria-label={
                        showPassword ? 'hide the password' : 'display the password'
                    }
                    onClick={handleClickShowPassword}
                    >
                    {showPassword ? <VisibilityOff /> : <Visibility />}
                    </IconButton>
                </InputAdornment>,
            },
        }}
    />)
}

export default PasswordField;