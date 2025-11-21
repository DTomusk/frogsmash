import { Box, Button, Card, CardActions, CardMedia, IconButton, Modal, Tooltip, Typography } from "@mui/material";
import type { LeaderboardItem } from "../../models/items";
import { format } from 'date-fns';
import InfoIcon from '@mui/icons-material/Info';
import { Image } from 'mui-image'

interface LeaderboardDetailModalProps {
    open: boolean;
    setOpen: (open: boolean) => void;
    item: LeaderboardItem;
}

function LeaderboardDetailModal({ open, setOpen, item }: LeaderboardDetailModalProps) {
    return (
        <Modal
        open={open}
        onClose={() => setOpen(false)}
        aria-labelledby="card-modal-title"
        aria-describedby="card-modal-description"
      >
        <Box
          sx={{
            position: "absolute" as const,
            top: "50%",
            left: "50%",
            transform: "translate(-50%, -50%)",
            outline: "none",
          }}
        >
          <Card sx={{ width: 400 }}>
            <CardMedia>
                <Image
                    src={item.image_url}
                    alt={item.name}
                    showLoading 
                    width="100%"
                    height={300}
                    fit="cover"
                    />
            </CardMedia>
            <Box sx={{display: 'flex', justifyContent: 'space-between', alignItems: 'center'}}>
            <Typography variant="h6" gutterBottom sx={{ px: 2, pt: 2 }}>
                {item.name}
            </Typography>
            <Tooltip sx={{mx: 2}} title={`Licensing info: ${item.license}`} arrow>
                <IconButton size="small">
                    <InfoIcon color="action" />
                </IconButton>
            </Tooltip>
            </Box>
            <Typography variant="body1" sx={{ px: 2, pb: 2 }}>
                Rank: #{item.rank} <br/>
                Active since: {format(new Date(item.created_at), 'yyyy-MM-dd')}<br/>
                Score: {item.score}
            </Typography>
            <CardActions sx={{ justifyContent: "flex-end" }}>
              <Button onClick={() => setOpen(false)}>Close</Button>
            </CardActions>
          </Card>
        </Box>
      </Modal>
    )
};

export default LeaderboardDetailModal;
