import { ContentWrapper, LinkButton } from "@/shared";
import { Typography } from "@mui/material";

export default function AlreadyVerified() {
    return (
        <ContentWrapper>
            <Typography variant='h3'>Already verified</Typography>
            <Typography variant='body1' sx={{ mb: 2}}>Your account is already verified. This means you have access to all our special features, like uploading your own contender. Give it a try!</Typography>
            <LinkButton color="primary" text="Submit a contender!" to="/upload" />
        </ContentWrapper>
    );
}