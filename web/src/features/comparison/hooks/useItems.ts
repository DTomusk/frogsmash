import { useQuery } from "@tanstack/react-query";
import type { ItemsResponse } from "../dtos/itemsResponse";
import { apiFetch } from "@/shared";

export function useItems() {
    return useQuery({
        queryKey: ['comparisonItems'],
        queryFn: async () => {
            const response = await apiFetch<ItemsResponse>('/items');
            return response.items;
        },
        refetchOnWindowFocus: false,  
        refetchOnReconnect: false,     
        refetchOnMount: false,         
        staleTime: Infinity,           
    });
}