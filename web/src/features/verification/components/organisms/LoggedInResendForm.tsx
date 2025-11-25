import { ContentWrapper } from "@/shared";
import { Typography } from "@mui/material";
import ResendVerificationButton from "../molecules/ResendVerificationButton";

interface LoggedInResendFormProps {
    title: string;
    message: string;
}

export default function LoggedInResendForm({ title, message }: LoggedInResendFormProps) {
    return (
        <ContentWrapper>
            <Typography variant='h3'>{title}</Typography>
            <Typography variant='body1' sx={{ mb: 2}}>{message}</Typography>
            <ResendVerificationButton />
        </ContentWrapper>
    );
}