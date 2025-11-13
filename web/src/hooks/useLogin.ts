import { useMutation } from "@tanstack/react-query";
import { apiFetch } from "../api/client";

export interface LoginResponse {
    token: string;
}

function useLogin() {
    return useMutation({
        mutationKey: ['login'],
        mutationFn: async (data: { username: string; password: string; }) => {
            const res = await apiFetch<LoginResponse>('/login', {
                method: 'POST',
                body: JSON.stringify(data),
            });
            return res;
        }
    })
}

export { useLogin };