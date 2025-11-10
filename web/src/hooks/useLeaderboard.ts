import { useQuery } from "@tanstack/react-query";
import { apiFetch } from "../api/client";
import type { LeaderboardResponse } from "../models/items";

export function useLeaderboard(page: number, limit: number) {
    return useQuery({
        queryKey: ['leaderboardItems', page, limit],
        queryFn: async () => {
            const response = await apiFetch<LeaderboardResponse>(`/leaderboard?page=${page}&limit=${limit}`);
            return response;
        }
    });
}