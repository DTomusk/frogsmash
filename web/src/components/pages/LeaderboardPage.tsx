import { List, Pagination, Typography } from "@mui/material";
import { useLeaderboard } from "../../hooks/useLeaderboard";
import LoadingSpinner from "../LoadingSpinner";
import LeaderboardEntry from "../LeaderboardEntry";
import { useState } from "react";
import LeaderboardDetailModal from "../LeaderboardDetailModal";
import type { LeaderboardItem } from "../../models/items";

function LeaderboardPage() {
    const [page, setPage] = useState(1);
    const { data: leaderboardData, isLoading } = useLeaderboard(page, 10);

    const [open, setOpen] = useState(false);
    const [selectedItem, setSelectedItem] = useState<LeaderboardItem | null>(null);

  return <>
        <Typography variant="h3" sx={{mt: 4, mb: 2}}>Leaderboard</Typography>
        <Typography variant="subtitle1" sx={{mb: 2}}>🐸See who the strongest contenders are!🐸</Typography>
        {isLoading && <LoadingSpinner />}
        {!isLoading && (<>
            <List sx={{ width: { xs: '100%', sm: '80%', md: '60%' }, mt: 4, mb: 4 }}>
            {leaderboardData?.data?.map((item) => (
                <LeaderboardEntry key={item.id} item={item} onClick={() => {setSelectedItem(item); setOpen(true);}} />
            ))}
            </List>
            {leaderboardData!.total_pages > 1 && <Pagination color="primary" count={leaderboardData?.total_pages} page={page} onChange={(_, value) => setPage(value)} />}
        </>)}
        {selectedItem && <LeaderboardDetailModal open={open} setOpen={setOpen} item={selectedItem!} />} 
    </>;
}

export default LeaderboardPage;