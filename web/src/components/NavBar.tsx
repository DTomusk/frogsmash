import { AppBar, Toolbar, Typography } from "@mui/material";

function NavBar() {
    return (<AppBar position="fixed" color="primary">
        <Toolbar>
            <Typography variant='h4'>
                🐸 FrogSmash
            </Typography>
        </Toolbar>
    </AppBar>);
}

export default NavBar;