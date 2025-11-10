import { Box, Typography } from "@mui/material";
import { Link } from "react-router-dom";

function Footer() {
  return (
    <Box component='footer'
    sx={{
        width: '100%',
        bgcolor: 'primary.main',
        position: 'relative',
        zIndex: 1000,
        display: 'flex',
        justifyContent: 'space-between'
    }}>
        <Box>
            <Link to="/privacy" style={{ textDecoration: 'none' }}>
                <Typography variant="body2" color="primary.contrastText" sx={{ px: 2, py: 2, display: 'inline-block' }}>
                    Privacy Policy
                </Typography>
            </Link>
            <Link to="/terms" style={{ textDecoration: 'none' }}>
                <Typography variant="body2" color="primary.contrastText" sx={{ px: 2, py: 2, display: 'inline-block' }}>
                    Terms of Service
                </Typography>
            </Link>
        </Box>
        <Typography variant="body2" color="primary.contrastText" align="center" sx={{ py: 2 }}>
            © {new Date().getFullYear()} FrogSmash. All rights reserved.
        </Typography>
    </Box>
  );
}

export default Footer;