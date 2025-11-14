import { Typography } from "@mui/material";
import { useUpload } from "../../hooks/useUpload";
import { useState } from "react";
import FileUploadButton from "../molecules/FileUploadButton";
import FormWrapper from "../atoms/FormWrapper";
import { useSnackbar } from "../../contexts/SnackbarContext";

function UploadPage() {
    const { mutate: upload, isPending } = useUpload();
    const [uploadDisabled, setUploadDisabled] = useState(false);
    const { showSnackbar } = useSnackbar();

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
            },
            onError: (err: any) => {
                showSnackbar({ message: err.message || "Upload failed", severity: "error" });
            },
        });
    };

  return (
    <FormWrapper onSubmit={(e) => e.preventDefault()}>
        <Typography variant="h3" sx={{ mb: 2}}>Submit a contender</Typography>
        <Typography variant="subtitle1" sx={{mb: 2}}>Does your champion have what it takes to take on the mighty frogs?🐸 Submit an image of anything of your choosing for the chance to have them appear alongside the frogs on the battlefield and let the people decide whether they triumph or fall.</Typography>
        <FileUploadButton
            onChange={handleFileChange}
            isPending={isPending}
            disabled={uploadDisabled}
        />
    </FormWrapper>
  );
}
export default UploadPage;