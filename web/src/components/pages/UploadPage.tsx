import { Typography } from "@mui/material";
import { useUpload } from "../../hooks/useUpload";
import { useState } from "react";
import FileUploadButton from "../molecules/FileUploadButton";
import AlertSnackbar from "../molecules/AlertSnackbar";
import FormWrapper from "../atoms/FormWrapper";

function UploadPage() {
    const { mutate: upload, data, isPending } = useUpload();
    const [uploadDisabled, setUploadDisabled] = useState(false);
    const [openSuccess, setOpenSuccess] = useState(false);
    const [errorMessage, setErrorMessage] = useState("");
    const [openError, setOpenError] = useState(false);

    const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const file = event.target.files?.[0];
        if (!file) {
            return;
        }
        if (!['image/png', 'image/jpeg', "image/jpg"].includes(file.type)) {
            setErrorMessage("Invalid file type. Please upload a PNG or JPEG image.");
            setOpenError(true);
            return;
        }
        if (file.size > 5 * 1024 * 1024) {
            setErrorMessage("File size exceeds the 5MB limit.");
            setOpenError(true);
            return;
        }
        upload(file, {
            onSuccess: () => {
                setUploadDisabled(true);
                setOpenSuccess(true);
            },
            onError: (err: any) => {
                setErrorMessage(err.message || "Upload failed");
                setOpenError(true);
            },
        });
    };

  return (
    <>
    <FormWrapper onSubmit={(e) => e.preventDefault()}>
        <Typography variant="h3" sx={{ mb: 2}}>Submit a contender</Typography>
        <Typography variant="subtitle1" sx={{mb: 2}}>Does your champion have what it takes to take on the mighty frogs?🐸 Submit an image of anything of your choosing for the chance to have them appear alongside the frogs on the battlefield and let the people decide whether they triumph or fall.</Typography>
        <FileUploadButton
            onChange={handleFileChange}
            isPending={isPending}
            disabled={uploadDisabled}
        />
    </FormWrapper>
    <AlertSnackbar
        open={openSuccess}
        onClose={() => setOpenSuccess(false)}
        severity="success"
        message={data?.message || "Upload successful!"}
    />
    <AlertSnackbar
        open={openError}
        onClose={() => setOpenError(false)}
        severity="error"
        message={errorMessage || "An error occurred during upload."}
    />
    </>
  );
}
export default UploadPage;