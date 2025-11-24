import { useSearchParams } from "react-router-dom";
import { useEffect, useState, type JSX } from "react";
import { useVerifyCode } from "../../hooks/useVerify";
import { AlreadyVerifiedCode, VerifiedCode } from "../../dtos/verificationResponse";
import { useCurrentUser } from "@/features/auth/hooks/useCurrentUser";
import { useQueryClient } from "@tanstack/react-query";
import PendingView from "../organisms/PendingView";
import AnonymousResendForm from "../organisms/AnonymousResendForm";
import LoggedInResendForm from "../organisms/LoggedInResendForm";
import AlreadyVerified from "../organisms/AlreadyVerified";
import VerificationFailed from "../organisms/VerificationFailed";
import VerifiedLoggedIn from "../organisms/VerifiedLoggedIn";
import VerifiedAnonymous from "../organisms/VerifiedAnonymous";

// This page should just show components based on the user state and the presence of the code param
// 1. if code, verify
// 1.a. success
// 1.a.i user not logged in -> show success and prompt to log in
// 1.a.ii user logged in to same account as code -> show success and prompt to continue to app
// 1.a.iii user logged in to different account -> show success and prompt to log out and log in with correct account
// 1.b. error -> show error and prompt to resend verification email
// 1.b.i user not logged in -> show error and prompt to resend verification email
// 1.b.ii user logged in and not verified -> show error and prompt to resend verification email
// 1.b.iii user logged in and verified -> show already verified message and prompt to continue to app
// 2. if no code
// 2.a user not logged in -> show email entry to resend verification email
// 2.b user logged in and not verified -> show button to resend verification email
// 2.c user logged in and verified -> show already verified message and prompt to continue to app

type VerificationStatus = "pending" 
    | "no_code_anonymous" 
    | "no_code_logged_in_unverified"
    | "no_code_logged_in_verified"
    | "code_error"
    | "code_success_logged_in"
    | "code_success_anonymous";

function VerificationPage() {
    const [params] = useSearchParams();
    const code = params.get("code");

    const queryClient = useQueryClient();

    const { mutate: verifyCode, isPending: isVerifying } = useVerifyCode();
    const { data: currentUser, isLoading: isLoadingCurrentUser } = useCurrentUser();

    const [status, setStatus] = useState<VerificationStatus>("pending");

    useEffect(() => {
        if (!code) {
            // Call auth/me to determine whether user is logged in and verified
            // Show resend with email entry if not logged in
            // Show resend button if logged in and not verified
            // Show already verified message if logged in and verified
            if (!isLoadingCurrentUser) {
                if (!currentUser) {
                    setStatus("no_code_anonymous");
                } else if (currentUser && !currentUser.isVerified) {
                    setStatus("no_code_logged_in_unverified");
                } else {
                    setStatus("no_code_logged_in_verified");
                }
            }
            return;
        }

        console.log("Verifying code:", code);
        
        verifyCode(code, {
            onSuccess: async (response) => {
                console.log("Verification response code:", response.code);
                if (response.code === VerifiedCode || response.code === AlreadyVerifiedCode) {
                    
                }
                // Invalidate current user to update verified status
                await queryClient.invalidateQueries({queryKey: ['currentUser']});
                // Recall current user to get updated verified status
                if (currentUser) {
                    setStatus("code_success_logged_in");
                } else {
                    setStatus("code_success_anonymous");
                }

            },
            onError: () => {
                setStatus("code_error");
            }
        });
    }, [code]);

    const views: Record<VerificationStatus, JSX.Element> = {
    pending: <PendingView />,
    no_code_anonymous: <AnonymousResendForm />,
    no_code_logged_in_unverified: <LoggedInResendForm />,
    no_code_logged_in_verified: <AlreadyVerified />,
    code_error: <VerificationFailed />,
    code_success_logged_in: <VerifiedLoggedIn />,
    code_success_anonymous: <VerifiedAnonymous />
  };

    return ( <> 
    {isVerifying ? <PendingView /> : views[status] ?? <PendingView />}
    </>);
}

export default VerificationPage;