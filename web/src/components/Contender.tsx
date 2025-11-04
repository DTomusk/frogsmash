import { Paper, Typography } from "@mui/material";
import { Image } from 'mui-image'

interface ContenderProps {
    imageUrl: string;
    name: string;
}

function Contender({ imageUrl, name }: ContenderProps) {
    return (
        <Paper elevation={3} sx={{ padding: 2, textAlign: 'center', display: 'flex', flexDirection: 'column', alignItems: 'center', height: 300, width: 200 }}>
            <Image src={imageUrl} alt={name} height='70%' />
            <Paper elevation={10} sx={{ backgroundColor: 'primary.main', width: '80%', py: 2 }}>
                <Typography variant="h4" color="white">{name}</Typography>
            </Paper>
        </Paper>
    );
}
export default Contender;