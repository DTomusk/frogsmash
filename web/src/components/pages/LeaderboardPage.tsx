import { useLeaderboard } from "../../hooks/useLeaderboard";
import LoadingSpinner from "../LoadingSpinner";

function LeaderboardPage() {
    const { data, isLoading } = useLeaderboard(1, 10);

    if (isLoading) {
        return <LoadingSpinner />;
    }
  return <div>
    <div>Leaderboard Page</div>
    <ul>
      {data?.map((item) => (
        <li key={item.id}>
          <span>Rank: {item.rank} </span>
          <span>Name: {item.name} </span>
          <span>Score: {item.score} </span>
          <img src={item.image_url} alt={item.name} width={50} height={50} />
        </li>
      ))}
    </ul>
    </div>;
}

export default LeaderboardPage;