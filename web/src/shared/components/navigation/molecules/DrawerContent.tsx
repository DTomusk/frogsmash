import { Box, List } from "@mui/material";
import ThemeSwitch from "./ThemeSwitch";
import { useAuth } from "@/app/providers";
import { DrawerLink, DrawerButton } from "../atoms/DrawerLink";

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
                    <DrawerLink to='/' label='Home' />
                    <DrawerLink to='/leaderboard' label='Leaderboard' />
                    <DrawerLink to='/upload' label='Upload' />
                    <DrawerButton onClick={logout} label='Logout' />
                </>) : <>
                    <DrawerLink to='/login' label='Login' />
                </>}
                <ThemeSwitch />
            </List>
        </Box>);
}

export default DrawerContent;