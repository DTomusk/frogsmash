import { useMutation } from "@tanstack/react-query";
import { apiFetch } from "@/shared";
import type { UploadResponse } from "../dtos/uploadResponse";
import type { LatestSubmissionResponse } from "../dtos/latestSubmissionResponse";

function useUpload() {
    return useMutation({
        mutationKey: ['upload'],
        mutationFn: async (file: File) => {
            const formData = new FormData();
            formData.append('image', file);
            const res = await apiFetch<UploadResponse>('/comparison/submit-contender', {
                method: 'POST',
                body: formData
            });
            return res;
        }
    });
}

function useLatestUploadTime() {
    return useMutation({
        mutationKey: ['latestUploadTime'],
        mutationFn: async () => {
            const res = await apiFetch<LatestSubmissionResponse>('/comparison/latest-submission', {
                method: 'GET'
            });
            return res;
        }
    });
}

export { useUpload, useLatestUploadTime };