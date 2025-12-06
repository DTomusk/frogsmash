import { ListItem, ListItemButton, ListItemIcon, ListItemText } from "@mui/material";
import { Link } from "react-router-dom";

interface DrawerLinkProps {
    to: string;
    label: string;
    onClick: () => void;
    children: React.ReactNode;
}

export function DrawerLink({ to, label, onClick, children }: DrawerLinkProps) {
    return (
        <Link to={to} style={{ textDecoration: 'none', color: 'inherit' }}>
            <ListItem disablePadding>
                <ListItemButton onClick={onClick}>
                    <ListItemIcon>
                        { children }
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
    children: React.ReactNode;
}

export function DrawerButton({ onClick, label, children }: DrawerButtonProps) {
    return (
        <ListItem disablePadding>
            <ListItemButton onClick={onClick}>
                <ListItemIcon>
                    { children}
                </ListItemIcon>
                <ListItemText primary={label} />
            </ListItemButton>
        </ListItem>
    )
}