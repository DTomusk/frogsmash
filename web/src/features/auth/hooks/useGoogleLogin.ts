import { useMutation } from "@tanstack/react-query";
import { apiFetch } from "@/shared";
import type { LoginResponse } from "../dtos/loginResponse";
import type { GoogleLoginRequest } from "../dtos/googleLoginRequest"

function useGoogleLogin() {
    return useMutation({
        mutationKey: ['googleLogin'],
        mutationFn: async (data: GoogleLoginRequest) => {
            const res = await apiFetch<LoginResponse>('/auth/google-login', {
                method: 'POST',
                body: JSON.stringify(data),
            });
            return res;
        }
    })
}

export { useGoogleLogin };