import { useMutation } from "@tanstack/react-query";

function useLogout() {
    return useMutation({
        mutationKey: ['logout'],
        mutationFn: async () => {
            await fetch('/auth/logout', {
                method: 'POST',
                credentials: 'include',
            });
        },
    })
}

export { useLogout };