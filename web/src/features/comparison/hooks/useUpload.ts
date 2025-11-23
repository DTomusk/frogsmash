import { useMutation } from "@tanstack/react-query";
import { apiFetch } from "@/shared";
import type { UploadResponse } from "../dtos/uploadResponse";

function useUpload() {
    return useMutation({
        mutationKey: ['upload'],
        mutationFn: async (file: File) => {
            const formData = new FormData();
            formData.append('image', file);
            const res = await apiFetch<UploadResponse>('/upload', {
                method: 'POST',
                body: formData
            });
            return res;
        }
    });
}

export { useUpload };