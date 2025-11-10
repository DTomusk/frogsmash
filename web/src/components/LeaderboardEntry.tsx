import { Avatar, Box, ListItem, ListItemAvatar, ListItemText, Paper, Typography } from "@mui/material";
import type { LeaderboardItem } from "../models/items";

function LeaderboardEntry({ item }: { item: LeaderboardItem }) {
    return (
        <Paper sx={{ marginBottom: 2, padding: 1, width: '100%' }}>
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
