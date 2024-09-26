import { jwtDecode } from "jwt-decode";

export const decryptToken = (token: string) => {
    try {
        const decodedPayload = jwtDecode(token);
        return decodedPayload;
    } catch (error) {
        console.error('Failed to decode token', error);
        return null;
    }
};

export const isTokenExpired = (token: string): boolean => {
    const decodedToken: any = decryptToken(token);

    if (!decodedToken || !decodedToken.exp) {
        return true;
    }

    const currentTime = Math.floor(Date.now() / 1000);
    return decodedToken.exp < currentTime;
};
