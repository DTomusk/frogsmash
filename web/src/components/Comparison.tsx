import { Box, Typography } from "@mui/material";
import { useItems } from "../hooks/useItems"
import Contender from "./Contender";
import { useComparison } from "../hooks/useComparison";

function Comparison() {
  const { isPending, error, data, refetch } = useItems();
  const { mutate: compareItems } = useComparison();

  if (isPending) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error loading items.</div>;
  }

  const { left_item, right_item } = data;

  const handleComparison = async (winnerId: string, loserId: string) => {
    await compareItems({ winner_id: winnerId, loser_id: loserId });
    await refetch();
  };

  return (
    <Box>
        <Typography variant="h4" sx={{ textAlign: 'center', mb: 4 }}>Ribbit ribbit</Typography>
        <Box sx={{ display: 'flex', gap: 4, flexDirection: { xs: 'column', md: 'row' } }}>
            <Contender imageUrl={left_item.image_url} name={left_item.name} onClick={() => handleComparison(left_item.id, right_item.id)} />
            <Typography variant="h3" sx={{ alignSelf: 'center' }}>VS</Typography>
            <Contender imageUrl={right_item.image_url} name={right_item.name} onClick={() => handleComparison(right_item.id, left_item.id)} />
        </Box>
    </Box>
  );
}

export default Comparison;
