import { Box, List, ListItem, ListItemIcon, ListItemText, Typography } from "@mui/material";
import { checkPasswordStrength } from "../../utils/PasswordStrength";
import CheckIcon from "@mui/icons-material/Check";
import CloseIcon from "@mui/icons-material/Close";

interface PasswordStrengthProps {
    password: string;
}

const labels = {
    minLength: "at least 8 characters",
    upperCase: "at least one uppercase letter (A-Z)",
    lowerCase: "at least one lowercase letter (a-z)",
    number: "at least one number (0-9)",
    specialChar: "at least one special character (!@#$%^&*)",
}

function PasswordStrength({ password }: PasswordStrengthProps) {
    const result = checkPasswordStrength(password);
    return (
        <Box sx={{ width: '100%'}}>
        <Typography variant="body2" sx={{mb: 1 }}>
            Make sure your password has:
        </Typography>
        <List dense sx={{ py: 0, my: 0 }}>
            {Object.entries(result).map(([key, passed]) => (
                <ListItem key={key} sx={{ p: 0, my: 0 }}>
                    <ListItemIcon>
                        {passed ? <CheckIcon color="success" /> : <CloseIcon color="error" />}
                    </ListItemIcon>
                    <ListItemText primary={labels[key as keyof typeof labels]} />
                </ListItem>
            ))}
        </List>
        </Box>
    );
}

export default PasswordStrength;