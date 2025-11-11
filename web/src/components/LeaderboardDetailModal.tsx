import { Box, Button, Card, CardActionArea, CardActions, CardMedia, Modal, Typography } from "@mui/material";
import type { LeaderboardItem } from "../models/items";

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
            <CardActionArea>
                <CardMedia component='img' image={item.image_url} alt={item.name} />
                <Typography id="card-modal-title" variant="h6" gutterBottom sx={{ p: 2 }}>
                {item.name}
              </Typography>
            </CardActionArea>
            <CardActions sx={{ justifyContent: "flex-end" }}>
              <Button onClick={() => setOpen(false)}>Close</Button>
            </CardActions>
          </Card>
        </Box>
      </Modal>
    )
};

export default LeaderboardDetailModal;
