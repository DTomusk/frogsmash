import { GlobalStyles, useTheme, CssBaseline } from "@mui/material";

export default function Background() {
  const theme = useTheme();

  // const gradient = `linear-gradient(110deg, 
  //   ${theme.palette.primary.main} 0%,
  //   ${theme.palette.primary.light} 45%,
  //   ${theme.palette.secondary.light} 55%,
  //   ${theme.palette.secondary.main} 100%)
  //   `;

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
