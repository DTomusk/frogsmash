import { GlobalStyles, useTheme, CssBaseline } from "@mui/material";

export default function Background() {
  const theme = useTheme();

  return (
    <>
      <CssBaseline />
      <GlobalStyles
        styles={{
          body: {
            margin: 0,
            minHeight: "100vh",
            background: theme.palette.background.default,
            backgroundAttachment: "fixed",
            backgroundSize: "cover",
            color: theme.palette.text.primary,
          },
          "#root": { height: "100%" },
        }}
      />
    </>
  );
}
