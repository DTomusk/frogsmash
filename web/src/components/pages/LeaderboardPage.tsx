import { Typography } from "@mui/material";
import { useLeaderboard } from "../../hooks/useLeaderboard";
import LoadingSpinner from "../LoadingSpinner";
import LeaderboardEntry from "../LeaderboardEntry";

function LeaderboardPage() {
    const { data, isLoading } = useLeaderboard(1, 10);

  return <>
    <Typography variant="h3" sx={{mt: 4}}>Leaderboard Page</Typography>
    {isLoading && <LoadingSpinner />}
    {!isLoading && (<ul>
      {data?.map((item) => (
        <LeaderboardEntry key={item.id} item={item} />
      ))}
    </ul>)}
    </>;
}

export default LeaderboardPage;