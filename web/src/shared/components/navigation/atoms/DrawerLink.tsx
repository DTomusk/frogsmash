import { Home } from "@mui/icons-material";
import { ListItem, ListItemButton, ListItemIcon, ListItemText } from "@mui/material";
import { Link } from "react-router-dom";

interface DrawerLinkProps {
    to: string;
    label: string;
    onClick: () => void;
}

export function DrawerLink({ to, label, onClick }: DrawerLinkProps) {
    return (
        <Link to={to} style={{ textDecoration: 'none', color: 'inherit' }}>
            <ListItem disablePadding>
                <ListItemButton onClick={onClick}>
                    <ListItemIcon>
                        <Home />
                    </ListItemIcon>
                    <ListItemText primary={label} />
                </ListItemButton>
            </ListItem>
        </Link>
    )
}

interface DrawerButtonProps {
    onClick: () => void;
    label: string;
}

export function DrawerButton({ onClick, label }: DrawerButtonProps) {
    return (
        <ListItem disablePadding>
            <ListItemButton onClick={onClick}>
                <ListItemIcon>
                    <Home />
                </ListItemIcon>
                <ListItemText primary={label} />
            </ListItemButton>
        </ListItem>
    )
}