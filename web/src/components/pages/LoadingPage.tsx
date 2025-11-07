import { Box, Typography } from "@mui/material";
import LoadingSpinner from "../LoadingSpinner";

function LoadingPage () {
    return (
        <Box sx={{ display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center' }}>
            <Typography variant="h4" sx={{ textAlign: 'center', mb: 4 }}>For those who wish to watch the frog spin forever</Typography>
            <LoadingSpinner />
        </Box>
    );
}

export default LoadingPage;