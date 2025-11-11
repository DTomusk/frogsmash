import { CssBaseline, ThemeProvider, useMediaQuery } from "@mui/material";
import { createContext, useContext, useEffect, useMemo, useState, type FC, type ReactNode } from "react";
import { darkTheme, lightTheme } from "./theme";

type ColorMode = 'light' | 'dark';

interface ThemeContextType {
  mode: ColorMode;
  toggleColorMode: () => void;
}

const ThemeContext = createContext<ThemeContextType>({
  mode: 'light',
  toggleColorMode: () => {},
});

export const useThemeMode = () => useContext(ThemeContext);

export const AppThemeProvider: FC<{ children: ReactNode }> = ({ children }) => {
    // Get system preference for theme
    const prefersDarkMode = useMediaQuery('(prefers-color-scheme: dark)');

    // Initialize theme mode from local storage or system preference
    // Local storage overrides system preference
    const [mode, setMode] = useState<ColorMode>(() => {
        const savedMode = localStorage.getItem('themeMode') as ColorMode | null;
        if (savedMode === 'light' || savedMode === 'dark') {
            return savedMode;
        }
        return prefersDarkMode ? 'dark' : 'light';
    });

    // Sync theme mode with local storage when system preference changes
    useEffect(() => {
        const storedMode = localStorage.getItem('themeMode');
        if (!storedMode) {
            setMode(prefersDarkMode ? 'dark' : 'light');
        }
    }, [prefersDarkMode]);

    // Set local storage on toggle
    const toggleColorMode = () => {
        setMode((prevMode) => {
            const nextMode = prevMode === 'light' ? 'dark' : 'light'
            localStorage.setItem('themeMode', nextMode);
            return nextMode;
        });
    };

    const theme = useMemo(
        () => (mode === "light" ? lightTheme : darkTheme), 
        [mode]
    );

    return (
        <ThemeContext.Provider value={{ mode, toggleColorMode }}>
            <ThemeProvider theme={theme}>
                <CssBaseline />
                {children}
            </ThemeProvider>
        </ThemeContext.Provider>
    );
}