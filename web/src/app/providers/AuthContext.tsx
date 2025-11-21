import { createContext, useContext, useState } from "react";
import type { User } from "../../features/auth/models/user";

interface AuthContextType {
    token: string | null;
    user: User | null;
    login: (jwt: string, user?: User) => void;
    logout: () => void;
}

const AuthContext = createContext<AuthContextType>({
    token: null,
    user: null,
    login: () => {},
    logout: () => {},
});

export const AuthProvider = ({children}: {children: React.ReactNode}) => {
    const [token, setToken] = useState<string | null>(() => localStorage.getItem("token"));
    const [user, setUser] = useState<User | null>(null);

    const login = (jwt: string, userData?: User) => {
        localStorage.setItem("token", jwt);
        setToken(jwt);
        if (userData) {
            setUser(userData);
        }
    };

    const logout = () => {
        localStorage.removeItem("token");
        setToken(null);
        setUser(null);
    };

    return (
    <AuthContext.Provider value={{ user, token, login, logout }}>
        {children}
    </AuthContext.Provider>
  );
}

export const useAuth = () => {
    const context = useContext(AuthContext);
    if (!context) {
        throw new Error("useAuth must be used within an AuthProvider");
    }
    return context;
}