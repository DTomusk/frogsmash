import { useMutation } from "@tanstack/react-query";
import { apiFetch } from "../api/client";

function useResendVerification() {
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

function useVerifyCode() {
    return useMutation({
        mutationKey: ['verifyCode'],
        mutationFn: async (code: string) => {
            const res = await apiFetch<void>('/verify', {
                method: 'POST',
                body: JSON.stringify({ code }),
            });
            return res;
        }
    })
}

export { useResendVerification, useVerifyCode };