import { Box } from "@mui/material";
import ThemeSwitch from "./ThemeSwitch";
import { useAuth } from "@/app/providers";
import { NavLink, NavButton }from "../atoms/NavLink";
import { useTenant } from "@/app/providers/TenantProvider";

function NavLinks() {
    const { token, logout } = useAuth();
    const config = useTenant();
    return (
        <Box sx={{ display: 'flex', gap: 4, alignItems: 'center' }}>
            {token ? (
                <>
                <NavLink to='/smash' label='Smash' />
                <NavLink to='/leaderboard' label='Leaderboard' />
                {config.tenantKey === "frog" && <NavLink to='/upload' label='Upload' />}
                <NavButton onClick={logout} label='Logout' />
                </>
            ) : (
                <>
                <NavLink to='/login' label='Login' />
                </>
            )}
            {config.tenantKey === "frog" && <ThemeSwitch />}
        </Box>
    );
}

export default NavLinks;