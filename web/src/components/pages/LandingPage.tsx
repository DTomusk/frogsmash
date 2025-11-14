import { Box, Button, Paper, Typography, useTheme } from "@mui/material";

function LandingPage() {
    const theme = useTheme(); 
    return <Box sx={{display: "flex", flexDirection: "column", alignItems: "center"}}>
            <Box
            sx={{
                width: "100%",
                height: {xs: 200, sm: 400, md: 500, lg: 600, xl: 800},
                backgroundSize: "cover",
                display: "flex",
                backgroundImage: theme.palette.mode === 'light'
                    ? `url(/froggy_light.png)`
                    : `url(/froggy_dark.png)`,
            }}
            >
            <Box sx={{ width: "50%", height: "100%", display: "flex", flexDirection: "column", alignItems: "center", justifyContent: "center" }}>
                <Typography variant="h4" sx={{ textAlign: 'center', mt: 4, color: theme.palette.mode === 'light' ? 'black' : 'white' }}>
                    Compare frogs
                </Typography>
                <Typography variant="h4" sx={{ textAlign: 'center', mt: 4, color: theme.palette.mode === 'light' ? 'black' : 'white' }}>
                    Compare other things
                </Typography>
                <Typography variant="h4" sx={{ textAlign: 'center', mt: 4, color: theme.palette.mode === 'light' ? 'black' : 'white' }}>
                    Ribbit ribbit
                </Typography>
                <Button variant="contained" size="large" color={theme.palette.mode === 'light' ? 'primary' : 'secondary'} sx={{ mt: 4, px: 2 }}>
                    <Typography variant="h5">Start smashing!</Typography>
                </Button>
            </Box>
        </Box>
        <Box width="50%">
            <Paper elevation={3} sx={{ p: 2 }}>
                <Typography variant="h4" sx={{ textAlign: 'center', mb: 2, mt: 2 }}>
                    What is FrogSmash?
                </Typography>
                <Typography variant="body1" sx={{ textAlign: 'center', mt: 4, mx: 2 }}>
                    Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
                </Typography>
            </Paper>
        </Box>
    </Box>;
}

export default LandingPage;