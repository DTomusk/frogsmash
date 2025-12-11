import { Typography } from "@mui/material";

interface BodyTextProps {
    text: string;
}

export default function BodyText({ text }: BodyTextProps) {
    return (
        <Typography variant="body1" sx={{ textAlign: 'center', mt: 4, mx: 2 }}>
            {text}
        </Typography>
    );
}