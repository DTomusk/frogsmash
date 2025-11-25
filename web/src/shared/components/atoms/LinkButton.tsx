import { Button, Typography } from "@mui/material";
import { useNavigate } from "react-router-dom";

interface LinkButtonProps {
    color: "primary" | "secondary" | "success" | "error" | "info" | "warning";
    text: string;
    to: string;
}

export default function LinkButton({ color, text, to }: LinkButtonProps) {
    const navigate = useNavigate();
    return (
        <Button variant="contained" 
            size="large" 
            color={color} 
            sx={{ mt: 4, px: 2 }}
            onClick={() => navigate(to)}>
            <Typography variant="h5">{text}</Typography>
        </Button>
    );
}