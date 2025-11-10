import { Pagination, Typography } from "@mui/material";
import { useLeaderboard } from "../../hooks/useLeaderboard";
import LoadingSpinner from "../LoadingSpinner";
import LeaderboardEntry from "../LeaderboardEntry";
import { useState } from "react";

function LeaderboardPage() {
    const [page, setPage] = useState(1);
    const { data, isLoading } = useLeaderboard(page, 10);

  return <>
    <Typography variant="h3" sx={{mt: 4}}>Leaderboard Page</Typography>
    {isLoading && <LoadingSpinner />}
    {!isLoading && (<>
    <ul>
      {data?.data?.map((item) => (
        <LeaderboardEntry key={item.id} item={item} />
      ))}
    </ul>
    <Pagination color="primary" count={data?.total_pages} page={page} onChange={(_, value) => setPage(value)} /></>)}
  </>;
}

export default LeaderboardPage;