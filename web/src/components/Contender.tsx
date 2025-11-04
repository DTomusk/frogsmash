import { Paper, Typography } from "@mui/material";
import { Image } from 'mui-image'

interface ContenderProps {
    imageUrl: string;
    name: string;
}

function Contender({ imageUrl, name }: ContenderProps) {
    return (
        <Paper elevation={3} sx={{ padding: 2, textAlign: 'center', display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
            <Image src={imageUrl} alt={name} />
            <Paper>
                <Typography variant="h6">{name}</Typography>
            </Paper>
        </Paper>
    );
}
export default Contender;