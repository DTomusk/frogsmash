import { AppBar, Box, Drawer, IconButton, List, ListItem, ListItemButton, ListItemIcon, ListItemText, Toolbar, Typography } from "@mui/material";
import { Link } from "react-router-dom";
import MenuIcon from "@mui/icons-material/Menu";
import { useState } from "react";
import { Home } from "@mui/icons-material";

function NavBar() {
    const [open, setOpen] = useState(false);
    const handleMenuToggle = (newOpen: boolean) => {
        setOpen(newOpen);
    };

    const drawerContent = (
        <Box sx={{ width: 250 }} role="presentation" onClick={() => handleMenuToggle(false)} onKeyDown={() => handleMenuToggle(false)}>
            <List>
                <Link to='/' style={{ textDecoration: 'none', color: 'inherit' }}>
                <ListItem disablePadding>
                    <ListItemButton>
                        <ListItemIcon>
                            <Home />
                        </ListItemIcon>
                        <ListItemText primary="Home" />
                    </ListItemButton>
                </ListItem>
                </Link>
                <Link to='/leaderboard' style={{ textDecoration: 'none', color: 'inherit' }}>
                <ListItem disablePadding>
                    <ListItemButton>
                        <ListItemIcon>
                            <Home />
                        </ListItemIcon>
                        <ListItemText primary="Leaderboard" />
                    </ListItemButton>
                </ListItem>
                </Link>
            </List>
        </Box>
    )

    return (
        <>
        <Drawer open={open} 
            onClose={() => handleMenuToggle(false)} 
            anchor="right">
            {drawerContent}
        </Drawer>
        <AppBar position="fixed" color="primary">
        <Toolbar sx={{ display: 'flex', justifyContent: 'space-between' }}>
            <Link to='/' style={{ textDecoration: 'none', color: 'inherit' }}>
                <Typography variant='h4'>
                    🐸 FrogSmash
                </Typography>
            </Link>
            <IconButton onClick={() => handleMenuToggle(true)}
                sx={{ display: { xs: "flex", md: "none" } }} // show on small only
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