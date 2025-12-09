import { Box, Typography } from "@mui/material";

function NotFoundPage() {
    return (
        <Box sx={{ textAlign: 'center', mt: 8 }}>
            <Typography variant="h1">404 - Page Not Found</Typography>
            <Typography variant="body1">Oopsie poopsie, this page does not exist!🫣</Typography>
        </Box>
    );
}

export default NotFoundPage;