import { Box } from "@mui/material";
import { useEffect } from "react";

interface GoogleSignInButtonProps {
  onLogin: (credential: string) => void;
}

export default function GoogleSignInButton({ onLogin }: GoogleSignInButtonProps) {
  useEffect(() => {
    if (!window.google) {
      console.error("Google Identity Services script not loaded");
      return;
    }

    window.google.accounts.id.initialize({
      client_id: import.meta.env.VITE_GOOGLE_CLIENT_ID,
      callback: (response: any) => {
        onLogin(response.credential);
      },
    });

    window.google.accounts.id.renderButton(
      document.getElementById("google-signin-btn")!,
      {
        theme: "outline",
        size: "large",
        type: "standard",
        shape: "rectangular",
      }
    );
  }, [onLogin]);

  return <Box 
    display="flex" 
    justifyContent="center" 
    mt={2}>
        <div id="google-signin-btn"></div>
    </Box>
}
