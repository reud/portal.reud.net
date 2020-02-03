import {getInstance} from "./authWrapper";
// @ts-ignore
import * as jwtCode from 'jwt-decode';


export const authGuard = (to: any, from: any, next: any) => {
    const authService = getInstance();

    const fn = () => {
        if (authService.isAuthenticated) {
            return next();
        }

        authService.loginWithRedirect({appState: {targetUrl: to.fullPath}});
    };

    if (!authService.loading) {
        return fn();
    }

    authService.$watch("loading", (loading: boolean) => {
        if (!loading) {
            return fn();
        }
    });
};

export const reudGuard = async (to: any, from: any, next: any) => {
    const authService = getInstance();

    const fn = async () => {
        if (authService.isAuthenticated) {
            const reudSub = "auth0|5d95a4585c8d900dd0c0327c";
            const decoded = jwtCode(await authService.getTokenSilently());
            if (decoded.sub == reudSub) {
                return next();
            }
            return next('/');
        }

        authService.loginWithRedirect({appState: {targetUrl: to.fullPath}});
    };

    if (!authService.loading) {
        return await fn();
    }

    authService.$watch("loading", async (loading: boolean) => {
        if (!loading) {
            return await fn();
        }
    });
};
