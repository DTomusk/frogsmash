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

const lightFrogTheme = createTheme({
    typography: typography,
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

const darkFrogTheme = createTheme({
    typography: typography,
    palette: {
        mode: 'dark',
        primary: {
            main: '#1f9408ff',
        },
        secondary: {
            main: '#d36a00ff',
        },
        background: {
            default: '#121212ff',
        },
        success: {
            main: '#257829ff',
        },
    },
});

export const frogThemes = {
  light: lightFrogTheme,
  dark: darkFrogTheme,
};
