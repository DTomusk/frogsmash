import { Avatar, ListItem, ListItemAvatar, ListItemText, Paper } from "@mui/material";
import type { LeaderboardItem } from "../models/items";

function LeaderboardEntry({ item }: { item: LeaderboardItem }) {
    return (
        <Paper sx={{ marginBottom: 2, padding: 1 }}>
        <ListItem>
            <ListItemAvatar>
                <Avatar src={item.image_url} alt={item.name} />
            </ListItemAvatar>
            <ListItemText
                primary={`Rank: ${item.rank} - ${item.name}`}
                secondary={`Score: ${item.score}`}
            />
        </ListItem>
        </Paper>
    );
}

export default LeaderboardEntry;
