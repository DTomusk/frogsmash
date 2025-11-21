import { Alert, Snackbar } from "@mui/material";

export interface AlertSnackbarProps {
    severity: "error" | "warning" | "info" | "success";
    message: string;
    autoHideDuration?: number;
    onClose?: () => void;
}

function AlertSnackbar({ severity, message, autoHideDuration = 6000, onClose }: AlertSnackbarProps) {
    return (
        <Snackbar open autoHideDuration={autoHideDuration} onClose={onClose} sx={{ boxShadow: 3 }}>
            <Alert onClose={onClose} severity={severity} variant="filled" sx={{ width: '100%' }}>
                {message}
            </Alert>
        </Snackbar>
    );
}

export default AlertSnackbar;
