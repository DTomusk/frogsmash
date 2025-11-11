import { Typography } from "@mui/material";
import { Link } from "react-router-dom";
import ThemeSwitch from "./ThemeSwitch";

function NavLinks() {
    return (
        <>
            <Link to='/leaderboard' style={{ color: 'inherit' }}>
                <Typography variant='h6'>
                    Leaderboard
                </Typography>
            </Link>
            <ThemeSwitch />
        </>
    );
}

export default NavLinks;