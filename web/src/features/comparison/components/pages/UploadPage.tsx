import { Typography } from "@mui/material";
import { useLatestUploadTime, useUpload } from "../../hooks/useUpload";
import { useEffect, useState } from "react";
import { useSnackbar } from "@/app/providers";
import { FormWrapper } from "@/shared";
import FileUploadButton from "../molecules/FileUploadButton";
import TomorrowCountdown from "../molecules/TomorrowCountdown";

function UploadPage() {
    const { mutate: upload, isPending } = useUpload();
    const [uploadDisabled, setUploadDisabled] = useState(false);
    const { showSnackbar } = useSnackbar();
    const { mutate: getLatestUploadTime } = useLatestUploadTime();
    const [uploadedToday, setUploadedToday] = useState(false);

    useEffect(() => {
        getLatestUploadTime(undefined, {
            onSuccess: (res) => {
                if (!res.uploaded_at) {
                    setUploadDisabled(false);
                    return;
                }

                const lastUpload = new Date(res.uploaded_at);
                const now = new Date();

                const lastY = lastUpload.getUTCFullYear();
                const lastM = lastUpload.getUTCMonth();
                const lastD = lastUpload.getUTCDate();


                const nowY = now.getUTCFullYear();
                const nowM = now.getUTCMonth();
                const nowD = now.getUTCDate();

                const didUploadTodayUTC =
                    lastY === nowY && lastM === nowM && lastD === nowD;

                setUploadedToday(didUploadTodayUTC);
            },
            onError: () => {
                setUploadedToday(false);
            }
        }
    );
    }, [getLatestUploadTime]);

    const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const file = event.target.files?.[0];
        if (!file) {
            return;
        }
        if (!['image/png', 'image/jpeg', "image/jpg"].includes(file.type)) {
            showSnackbar({ message: "Invalid file type. Please upload a PNG or JPEG image.", severity: "error" });
            return;
        }
        if (file.size > 5 * 1024 * 1024) {
            showSnackbar({ message: "File size exceeds the 5MB limit.", severity: "error" });
            return;
        }
        upload(file, {
            onSuccess: () => {
                setUploadDisabled(true);
                showSnackbar({ message: "Upload successful! Your contender will appear soon.", severity: "success" });
                setUploadedToday(true);
            },
            onError: (err: any) => {
                showSnackbar({ message: err.message || "Upload failed", severity: "error" });
            },
        });
    };

  return (
    <FormWrapper onSubmit={(e) => e.preventDefault()}>
        <Typography variant="h3" sx={{ mb: 2}}>Submit a contender</Typography>
        {uploadedToday? <TomorrowCountdown onFinish={() => setUploadedToday(false)} />:
        <>
        <Typography variant="subtitle1" sx={{mb: 2}}>Does your champion have what it takes to take on the mighty frogs?🐸 Submit an image of anything of your choosing for the chance to have them appear alongside the frogs on the battlefield and let the people decide whether they triumph or fall.</Typography>
        <FileUploadButton
            onChange={handleFileChange}
            isPending={isPending}
            disabled={uploadDisabled}
        />
        </>}
    </FormWrapper>
  );
}
export default UploadPage;