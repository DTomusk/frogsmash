import { Box, Typography } from "@mui/material";
import { Link } from "react-router-dom";
import ThemeSwitch from "./ThemeSwitch";

function NavLinks() {
    return (
        <Box sx={{ display: 'flex', gap: 4, alignItems: 'center' }}>
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
            <ThemeSwitch />
        </Box>
    );
}

export default NavLinks;