import { Box, Divider, IconButton, List } from "@mui/material";
import ThemeSwitch from "./ThemeSwitch";
import { useAuth } from "@/app/providers";
import { DrawerLink, DrawerButton } from "../atoms/DrawerLink";
import { Close, Home } from "@mui/icons-material";
import EmojiEventsIcon from '@mui/icons-material/EmojiEvents';
import UploadIcon from '@mui/icons-material/Upload';
import PersonIcon from '@mui/icons-material/Person';

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
                <DrawerLink to='/' label='Home' onClick={onClick} ><Home /></DrawerLink>
                {token ? (
                    <>
                    <DrawerLink to='/leaderboard' label='Leaderboard' onClick={onClick} ><EmojiEventsIcon /></DrawerLink>
                    <DrawerLink to='/upload' label='Upload' onClick={onClick} ><UploadIcon /></DrawerLink>
                    <DrawerButton onClick={logoutClick} label='Logout'><PersonIcon  /></DrawerButton>
                </>) : <>
                    <DrawerLink to='/login' label='Login' onClick={onClick} ><PersonIcon /></DrawerLink>
                </>}
                <ThemeSwitch />
            </List>
        </Box>);
}

export default DrawerContent;