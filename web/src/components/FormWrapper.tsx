import { Paper } from "@mui/material";

function FormWrapper({ children, onSubmit }: { children: React.ReactNode; onSubmit: (event: React.FormEvent<HTMLFormElement>) => void }) {
    return (
        <Paper
        component="form"
        onSubmit={onSubmit}
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
export default FormWrapper;