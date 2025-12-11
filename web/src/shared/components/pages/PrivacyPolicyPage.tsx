import { useTenant } from "@/app/providers/TenantProvider";
import BigTextWrapper from "../atoms/BigTextWrapper";
import BodyText from "../atoms/BodyText";
import HeaderText from "../atoms/HeaderText";

export default function PrivacyPolicyPage() {
    const config = useTenant();

    const text = `At ${config.title}, we don\'t know how to sell data to advertisers yet, so you can rest assured knowing that none of the data you provide is being used to subliminally make you buy more stuff in the future.`
    return (
        <BigTextWrapper>
            <HeaderText text="Privacy Policy" />
            <BodyText text="First off, thanks for stopping by! This page can get a little lonely at times." />
            <BodyText text={text} />
            <BodyText text="For now, the only piece of information we collect is your email address, and that only gets used for sending verification emails at the moment. We might add a feature to send you annoying newsletters that you have to jump through the nine circles of Hell to opt out of in the future, but when that happens, we'll keep you posted (come back soon!)." />
            <BodyText text="Lastly, we do store some stuff in your browser (uwu). We save your theme preference, so whether you choose light or dark mode gets remembered. We also store a cookie that automatically logs you in if you've logged in within the last seven days." />
            <BodyText text="That's all I can remember at least, happy smashing!" />
        </BigTextWrapper>
    );
}