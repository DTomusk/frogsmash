import { Alert, Snackbar } from "@mui/material";

interface AlertSnackbarProps {
    open: boolean;
    onClose: () => void;
    severity: "error" | "warning" | "info" | "success";
    message: string;
}

function AlertSnackbar({ open, onClose, severity, message }: AlertSnackbarProps) {
    return (
        <Snackbar open={open} autoHideDuration={6000} onClose={onClose}>
            <Alert onClose={onClose} severity={severity} variant="outlined" sx={{ width: '100%', bgcolor: 'background.paper' }}>
                {message}
            </Alert>
        </Snackbar>
    );
}

export default AlertSnackbar;
