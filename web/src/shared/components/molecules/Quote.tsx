import { Box, Typography } from "@mui/material";

interface QuoteProps {
  text: string;
  author: string;
}

export default function Quote({ text, author }: QuoteProps) {
  return (
    <Box
      sx={{
        borderLeft: 4,
        borderColor: "primary.main",
        pl: 2,
        my: 4,
      }}
    >
      <Typography
        variant="h6"
        sx={{ fontStyle: "italic", mb: 1 }}
      >
        “{text}”
      </Typography>

      <Typography
        variant="subtitle2"
        sx={{ textAlign: "right", fontWeight: "bold" }}
      >
        — {author}
      </Typography>
    </Box>
  );
}
