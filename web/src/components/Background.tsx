import { GlobalStyles, useTheme, CssBaseline } from "@mui/material";

export default function Background() {
  const theme = useTheme();

  // const gradient = `linear-gradient(180deg, 
  //   ${theme.palette.secondary.light} 0%,
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
