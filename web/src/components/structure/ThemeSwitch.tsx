import { Box, Typography, Switch, useTheme } from "@mui/material";
import { useThemeMode } from "../../theme/ThemeProvider";

function ThemeSwitch() {
    const theme = useTheme();
    const { toggleColorMode } = useThemeMode();

    return (
        <Box sx={{ display: 'flex', alignItems: 'center', ml: 2 }}>
            { theme.palette.mode === 'light' ? 
                <Typography variant='h6'>Light mode</Typography> : 
                <Typography variant='h6'>Dark mode</Typography> 
            }
            <Switch checked={theme.palette.mode === 'dark'} onChange={toggleColorMode} />
        </Box>
    );
}

export default ThemeSwitch;