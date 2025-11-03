import { useQuery } from "@tanstack/react-query";
import { apiFetch } from "../api/client";

export function useItems() {
    // Hook logic to fetch and return items would go here
    return useQuery({
        queryKey: ['comparisonItems'],
        queryFn: async () => {
            const response = await apiFetch<ItemsResponse>('/items');
            return response.items;
        },
    });
}

// TODO: move to models folder
export interface Item {
  id: string;
  name: string;
  image_url: string;
}

export interface ItemsResponse {
  items: {
    left_item: Item;
    right_item: Item;
  };
}