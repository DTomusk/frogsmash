import { TextField, type TextFieldProps } from "@mui/material";
import type { FieldError, UseFormRegisterReturn } from "react-hook-form";

interface EmailFieldProps extends Omit<TextFieldProps, "name"> {
  registration: UseFormRegisterReturn;
  fieldError?: FieldError;
}

function EmailField({ registration, fieldError, ...props }: EmailFieldProps) {
    return (
    <TextField
        label="Email"
        variant="outlined"
        required
        fullWidth
        sx={{ mb: 2 }}
        {...registration}
        error={!!fieldError}
        helperText={fieldError ? fieldError.message?.toString() : ""}
        {...props}
    />)
}

export default EmailField;