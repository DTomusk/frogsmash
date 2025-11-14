import { Alert } from "@mui/material";

interface ErrorMessageProps {
    message: string;
}

function ErrorMessage({ message }: ErrorMessageProps) {
    return <Alert severity="error" variant="filled">{message}</Alert>;
}

export default ErrorMessage;