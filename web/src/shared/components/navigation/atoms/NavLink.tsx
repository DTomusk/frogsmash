import { Typography } from "@mui/material";
import { Link } from "react-router-dom";

interface NavLinkProps {
    to: string;
    label: string;
}

export function NavLink({ to, label }: NavLinkProps) {
    return (
        <Link to={to} style={{ color: 'inherit' }}>
            <Typography variant='h6'>
                {label}
            </Typography>
        </Link>
    )
}

interface NavButtonProps {
    onClick: () => void;
    label: string;
}

export function NavButton({ onClick, label }: NavButtonProps) {
    return (
        <Typography variant='h6' onClick={onClick} style={{ cursor: 'pointer' }}>
            {label}
        </Typography>
    )
}