import { Box } from "@mui/material";
import ThemeSwitch from "./ThemeSwitch";
import { useAuth } from "@/app/providers";
import { NavLink, NavButton }from "../atoms/NavLink";

function NavLinks() {
    const { token, logout } = useAuth();
    return (
        <Box sx={{ display: 'flex', gap: 4, alignItems: 'center' }}>
            {token ? (
                <>
                <NavLink to='/leaderboard' label='Leaderboard' />
                <NavLink to='/upload' label='Upload' />
                <NavButton onClick={logout} label='Logout' />
                </>
            ) : (
                <>
                <NavLink to='/login' label='Login' />
                </>
            )}
            <ThemeSwitch />
        </Box>
    );
}

export default NavLinks;