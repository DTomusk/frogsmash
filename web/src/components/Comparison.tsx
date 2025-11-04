import { Box, Typography } from "@mui/material";
import { useItems } from "../hooks/useItems"
import Contender from "./Contender";

function Comparison() {
  const { isPending, error, data } = useItems();

  if (isPending) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error loading items.</div>;
  }

  const { left_item, right_item } = data;

  return (
    <Box>
        <Typography variant="h4" sx={{ textAlign: 'center', mb: 4 }}>Ribbit ribbit</Typography>
        <Box sx={{ display: 'flex', gap: 4, flexDirection: { xs: 'column', md: 'row' } }}>
            <Contender imageUrl={left_item.image_url} name={left_item.name} />
            <Typography variant="h3" sx={{ alignSelf: 'center' }}>VS</Typography>
            <Contender imageUrl={right_item.image_url} name={right_item.name} />
        </Box>
    </Box>
  );
}

export default Comparison;
