import { Box, Divider, IconButton, List } from "@mui/material";
import ThemeSwitch from "./ThemeSwitch";
import { useAuth } from "@/app/providers";
import { DrawerLink, DrawerButton } from "../atoms/DrawerLink";
import { Close } from "@mui/icons-material";

interface DrawerContentProps {
    onClick: () => void;
}

function DrawerContent({ onClick }: DrawerContentProps) {
    const { token, logout } = useAuth();
    const logoutClick = () => {
        logout();
        onClick();
    }
    return (
        <Box sx={{ width: { xs: 300, sm: 250} }} role="presentation" onKeyDown={onClick}>
            <Box sx={{ display: 'flex', justifyContent: 'flex-end', p: 1 }}>
                <IconButton onClick={onClick}>
                    <Close />
                </IconButton>
            </Box>
            <Divider />
            <List>
                <DrawerLink to='/' label='Home' onClick={onClick} />
                {token ? (
                    <>
                    <DrawerLink to='/leaderboard' label='Leaderboard' onClick={onClick} />
                    <DrawerLink to='/upload' label='Upload' onClick={onClick} />
                    <DrawerButton onClick={logoutClick} label='Logout'/>
                </>) : <>
                    <DrawerLink to='/login' label='Login' onClick={onClick} />
                </>}
                <ThemeSwitch />
            </List>
        </Box>);
}

export default DrawerContent;