import { useMutation } from "@tanstack/react-query";
import { apiFetch } from "@/shared";

function useLogout() {
    return useMutation({
        mutationKey: ['logout'],
        mutationFn: async () => {
            await apiFetch('/auth/logout', {
                method: 'POST',
                credentials: 'include',
            });
        },
    })
}

export { useLogout };