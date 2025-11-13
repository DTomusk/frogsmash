import { Button, styled } from "@mui/material";
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

interface FileUploadButtonProps {
    onChange: (event: React.ChangeEvent<HTMLInputElement>) => void;
    isPending: boolean;
    disabled: boolean;
}

function FileUploadButton({ onChange, isPending, disabled }: FileUploadButtonProps) {
    return (
        <Button
            component="label"
            role={undefined}
            variant="contained"
            tabIndex={-1}
            startIcon={<CloudUploadIcon />}
            sx={{ mb: 4}}
            loading={isPending}
            disabled={isPending || disabled}
        >
            Upload your image
            <VisuallyHiddenInput
                type="file"
                onChange={onChange}
                accept=".png, .jpg, .jpeg"
            />
        </Button>
    )
}

export default FileUploadButton;