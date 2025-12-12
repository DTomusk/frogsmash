import { Box, Paper, Typography } from "@mui/material";
import { Image } from 'mui-image'

interface ContenderProps {
    imageUrl: string;
    name: string;
    onClick: () => void;
    variant: 'left' | 'right';
}

function Contender({ imageUrl, name, onClick, variant }: ContenderProps) {
    return (
    <Box display="flex" flexDirection="column" alignItems="center" maxHeight="70%" position='relative' onClick={onClick}
    sx={{
        transition: "transform 0.3s ease",
        transformOrigin: "center center",
        "&:hover": {
          transform: "scale(1.05)",
          cursor: "pointer",
        },
      }}>
        <Paper elevation={3} sx={{borderRadius: 2, overflow: 'hidden', width: {xs: 200, sm: 250, md: 300, lg: 350}, height: {xs: 200, sm: 250, md: 300, lg: 350}}}>
            <Image src={imageUrl} alt={name} />
        </Paper>
        <Paper elevation={10} 
        sx={{ 
            backgroundColor: variant === 'left' 
            ? 'primary.main' 
            : 'secondary.main', 
            width: 'fit-content', 
            minWidth: '80%', 
            px: 2, 
            py: {xs: 1, sm: 2},
            display: 'flex', 
            justifyContent: 'center', 
            mt: -5, 
            zIndex: 1, 
            borderRadius: 4 }}>
            <Typography variant="h4" color="white">{name}</Typography>
        </Paper>
    </Box>
    );
}
export default Contender;