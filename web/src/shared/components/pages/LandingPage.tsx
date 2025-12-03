import { Box, List, ListItem, ListItemText, Paper, Typography, useTheme } from "@mui/material";
import LinkButton from "../atoms/LinkButton";
import { FROGSMASH_RULE_CONTENT, FROGSMASH_RULE_EXPLANATION_CONTENT, LIST_ITEM_1, LIST_ITEM_2, LIST_ITEM_3, LIST_ITEM_4, LIST_ITEM_5, LIST_ITEM_6, NO, QUOTE_1, QUOTE_2, WHAT_IS_FROGSMASH_CONTENT_1, WHAT_IS_FROGSMASH_CONTENT_1_2, WHAT_IS_FROGSMASH_CONTENT_1_3, WHAT_IS_FROGSMASH_CONTENT_1_4, WHAT_IS_FROGSMASH_CONTENT_1_5, WHAT_IS_FROGSMASH_CONTENT_2, WHAT_IS_FROGSMASH_CONTENT_2_1, WHAT_IS_FROGSMASH_CONTENT_2_2 } from "../content";
import Quote from "../molecules/Quote";

function LandingPage() {
    const theme = useTheme(); 
    return <Box sx={{display: "flex", flexDirection: "column", alignItems: "center"}}>
            <Box
            sx={{
                width: "100%",
                height: {xs: 400, sm: 400, md: 500, lg: 600, xl: 800},
                backgroundSize: "cover",
                display: "flex",
                backgroundColor: theme.palette.mode === 'light' ? 'white' : 'primary.dark',
                backgroundImage: {xs: "none", 
                    sm:
                    theme.palette.mode === 'light'
                    ? `url(/froggy_light.png)`
                    : `url(/froggy_dark.png)`
                }
            }}
            >
            <Box sx={{ width: { xs: "100%", sm: "50%", md: "50%" }, height: "100%", display: "flex", flexDirection: "column", alignItems: "center", justifyContent: "center" }}>
                <Typography variant="h4" sx={{ textAlign: 'center', mt: 4, color: theme.palette.mode === 'light' ? 'black' : 'white' }}>
                    Compare frogs
                </Typography>
                <Typography variant="h4" sx={{ textAlign: 'center', mt: 4, color: theme.palette.mode === 'light' ? 'black' : 'white' }}>
                    Compare other things
                </Typography>
                <Typography variant="h4" sx={{ textAlign: 'center', mt: 4, color: theme.palette.mode === 'light' ? 'black' : 'white' }}>
                    Ribbit ribbit
                </Typography>
                <LinkButton color={theme.palette.mode === 'light' ? 'primary' : 'secondary'} text="Start smashing!" to="/smash" />
            </Box>
        </Box>
        <Box width={{ xs: "90%", sm: "70%", md: "50%" }}>
            <Paper elevation={3} sx={{ p: 2, pb: 4, textAlign: 'center', justifyContent: 'center', mt: 4, mb: 4, display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
                <Typography variant="h3" sx={{ textAlign: 'center', mb: 2, mt: 2 }}>
                    What is FrogSmash?
                </Typography>
                <Typography variant="body1" sx={{ textAlign: 'center', mt: 4, mx: 2 }}>
                    {WHAT_IS_FROGSMASH_CONTENT_1}
                </Typography>
                <Quote text={QUOTE_1} author="You, probably" />
                <Typography variant="body1" sx={{ textAlign: 'center', mt: 4, mx: 2 }}>
                    {WHAT_IS_FROGSMASH_CONTENT_1_2}
                </Typography>
                <Typography variant="body1" sx={{ textAlign: 'center', mt: 4, mx: 2 }}>
                    {WHAT_IS_FROGSMASH_CONTENT_1_3}
                </Typography>
                <Quote text={QUOTE_2} author="You were definitely thinking this last week" />
                <Typography variant="body1" sx={{ textAlign: 'center', mt: 4, mx: 2 }}>
                    {WHAT_IS_FROGSMASH_CONTENT_1_4}
                </Typography>
                <Typography variant="h4" sx={{ textAlign: 'center', mt: 4, mx: 2 }}>
                    {WHAT_IS_FROGSMASH_CONTENT_1_5}
                </Typography>
                <Typography variant="body1" sx={{ textAlign: 'center', mt: 4, mx: 2 }}>
                    {WHAT_IS_FROGSMASH_CONTENT_2}
                </Typography>
                <List sx={{ listStyleType: 'disc', pl: 4, width: {sx: "100%", md: "50%"} }}>
                    <ListItem sx={{ display: 'list-item' }}>
                        <ListItemText>{LIST_ITEM_1}</ListItemText>
                    </ListItem>
                    <ListItem sx={{ display: 'list-item' }}>
                        <ListItemText>{LIST_ITEM_2}</ListItemText>
                    </ListItem>
                    <ListItem sx={{ display: 'list-item' }}>
                        <ListItemText>{LIST_ITEM_3}</ListItemText>
                    </ListItem>
                </List>
                <Typography variant="body1" sx={{ textAlign: 'center', mt: 4, mx: 2 }}>
                    {WHAT_IS_FROGSMASH_CONTENT_2_1}
                </Typography>
                <List sx={{ listStyleType: 'disc', pl: 4, width: {sx: "100%", md: "50%"} }}>
                    <ListItem sx={{ display: 'list-item' }}>
                        <ListItemText>{LIST_ITEM_4}</ListItemText>
                    </ListItem>
                    <ListItem sx={{ display: 'list-item' }}>
                        <ListItemText>{LIST_ITEM_5}</ListItemText>
                    </ListItem>
                    <ListItem sx={{ display: 'list-item' }}>
                        <ListItemText>{LIST_ITEM_6}</ListItemText>
                    </ListItem>
                </List>
                <Typography variant="h4" sx={{ textAlign: 'center', mt: 4, mx: 2 }}>
                    {NO}
                </Typography>
                <Typography variant="body1" sx={{ textAlign: 'center', mt: 4, mx: 2 }}>
                    {WHAT_IS_FROGSMASH_CONTENT_2_2}
                </Typography>
                <Typography variant="h4" sx={{ textAlign: 'center', mt: 4, mx: 2 }}>
                    {FROGSMASH_RULE_CONTENT}
                </Typography>
                <Typography variant="body1" sx={{ textAlign: 'center', mt: 4, mx: 2 }}>
                    {FROGSMASH_RULE_EXPLANATION_CONTENT}
                </Typography>
                <LinkButton color={theme.palette.mode === 'light' ? 'primary' : 'secondary'} text="Start smashing!" to="/smash" />
            </Paper>
        </Box>
    </Box>;
}

export default LandingPage;