import { Typography } from "@mui/material";
import { Link } from "react-router-dom";

function LogoLink() {
    return (<Link to='/smash' style={{ textDecoration: 'none', color: 'inherit' }}>
                <Typography variant='h4'>
                    🐸 FrogSmash
                </Typography>
            </Link>)
}

export default LogoLink;