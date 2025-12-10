import { Typography } from "@mui/material";

interface HeaderTextProps {
    text: string;
}

export default function HeaderText({ text }: HeaderTextProps) {
    return (
        <Typography variant="h4" sx={{ textAlign: 'center', mt: 4, mx: 2 }}>
            {text}
        </Typography>
    );
}