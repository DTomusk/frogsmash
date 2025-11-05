import { AppBar, Toolbar, Typography } from "@mui/material";

function NavBar() {
    return (<AppBar position="static">
        <Toolbar>
            <Typography variant='h5'>
                🐸 FrogSmash
            </Typography>
        </Toolbar>
    </AppBar>);
}

export default NavBar;