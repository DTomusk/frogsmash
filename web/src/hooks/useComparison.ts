import { useMutation } from "@tanstack/react-query";
import { apiFetch } from "../api/client";

function useComparison() {
    return useMutation({
        mutationKey: ['comparison'],
        mutationFn: async ({ winner_id, loser_id }: { winner_id: string; loser_id: string; }) => {
            await apiFetch<void>('/comparison', {
                method: 'POST',
                body: JSON.stringify({ winner_id, loser_id }),
            });
            return;
        },
    });
}

export { useComparison };