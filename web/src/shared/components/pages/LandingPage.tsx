import { Box } from "@mui/material";
import IntroText from "../organisms/IntroText";
import TitleImage from "../organisms/TitleImage";
import { useTenant } from "@/app/providers/TenantProvider";

function LandingPage() {
    const config = useTenant();
    return <Box sx={{display: "flex", flexDirection: "column", alignItems: "center", width: "100%", justifyContent: "flex-start"}}>
                <TitleImage />
                {config.tenantKey === "frog" && <IntroText />}
            </Box>;
}

export default LandingPage;