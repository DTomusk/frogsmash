import { useMutation } from "@tanstack/react-query";
import { apiFetch } from "@/shared";

export interface LoginResponse {
    jwt: string;
    user: {
        id: string;
        email: string;
        isVerified: boolean;
    };
}

function useLogin() {
    return useMutation({
        mutationKey: ['login'],
        mutationFn: async (data: { email: string; password: string; }) => {
            const res = await apiFetch<LoginResponse>('/login', {
                method: 'POST',
                body: JSON.stringify(data),
            });
            return res;
        }
    })
}

export { useLogin };