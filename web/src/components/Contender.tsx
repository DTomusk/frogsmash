import { Box, Paper, Typography } from "@mui/material";
import { Image } from 'mui-image'

interface ContenderProps {
    imageUrl: string;
    name: string;
    onClick: () => void;
}

function Contender({ imageUrl, name, onClick }: ContenderProps) {
    return (
    <Box display="flex" flexDirection="column" alignItems="center" maxHeight="70%" position='relative'
    sx={{
        transition: "transform 0.3s ease",
        transformOrigin: "center center",
        "&:hover": {
          transform: "scale(1.05)",
          cursor: "pointer",
        },
      }}>
        <Paper elevation={3}>
            <Image src={imageUrl} alt={name} />
        </Paper>
        <Paper elevation={10} sx={{ backgroundColor: 'primary.main', width: '80%', py: 2, display: 'flex', justifyContent: 'center', mt: -5, zIndex: 1, borderRadius: 4 }}>
            <Typography variant="h4" color="white">{name}</Typography>
        </Paper>
    </Box>
    );
}
export default Contender;