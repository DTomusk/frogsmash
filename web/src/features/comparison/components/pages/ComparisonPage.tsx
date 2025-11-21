import { Box, Button, Typography } from "@mui/material";
import { useItems } from "../../hooks/useItems"
import Contender from "../organisms/Contender";
import { useComparison } from "../../hooks/useComparison";
import LoadingSpinner from "../../../../shared/components/molecules/LoadingSpinner";
import ErrorMessage from "../../../../shared/components/atoms/ErrorMessage";

function Comparison() {
  const { isPending, error, data, refetch } = useItems();
  const { mutate: compareItems } = useComparison();

  if (isPending) {
    return <LoadingSpinner />;
  }

  if (error) {
    return (
      <Box sx={{ display: 'flex', flexDirection: 'column', alignItems: 'center', mt: 4 }}>
        <ErrorMessage message={'Error loading contenders'} />
        <Button variant="contained" onClick={() => refetch()} sx={{ mt: 2, textAlign: 'center' }}>
          Retry
        </Button>
      </Box>
    );
  }

  const { left_item, right_item } = data;

  const handleComparison = async (winnerId: string, loserId: string) => {
    await compareItems({ winner_id: winnerId, loser_id: loserId });
    await refetch();
  };

  return (
    <Box>
        <Typography variant="h4" sx={{ textAlign: 'center', mb: 4 }}>Ribbit ribbit? Ribbit.</Typography>
        <Box sx={{ display: 'flex', gap: 4, justifyContent: 'space-between', flexDirection: { xs: 'column', sm: 'row' } }}>
            <Contender imageUrl={left_item.image_url} name={left_item.name} onClick={() => handleComparison(left_item.id, right_item.id)} variant="left" />
            <Typography variant="h3" sx={{ alignSelf: 'center' }}>VS</Typography>
            <Contender imageUrl={right_item.image_url} name={right_item.name} onClick={() => handleComparison(right_item.id, left_item.id)} variant="right" />
        </Box>
    </Box>
  );
}

export default Comparison;
