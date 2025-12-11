import { Typography, useTheme } from "@mui/material";
import Quote from "../molecules/Quote";
import { FROGSMASH_RULE_CONTENT, FROGSMASH_RULE_EXPLANATION_CONTENT, LIST_ITEM_1, LIST_ITEM_2, LIST_ITEM_3, LIST_ITEM_4, LIST_ITEM_5, LIST_ITEM_6, NO, QUOTE_1, QUOTE_2, WHAT_IS_FROGSMASH_CONTENT_1, WHAT_IS_FROGSMASH_CONTENT_1_2, WHAT_IS_FROGSMASH_CONTENT_1_3, WHAT_IS_FROGSMASH_CONTENT_1_4, WHAT_IS_FROGSMASH_CONTENT_1_5, WHAT_IS_FROGSMASH_CONTENT_2, WHAT_IS_FROGSMASH_CONTENT_2_1, WHAT_IS_FROGSMASH_CONTENT_2_2 } from "../content";
import LinkButton from "../atoms/LinkButton";
import BodyText from "../atoms/BodyText";
import HeaderText from "../atoms/HeaderText";
import BulletList from "../molecules/BulletList";
import BigTextWrapper from "../atoms/BigTextWrapper";

export default function IntroText() {
    const theme = useTheme();
    return (
        <BigTextWrapper>
            <Typography variant="h3" sx={{ textAlign: 'center', mb: 2, mt: 2 }}>
                What is FrogSmash?
            </Typography>
            <BodyText text={WHAT_IS_FROGSMASH_CONTENT_1} />
            <Quote text={QUOTE_1} author="You, probably" />
            <BodyText text={WHAT_IS_FROGSMASH_CONTENT_1_2} />
            <BodyText text={WHAT_IS_FROGSMASH_CONTENT_1_3} />
            <Quote text={QUOTE_2} author="You were definitely thinking this last week" />
            <BodyText text={WHAT_IS_FROGSMASH_CONTENT_1_4} />
            <HeaderText text={WHAT_IS_FROGSMASH_CONTENT_1_5} />
            <BodyText text={WHAT_IS_FROGSMASH_CONTENT_2} />
            <BulletList items={[LIST_ITEM_1, LIST_ITEM_2, LIST_ITEM_3]} />
            <BodyText text={WHAT_IS_FROGSMASH_CONTENT_2_1} />
            <BulletList items={[LIST_ITEM_4, LIST_ITEM_5, LIST_ITEM_6]} />
            <HeaderText text={NO} />
            <BodyText text={WHAT_IS_FROGSMASH_CONTENT_2_2} />
            <HeaderText text={FROGSMASH_RULE_CONTENT} />
            <BodyText text={FROGSMASH_RULE_EXPLANATION_CONTENT} />
            <LinkButton color={theme.palette.mode === 'light' ? 'primary' : 'secondary'} text="Start smashing!" to="/smash" />
        </BigTextWrapper>
    )
}