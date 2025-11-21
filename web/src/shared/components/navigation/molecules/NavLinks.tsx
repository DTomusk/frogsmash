import { Box, Button, Typography } from "@mui/material";
import { Link } from "react-router-dom";
import ThemeSwitch from "./ThemeSwitch";
import { useAuth } from "../../../../app/providers/AuthContext";

function NavLinks() {
    const { token, logout } = useAuth();
    return (
        <Box sx={{ display: 'flex', gap: 4, alignItems: 'center' }}>
            {token ? (
                <>
                <Link to='/leaderboard' style={{ color: 'inherit' }}>
                    <Typography variant='h6'>
                        Leaderboard
                    </Typography>
                </Link>
                <Link to='/upload' style={{ color: 'inherit' }}>
                    <Typography variant='h6'>
                        Upload
                    </Typography>
                </Link>
                <Button color="inherit" onClick={logout}>
                    <Typography variant='h6'>
                        Logout
                    </Typography>
                </Button>
                </>
            ) : (
                <>
                <Link to='/login' style={{ color: 'inherit' }}>
                    <Typography variant='h6'>
                        Login
                    </Typography>
                </Link>
                </>
            )}
            <ThemeSwitch />
        </Box>
    );
}

export default NavLinks;