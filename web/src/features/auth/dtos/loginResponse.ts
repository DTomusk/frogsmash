export interface LoginResponse {
    jwt: string;
    user: {
        id: string;
        email: string;
        isVerified: boolean;
    };
}