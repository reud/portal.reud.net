import {getInstance} from "./authWrapper";

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
