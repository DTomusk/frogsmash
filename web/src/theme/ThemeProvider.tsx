import { CssBaseline, ThemeProvider } from "@mui/material";
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
    const [mode, setMode] = useState<ColorMode>(() => {
        const savedMode = localStorage.getItem('themeMode') as ColorMode | null;
        return savedMode ? savedMode : 'light';
    });

    const toggleColorMode = () => {
        setMode((prevMode) => (prevMode === 'light' ? 'dark' : 'light'));
    };

    const theme = useMemo(() => (mode === "light" ? lightTheme : darkTheme), [mode]);

    useEffect(() => {
        const storedMode = localStorage.getItem('themeMode');
        if (storedMode !== mode) {
            localStorage.setItem('themeMode', mode);
        }
    }, [mode]);

    return (
        <ThemeContext.Provider value={{ mode, toggleColorMode }}>
            <ThemeProvider theme={theme}>
                <CssBaseline />
                {children}
            </ThemeProvider>
        </ThemeContext.Provider>
    );
}