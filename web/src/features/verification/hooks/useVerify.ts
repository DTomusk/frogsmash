import { useMutation } from "@tanstack/react-query";
import { apiFetch } from "@/shared";
import type { VerificationResponse } from "../dtos/verificationResponse";

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
            const res = await apiFetch<VerificationResponse>('/verify', {
                method: 'POST',
                body: JSON.stringify({ code }),
            });
            return res;
        }
    })
}

export { useResendVerification, useVerifyCode };