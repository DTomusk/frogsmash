import { AppBar, Toolbar, Typography } from "@mui/material";

function NavBar() {
    return (<AppBar position="static" color="primary">
        <Toolbar>
            <Typography variant='h5'>
                🐸 FrogSmash
            </Typography>
        </Toolbar>
    </AppBar>);
}

export default NavBar;