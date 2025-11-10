import { AppBar, Box, Drawer, IconButton, Toolbar, Typography } from "@mui/material";
import { Link } from "react-router-dom";
import MenuIcon from "@mui/icons-material/Menu";
import { useState } from "react";
import DrawerContent from "./DrawerContent";

function NavBar() {
    const [open, setOpen] = useState(false);
    const handleMenuToggle = (newOpen: boolean) => {
        setOpen(newOpen);
    };

    return (
        <>
        <Drawer open={open} 
            onClose={() => handleMenuToggle(false)} 
            anchor="right">
            <DrawerContent onClick={() => handleMenuToggle(false)} />
        </Drawer>
        <AppBar position="fixed" color="primary">
        <Toolbar sx={{ display: 'flex', justifyContent: 'space-between' }}>
            <Link to='/' style={{ textDecoration: 'none', color: 'inherit' }}>
                <Typography variant='h4'>
                    🐸 FrogSmash
                </Typography>
            </Link>
            <IconButton onClick={() => handleMenuToggle(true)}
                sx={{ display: { xs: "flex", sm: "none" } }} // show on small only
                color="inherit"
                >
                <MenuIcon />
            </IconButton>
            <Box sx={{ display: { xs: 'none', sm: 'flex' } }}>
                <Link to='/leaderboard' style={{ color: 'inherit' }}>
                    <Typography variant='h6'>
                        Leaderboard
                    </Typography>
                </Link>
            </Box>
        </Toolbar>
    </AppBar></>);
}

export default NavBar;