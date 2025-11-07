import { Box } from "@mui/material";
import { Outlet } from "react-router-dom";
import NavBar from "../NavBar";

function Template() {
    return (
         <Box sx={{ display: 'flex', flexDirection: 'column', width: '100vw', minHeight: '100vh' }}>
            <NavBar />
            <Box sx={{ display: "flex", flexDirection: "column", alignItems: "center", justifyContent: "center", flexGrow: 1, mt: 8 }}>
                <Outlet />
            </Box>
        </Box>
    );
}

export default Template;