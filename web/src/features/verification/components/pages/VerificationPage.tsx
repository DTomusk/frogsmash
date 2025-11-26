import { useSearchParams } from "react-router-dom";
import { useEffect, useState, type JSX } from "react";
import { useVerifyCode } from "../../hooks/useVerify";
import { AlreadyVerifiedCode, VerifiedCode } from "../../dtos/verificationResponse";
import { useCurrentUser } from "@/features/auth/hooks/useCurrentUser";
import { useQueryClient } from "@tanstack/react-query";
import PendingView from "../organisms/PendingView";
import LoggedInResendForm from "../organisms/LoggedInResendForm";
import AlreadyVerified from "../organisms/AlreadyVerified";
import VerifiedLoggedIn from "../organisms/VerifiedLoggedIn";
import VerifiedAnonymous from "../organisms/VerifiedAnonymous";
import ResendVerificationEmailForm from "../organisms/ResendVerificationEmailForm";

type VerificationStatus = "pending" 
    | "no_code_anonymous" 
    | "no_code_logged_in_unverified"
    | "logged_in_verified"
    | "code_error_logged_in"
    | "code_error_anonymous"
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
        if (isLoadingCurrentUser) {
            return;
        }
        if (currentUser && currentUser.isVerified) {
            setStatus("logged_in_verified");
            return;
        }
        if (!code) {
            if (!currentUser) {
                setStatus("no_code_anonymous");
            } else if (currentUser && !currentUser.isVerified) {
                setStatus("no_code_logged_in_unverified");
            } else {
                setStatus("logged_in_verified");
            }
            return;
        }
        
        verifyCode(code, {
            onSuccess: async (response) => {
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
                if (currentUser) {
                    setStatus("code_error_logged_in");
                } else {
                    setStatus("code_error_anonymous");
                }
            }
        });
    }, [code, isLoadingCurrentUser, currentUser]);

    const views: Record<VerificationStatus, JSX.Element> = {
    pending: <PendingView />,
    no_code_anonymous: <ResendVerificationEmailForm title="Verification required" message="Your account is not yet verified. Please try resending the verification email." />,
    no_code_logged_in_unverified: <LoggedInResendForm title="Verification required" message="Your account is not yet verified. Please click the button below to resend the verification email." />,
    logged_in_verified: <AlreadyVerified />,
    code_error_logged_in: <LoggedInResendForm title="Verification failed" message="There was an error verifying your account. Please try resending the verification email." />,
    code_error_anonymous: <ResendVerificationEmailForm title="Verification failed" message="There was an error verifying your account. Please try resending the verification email." />,
    code_success_logged_in: <VerifiedLoggedIn />,
    code_success_anonymous: <VerifiedAnonymous />
  };

    return ( <> 
    {isVerifying ? <PendingView /> : views[status] ?? <PendingView />}
    </>);
}

export default VerificationPage;