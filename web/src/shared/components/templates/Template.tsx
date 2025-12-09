import { Box } from "@mui/material";
import { Outlet } from "react-router-dom";
import NavBar from "../navigation/organisms/NavBar";
import Footer from "../navigation/organisms/Footer";

function Template() {
    return (
         <Box sx={{ display: 'flex', flexDirection: 'column', width: '100vw', minHeight: '100vh' }}>
            <NavBar />
            <Box sx={{ display: "flex", flexDirection: "column", alignItems: "center", justifyContent: "flex-start", flexGrow: 1, mt: { xs: 6, sm: 8 }, mb: 4 }}>
                <Outlet />
            </Box>
            <Footer />
        </Box>
    );
}

export default Template;