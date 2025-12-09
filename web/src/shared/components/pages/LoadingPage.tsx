import { Box, Typography } from "@mui/material";
import LoadingSpinner from "../molecules/LoadingSpinner";
import { useTenant } from "@/app/providers/TenantProvider";

function LoadingPage () {
    const config = useTenant();
    return (
        <Box sx={{ display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center', mt: "64px" }}>
            <Typography variant="h4" sx={{ textAlign: 'center', mb: 4 }}>For those who wish to watch the {config.tenantKey === "book" ? "book" : "frog"} spin forever</Typography>
            <LoadingSpinner />
        </Box>
    );
}

export default LoadingPage;