import { Button, Paper, styled, Typography } from "@mui/material";
import CloudUploadIcon from '@mui/icons-material/CloudUpload';

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
  return (
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
        >
            Upload your image
            <VisuallyHiddenInput
                type="file"
                onChange={(event) => console.log(event.target.files)}
                accept=".png, .jpg, .jpeg"
            />
        </Button>
    </Paper>
  );
}
export default UploadPage;