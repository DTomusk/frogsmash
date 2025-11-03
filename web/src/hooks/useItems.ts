import { useQuery } from "@tanstack/react-query";
import { apiFetch } from "../api/client";
import type { ItemsResponse } from "../models/items";

export function useItems() {
    return useQuery({
        queryKey: ['comparisonItems'],
        queryFn: async () => {
            const response = await apiFetch<ItemsResponse>('/items');
            return response.items;
        },
    });
}