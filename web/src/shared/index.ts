// TODO: split this barrel file into multiple smaller ones
// Pages
export { default as LandingPage } from "./components/pages/LandingPage";
export { default as LoadingPage } from "./components/pages/LoadingPage";
export { default as NotFoundPage } from "./components/pages/NotFoundPage";

// Templates
export { default as Template } from "./components/templates/Template";

// Themes
export { darkTheme, lightTheme } from "./theme/theme";

// API Client
export { apiFetch } from "./api/client";

// Atoms
export { default as FormWrapper } from "./components/atoms/FormWrapper";
export { default as EmailField } from "./components/atoms/EmailField";
export { default as StyledLink } from "./components/atoms/StyledLink";
export { default as ContentWrapper } from "./components/atoms/ContentWrapper";
export { default as ErrorMessage } from "./components/atoms/ErrorMessage";

// Molecules
export { default as LoadingSpinner } from "./components/molecules/LoadingSpinner";