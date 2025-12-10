import { Box, Paper } from "@mui/material";

export default function BigTextWrapper({ children }: { children: React.ReactNode }) {
    return (
        <Box width={{ xs: "90%", sm: "70%", md: "50%" }}>
            <Paper elevation={3} sx={{ p: 2, pb: 4, textAlign: 'center', justifyContent: 'center', mt: 4, mb: 4, display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
                {children}
            </Paper>
        </Box>
    );
}