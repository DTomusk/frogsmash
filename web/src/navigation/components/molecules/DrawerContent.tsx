import { Home } from "@mui/icons-material";
import { Box, List, ListItem, ListItemButton, ListItemIcon, ListItemText } from "@mui/material";
import { Link } from "react-router-dom";
import ThemeSwitch from "./ThemeSwitch";
import { useAuth } from "../../../auth/contexts/AuthContext";

interface DrawerContentProps {
    onClick: () => void;
}

function DrawerContent({ onClick }: DrawerContentProps) {
    const { token, logout } = useAuth();
    return (
        <Box sx={{ width: 250 }} role="presentation" onClick={onClick} onKeyDown={onClick}>
            <List>
                {token ? (
                    <>
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

                    <Link to='/upload' style={{ textDecoration: 'none', color: 'inherit' }}>
                        <ListItem disablePadding>
                            <ListItemButton>
                                <ListItemIcon>
                                    <Home />
                                </ListItemIcon>
                                <ListItemText primary="Upload" />
                            </ListItemButton>
                        </ListItem>
                    </Link>

                    <ListItem disablePadding>
                        <ListItemButton onClick={logout}>
                            <ListItemIcon>
                                <Home />
                            </ListItemIcon>
                            <ListItemText primary="Logout" />
                        </ListItemButton>
                    </ListItem>
                </>) : <>
                <Link to='/login' style={{ textDecoration: 'none', color: 'inherit' }}>
                    <ListItem disablePadding>
                        <ListItemButton>
                            <ListItemIcon>
                                <Home />
                            </ListItemIcon>
                            <ListItemText primary="Login" />
                        </ListItemButton>
                    </ListItem>
                </Link>
                </>}
                <ThemeSwitch />
            </List>
        </Box>);
}

export default DrawerContent;