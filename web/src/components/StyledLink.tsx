import { Typography } from "@mui/material";
import { Link } from "react-router-dom";

interface StyledLinkProps {
    to: string;
    text: string;
}

function StyledLink({ to, text }: StyledLinkProps) {
    return (<Link to={to}>
        <Typography component="span" sx={{ fontWeight: 'bold', color: 'primary.main', "&:hover": { textDecoration: 'underline' } }}>{text}</Typography>
    </Link>);
}

export default StyledLink;