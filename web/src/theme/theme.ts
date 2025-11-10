import { createTheme } from "@mui/material";

export const lightTheme = createTheme({
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

export const darkTheme = createTheme({
    typography: {
        fontFamily: "'Luckiest Guy', system-ui, Avenir, Helvetica, Arial, sans-serif",
    },
    palette: {
        mode: 'dark',
        primary: {
            main: '#327425ff',
        },
        secondary: {
            main: '#e8a106ff',
        },
        background: {
            default: '#121212',
        },
    },
});
