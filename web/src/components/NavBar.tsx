import { AppBar, Toolbar, Typography } from "@mui/material";
import { Link } from "react-router-dom";

function NavBar() {
    return (<AppBar position="fixed" color="primary">
        <Toolbar>
            <Link to='/' style={{ textDecoration: 'none', color: 'inherit' }}>
            <Typography variant='h4'>
                🐸 FrogSmash
            </Typography>
            </Link>
        </Toolbar>
    </AppBar>);
}

export default NavBar;