import { useTenant } from "@/app/providers/TenantProvider";
import BigTextWrapper from "../atoms/BigTextWrapper";
import BodyText from "../atoms/BodyText";
import HeaderText from "../atoms/HeaderText";

export default function TermsAndConditionsPage() {
    const config = useTenant();

    const tenantText = `Any images uploaded by the ${config.appTitle} team are either in the public domain or licensed under creative commons. We store the exact licence for each image (which you can see on the leaderboard page), as well as links to the original source.`
    return (
        <BigTextWrapper>
            <HeaderText text="Terms and conditions"/>
            <BodyText text="Hello there!"/>
            <BodyText text="I'm going to cut straight to the chase. There isn't a massive amount you can do on this website, so there aren't a great many ways to misuse it (although I'm always happy to hear ideas). " />
            {config.tenantKey === 'frog' && <><BodyText text="The main thing we do take seriously (surprisingly seriously, given how unserious the rest of the place is) is copyright law."/>
            <BodyText text={tenantText}/>
            <BodyText text="All we ask of you, our dear and lovely user, is, if you upload any images via the upload feature, that you own the copyright to those images. Uploaded images are reviewed by a member of our team before being added to the site, and we will do our best to ensure work that is obviously not the user's own isn't served on the platform. But at the end of the day, it's your responsibility to make sure that you own any work you upload. If you find there is any content on the site that shouldn't be there, please don't hesitate to contact us and we will be happy to remove it if the reasons for it are fair." />
            <BodyText text="And one last thing: please please please don't upload anything profane. This is a sacred place💅. Nothing disgusting will make its way to our users, but our team will have to review it, and they've seen enough as it is." /></>}
        </BigTextWrapper>
    );
}