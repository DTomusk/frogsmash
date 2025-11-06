import { createTheme } from "@mui/material";

const theme = createTheme({
    typography: {
        fontFamily: "'Luckiest Guy', system-ui, Avenir, Helvetica, Arial, sans-serif",
    },
    palette: {
        primary: {
            main: '#327425ff',
        },
        secondary: {
            main: '#e8a106ff',
        },
        background: {
            default: '#f0f0f0ff',
        },
    },
});
export default theme;