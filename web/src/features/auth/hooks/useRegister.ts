import { useMutation } from "@tanstack/react-query";
import { apiFetch } from "@/shared";

function useRegister() {
    return useMutation({
        mutationKey: ['register'],
        mutationFn: async (data: { email: string; password: string; }) => {
            const res = await apiFetch<{ message: string }>('/register', {
                method: 'POST',
                body: JSON.stringify(data),
            });
            return res;
        },
    })
}

export { useRegister };