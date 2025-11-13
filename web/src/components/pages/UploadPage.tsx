import { Alert, Button, Paper, Snackbar, styled, Typography } from "@mui/material";
import CloudUploadIcon from '@mui/icons-material/CloudUpload';
import { useUpload } from "../../hooks/useUpload";
import { useState } from "react";

const VisuallyHiddenInput = styled('input')({
  clip: 'rect(0 0 0 0)',
  clipPath: 'inset(50%)',
  height: 1,
  overflow: 'hidden',
  position: 'absolute',
  bottom: 0,
  left: 0,
  whiteSpace: 'nowrap',
  width: 1,
});

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
    <Paper
        component="form"
        onSubmit={() => {}}
        sx={{
            display: "flex",
            flexDirection: "column",
            gap: 2,
            alignItems: "center",
            m: 4,
            p: 4,
            borderRadius: 2,
            maxWidth: 600,
        }}
    >
        <Typography variant="h3" sx={{ mb: 2}}>Submit a contender</Typography>
        <Typography variant="subtitle1" sx={{mb: 2}}>Does your champion have what it takes to take on the mighty frogs?🐸 Submit an image of anything of your choosing for the chance to have them appear alongside the frogs on the battlefield and let the people decide whether they triumph or fall.</Typography>
        <Button
            component="label"
            role={undefined}
            variant="contained"
            tabIndex={-1}
            startIcon={<CloudUploadIcon />}
            sx={{ mb: 4}}
            loading={isPending}
            disabled={isPending || uploadDisabled}
        >
            Upload your image
            <VisuallyHiddenInput
                type="file"
                onChange={(event) => handleFileChange(event)}
                accept=".png, .jpg, .jpeg"
            />
        </Button>
    </Paper>
    <Snackbar open={openSuccess} autoHideDuration={6000} onClose={() => setOpenSuccess(false)}>
        <Alert onClose={() => setOpenSuccess(false)} severity="success" variant="outlined" sx={{ width: '100%', bgcolor: 'background.paper' }}>
            {data?.message}
        </Alert>
    </Snackbar>
    <Snackbar open={openError} autoHideDuration={6000} onClose={() => setOpenError(false)}>
        <Alert onClose={() => setOpenError(false)} severity="error" variant="outlined" sx={{ width: '100%', bgcolor: 'background.paper' }}>
            {errorMessage || "An error occurred during upload."}
        </Alert>
    </Snackbar>
    </>
  );
}
export default UploadPage;