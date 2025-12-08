import { useTenant } from "@/app/providers/TenantProvider";
import { Typography } from "@mui/material";
import { Link } from "react-router-dom";

function LogoLink() {
    const config = useTenant();

    return (<Link to='/smash' style={{ textDecoration: 'none', color: 'inherit' }}>
                <Typography variant='h4'>
                    {config.title}
                </Typography>
            </Link>)
}

export default LogoLink;