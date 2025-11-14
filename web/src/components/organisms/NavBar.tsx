import { AppBar, Box, Drawer, IconButton, Toolbar } from "@mui/material";
import MenuIcon from "@mui/icons-material/Menu";
import { useState } from "react";
import DrawerContent from "../molecules/DrawerContent";
import LogoLink from "../molecules/LogoLink";
import NavLinks from "../molecules/NavLinks";

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
                <LogoLink />
                <IconButton onClick={() => handleMenuToggle(true)}
                    sx={{ display: { xs: "flex", sm: "none" } }} // show on small only
                    color="inherit"
                    >
                    <MenuIcon />
                </IconButton>
                <Box sx={{ display: { xs: 'none', sm: 'flex' } }}>
                    <NavLinks />
                </Box>
            </Toolbar>
        </AppBar>
    </>);
}

export default NavBar;