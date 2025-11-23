import { useMutation } from "@tanstack/react-query";
import { apiFetch } from "@/shared";
import type { LoginResponse } from "../dtos/loginResponse";
import type { LoginRequest } from "../dtos/loginRequest";

function useLogin() {
    return useMutation({
        mutationKey: ['login'],
        mutationFn: async (data: LoginRequest) => {
            const res = await apiFetch<LoginResponse>('/login', {
                method: 'POST',
                body: JSON.stringify(data),
            });
            return res;
        }
    })
}

export { useLogin };