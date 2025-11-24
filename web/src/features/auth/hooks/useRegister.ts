import { useMutation } from "@tanstack/react-query";
import { apiFetch } from "@/shared";
import type { ApiResponse } from "@/shared";
import type { RegistrationRequest } from "../dtos/registrationRequest";

function useRegister() {
    return useMutation({
        mutationKey: ['register'],
        mutationFn: async (data: RegistrationRequest) => {
            const res = await apiFetch<ApiResponse>('/auth/register', {
                method: 'POST',
                body: JSON.stringify(data),
            });
            return res;
        },
    })
}

export { useRegister };