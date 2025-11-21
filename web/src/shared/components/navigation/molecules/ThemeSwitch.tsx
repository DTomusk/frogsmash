import { Box, Switch, useTheme, styled } from "@mui/material";
import { useThemeMode } from "../../../theme/ThemeProvider";

const StyledSwitch = styled(Switch)(() => ({
    width: 56,
    height: 32,
    padding: 8,
    "& .MuiSwitch-thumb": {
        width: 24,
        height: 24,
        position: "relative",
        "&::before": {
        content: "'🌞'",
        position: "absolute",
        top: "50%",
        left: "50%",
        transform: "translate(-50%, -50%)",
        fontSize: "20px",
        },
    },
    "& .MuiSwitch-switchBase": {
        padding: 3,
        "&.Mui-checked": {
            transform: "translateX(22px)",
        },
    },
    "& .Mui-checked .MuiSwitch-thumb::before": {
        content: "'🌜'", 
    },
    "& .MuiSwitch-track": {
        borderRadius: 38 / 2,
    },
}));

function ThemeSwitch() {
    const theme = useTheme();
    const { toggleColorMode } = useThemeMode();

    return (
        <Box sx={{ display: 'flex', alignItems: 'center', ml: 2 }}>
            <StyledSwitch checked={theme.palette.mode === 'dark'} onChange={toggleColorMode} />
        </Box>
    );
}

export default ThemeSwitch;