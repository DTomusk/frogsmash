import { Paper, Typography } from "@mui/material";
import { useEffect, useState } from "react";

export default function TomorrowCountdown({ onFinish }: { onFinish?: () => void }) {
    const [timeLeft, setTimeLeft] = useState("");

    const updateCountdown = () => {
        const now = new Date();

        // Next UTC midnight
        const nextUTC = new Date(Date.UTC(
            now.getUTCFullYear(),
            now.getUTCMonth(),
            now.getUTCDate() + 1, // tomorrow
            0, 0, 0, 0
        ));

        const diff = nextUTC.getTime() - now.getTime();

        if (diff <= 0) {
            setTimeLeft("00:00:00");
            if (onFinish) {
                onFinish();
            }
            return;
        }

        const hours = Math.floor(diff / (1000 * 60 * 60));
        const minutes = Math.floor((diff % (1000 * 60 * 60)) / (1000 * 60));
        const seconds = Math.floor((diff % (1000 * 60)) / 1000);

        const pad = (n: number) => String(n).padStart(2, "0");

        setTimeLeft(`${pad(hours)}:${pad(minutes)}:${pad(seconds)}`);
    };

    useEffect(() => {
        updateCountdown(); // initial
        const interval = setInterval(updateCountdown, 1000);
        return () => clearInterval(interval);
    }, []);

    return (<>
        <Typography variant="body1" gutterBottom>
            You've already made a submission today!
        </Typography>
        <Typography variant="body1" gutterBottom>
            You can submit again in:
        </Typography>
        <Paper elevation={3} sx={{ width: 200, textAlign: 'center', borderRadius: 1, backgroundColor: 'primary.main', color: 'primary.contrastText' }}>
            <Typography variant="h6" sx={{ p: 1 }}>
                {timeLeft}
            </Typography>
        </Paper>
    </>);
}