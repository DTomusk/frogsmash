import { useMutation } from "@tanstack/react-query";
import { apiFetch } from "../api/client";

function useUpload() {
    return useMutation({
        mutationKey: ['upload'],
        mutationFn: async (file: File) => {
            const formData = new FormData();
            formData.append('image', file);
            const res = await apiFetch<void>('/upload', {
                method: 'POST',
                body: formData
            });
            return res;
        }
    });
}

export { useUpload };