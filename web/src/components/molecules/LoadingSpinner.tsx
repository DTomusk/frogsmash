import { Box, keyframes, Typography } from "@mui/material";
import { useEffect, useState } from "react";

const spin = keyframes`
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
`;

function LoadingSpinner() {
    // State of dots in loading text
    const [dots, setDots] = useState("");

    useEffect(() => {
        // Update dots every 500ms
        const interval = setInterval(() => {
            setDots((prev) => (prev.length < 3 ? prev + "." : ""));
        }, 500);
        return () => clearInterval(interval);
    }, []);

    return <Box 
        display="flex" 
        flexDirection="column" 
        alignItems="center"
        sx={{ width: '8ch'}}>
        <Box sx={{ display: 'inline-block', animation: `${spin} 1s linear infinite`, fontSize: 64 }} >🐸</Box>
        <Typography variant="h6" sx={{ textAlign: 'center', mt: 2, width: '100%' }}>Loading{dots}</Typography>
    </Box>;
}
export default LoadingSpinner;