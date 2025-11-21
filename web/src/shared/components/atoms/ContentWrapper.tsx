import { Paper } from "@mui/material";

function ContentWrapper({ children }: { children: React.ReactNode; }) {
    return (
        <Paper
        sx={{
            display: "flex",
            flexDirection: "column",
            gap: 1,
            alignItems: "center",
            m: 4,
            p: 4,
            borderRadius: 2,
            maxWidth: 600,
            width: "100%",
        }}
    >
        {children}
    </Paper>
    )
}
export default ContentWrapper;