import { createTheme } from "@mui/material";

const bodyFont = "system-ui, Avenir, Helvetica, Arial, sans-serif";
const headingFont = "'Luckiest Guy', cursive";
const typography = {
    fontFamily: bodyFont,
    h1: { fontFamily: headingFont },
    h2: { fontFamily: headingFont },
    h3: { fontFamily: headingFont },
    h4: { fontFamily: headingFont },
    h5: { fontFamily: headingFont },
    h6: { fontFamily: headingFont},
};

const lightBookTheme = createTheme({
    typography: typography,
    palette: {
        primary: {
            main: '#981a0fff',
        },
        secondary: {
            main: '#981a0fff',
        },
        background: {
            default: '#f0f0f0ff',
        },
    },
});

const darkBookTheme = createTheme({
    typography: typography,
    palette: {
        mode: 'dark',
        primary: {
            main: '#981a0fff',
        },
        secondary: {
            main: '#981a0fff',
        },
        background: {
            default: '#121212ff',
        },
    },
});


export const bookThemes = {
  light: lightBookTheme,
  dark: darkBookTheme,
};