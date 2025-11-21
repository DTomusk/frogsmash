import { createContext, useCallback, useContext, useState } from "react";
import type { AlertSnackbarProps } from "../../shared/components/molecules/AlertSnackbar";
import AlertSnackbar from "../../shared/components/molecules/AlertSnackbar";

type SnackbarContextType = {
    showSnackbar: (props: AlertSnackbarProps) => void;
};

const SnackbarContext = createContext<SnackbarContextType | undefined>(undefined);

export const useSnackbar = () => {
    const context = useContext(SnackbarContext);
    if (!context) {
        throw new Error("useSnackbar must be used within a SnackbarProvider");
    }
    return context;
}

export const SnackbarProvider = ({ children }: { children: React.ReactNode }) => {
    const [snackbars, setSnackbars] = useState<AlertSnackbarProps[]>([]);

    const showSnackbar = useCallback((options: AlertSnackbarProps) => {
        setSnackbars((prev) => [...prev, { ...options }]);
    }, []);

    const handleClose = (index: number) => {
        setSnackbars((prev) => prev.filter((_, i) => i !== index));
    }

    return (
        <SnackbarContext.Provider value={{ showSnackbar }}>
            {children}
            {snackbars.map((props, index) => (
                <AlertSnackbar
                    key={index}
                    {...props}
                    onClose={() => handleClose(index)}
                />
            ))}
        </SnackbarContext.Provider>
    );
}