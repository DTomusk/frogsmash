import { Button } from "@mui/material";
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

    // Initialize once
    window.google.accounts.id.initialize({
      client_id: import.meta.env.VITE_GOOGLE_CLIENT_ID,
      callback: (response: any) => {
        onLogin(response.credential);
      },
    });
  }, [onLogin]);

  // Trigger Google popup manually
  const handleGoogleLogin = () => {
    if (!window.google) return;

    window.google.accounts.id.prompt(); // Opens Google login
  };

  return (
    <Button
      variant="contained"
      fullWidth
      onClick={handleGoogleLogin}
      startIcon={<img src="/google_logo.svg" width={20} height={20} />}
      sx={{
        backgroundColor: "white",
        color: "black",
        fontWeight: 600,
        border: "1px solid #ddd",
        py: 1,
        "&:hover": { backgroundColor: "#f7f7f7" }
      }}
    >
      Continue with Google
    </Button>
  );
}
