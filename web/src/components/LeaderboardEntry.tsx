import { Avatar, Box, ListItem, ListItemAvatar, ListItemText, Paper, Typography } from "@mui/material";
import type { LeaderboardItem } from "../models/items";

interface LeaderboardEntryProps {
    item: LeaderboardItem;
    onClick?: () => void;
}

function LeaderboardEntry({ item, onClick }: LeaderboardEntryProps) {
    return (
        <Paper 
            sx={{ marginBottom: 2, 
                padding: 1, 
                width: '100%', 
                transition: "transform 0.3s ease",
                transformOrigin: "center center",
                "&:hover": {
                    transform: "scale(1.05)",
                    cursor: "pointer",
                    backgroundColor: 'action.hover',
                },
            }} 
            onClick={onClick}>
        <ListItem sx={{ display: 'flex', justifyContent: 'space-between' }}>
            <Box sx={{ display: 'flex', alignItems: 'center', width: '40%' }}>
            <ListItemText primary={
              <Typography variant="h6">
                #{item.rank}
                {item.rank === 1 && " 🏆"}
              </Typography>
            } sx={{width: '50%'}}/>
            <ListItemAvatar>
                <Avatar src={item.image_url} alt={item.name} />
            </ListItemAvatar>
            
            <ListItemText sx={{ width: '100%' }}
                primary={item.name}
            />
            </Box>
            <ListItemText primary={item.score} sx={{ flex: 1, textAlign: 'right' }} />
        </ListItem>
        </Paper>
    );
}

export default LeaderboardEntry;
