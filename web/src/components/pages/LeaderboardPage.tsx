import { List, Pagination, Typography } from "@mui/material";
import { useLeaderboard } from "../../hooks/useLeaderboard";
import LoadingSpinner from "../LoadingSpinner";
import LeaderboardEntry from "../LeaderboardEntry";
import { useState } from "react";

function LeaderboardPage() {
    const [page, setPage] = useState(1);
    const { data: leaderboardData, isLoading } = useLeaderboard(page, 10);

  return <>
        <Typography variant="h3" sx={{mt: 4, mb: 2}}>Leaderboard</Typography>
        <Typography variant="subtitle1" sx={{mb: 2}}>🐸See who the strongest contenders are!🐸</Typography>
        {isLoading && <LoadingSpinner />}
        {!isLoading && (<>
            <List sx={{ width: { xs: '100%', sm: '80%', md: '60%' }, mt: 4, mb: 4 }}>
            {leaderboardData?.data?.map((item) => (
                <LeaderboardEntry key={item.id} item={item} />
            ))}
            </List>
            {leaderboardData!.total_pages > 1 && <Pagination color="primary" count={leaderboardData?.total_pages} page={page} onChange={(_, value) => setPage(value)} />}
        </>)}
    </>;
}

export default LeaderboardPage;