import { List, ListItem, ListItemText } from "@mui/material";

interface BulletListProps {
    items: string[];
}

export default function BulletList({ items }: BulletListProps) {
    return (
        <List sx={{ listStyleType: 'disc', pl: 4, width: {sx: "100%", md: "50%"} }}>
            {items.map((item, index) => (
                <ListItem key={index} sx={{ display: 'list-item' }}>
                    <ListItemText>{item}</ListItemText>
                </ListItem>
            ))}
        </List>
    );
}