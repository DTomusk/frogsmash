import { useQuery } from "@tanstack/react-query";
import type { User } from "../models/user";
import { apiFetch } from "@/shared";

function useCurrentUser() {
    return useQuery({
        queryKey: ['currentUser'],
        queryFn: async () => {
            const res = await apiFetch<User>('/auth/me', {
                method: 'GET',
            });
            return res;
        },
        retry: false,
    });
}

export { useCurrentUser };