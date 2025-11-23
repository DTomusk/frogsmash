import { useMutation } from "@tanstack/react-query";
import { apiFetch } from "@/shared";
import type { ComparisonRequest } from "../dtos/comparisonRequest";

function useComparison() {
    return useMutation({
        mutationKey: ['comparison'],
        mutationFn: async ({ winner_id, loser_id }: ComparisonRequest) => {
            await apiFetch<void>('/compare', {
                method: 'POST',
                body: JSON.stringify({ winner_id, loser_id }),
            });
            return;
        },
    });
}

export { useComparison };