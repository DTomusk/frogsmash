import { Home } from "@mui/icons-material";
import { Box, List, ListItem, ListItemButton, ListItemIcon, ListItemText } from "@mui/material";
import { Link } from "react-router-dom";

interface DrawerContentProps {
    onClick: () => void;
}

function DrawerContent({ onClick }: DrawerContentProps) {
    return (
        <Box sx={{ width: 250 }} role="presentation" onClick={onClick} onKeyDown={onClick}>
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
            </Box>);
}

export default DrawerContent;