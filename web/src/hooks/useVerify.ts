import { useMutation } from "@tanstack/react-query";
import { apiFetch } from "../api/client";

function useVerify() {
    return useMutation({
        mutationKey: ['resendVerification'],
        mutationFn: async () => {
            const res = await apiFetch<void>('/resend-verification', {
                method: 'POST',
            });
            return res;
        }
    })
}

export { useVerify };