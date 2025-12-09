import { Box, Typography, useTheme } from "@mui/material";
import LinkButton from "../atoms/LinkButton";
import { useTenant } from "@/app/providers/TenantProvider";

export default function TitleImage() {
    const theme = useTheme();
    const config = useTenant();
    return (
        <Box
            sx={{
                width: "100%",
                height: {xs: 400, sm: 400, md: 500, lg: 600, xl: 800},
                backgroundSize: "cover",
                display: "flex",
                backgroundColor: theme.palette.mode === 'light' ? 'white' : 'primary.dark',
                backgroundImage: {xs: "none", 
                    sm:
                    theme.palette.mode === 'light'
                    ? `url(/froggy_light.png)`
                    : `url(/froggy_dark.png)`
                }
            }}
            >
            <Box sx={{ width: { xs: "100%", sm: "50%", md: "50%" }, height: "100%", display: "flex", flexDirection: "column", alignItems: "center", justifyContent: "center" }}>
                <Typography variant="h4" sx={{ textAlign: 'center', mt: 4, color: theme.palette.mode === 'light' ? 'black' : 'white' }}>
                    {config.titleImageText1}
                </Typography>
                <Typography variant="h4" sx={{ textAlign: 'center', mt: 4, color: theme.palette.mode === 'light' ? 'black' : 'white' }}>
                    {config.titleImageText2}
                </Typography>
                <Typography variant="h4" sx={{ textAlign: 'center', mt: 4, color: theme.palette.mode === 'light' ? 'black' : 'white' }}>
                    {config.titleImageText3}
                </Typography>
                <LinkButton color={theme.palette.mode === 'light' ? 'primary' : 'secondary'} text="Start smashing!" to="/smash" />
            </Box>
        </Box>
    )
}