import { ContentWrapper } from "@/shared";
import { Typography } from "@mui/material";

export default function PendingView() {
    return (
        <ContentWrapper>
            <Typography variant='h3'>Your verification is pending...</Typography>
            <Typography variant='body1' sx={{ mb: 2}}>Please check your email for the verification link. If you haven't received it yet, please be patient as it may take a few minutes to arrive.</Typography>
        </ContentWrapper>
    );
}