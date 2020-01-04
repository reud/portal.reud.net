import Vue from "vue";
import Component from 'vue-class-component'
import createAuth0Client from "@auth0/auth0-spa-js";
import { Emit } from 'vue-property-decorator';
import Auth0Client from "@auth0/auth0-spa-js/dist/typings/Auth0Client";

const DEFAULT_REDIRECT_CALLBACK = () =>
    window.history.replaceState({}, document.title, window.location.pathname);

let instance: any;

export const getInstance = () => instance;

export const useAuth0 = ({
                             onRedirectCallback = DEFAULT_REDIRECT_CALLBACK,
                             redirectUri = window.location.origin,
                             ...options
                         }) => {
    if (instance) return instance;
    @Component
    class authWrapper extends Vue {
        loading = true;
        isAuthenticated =  false;
        user = {};
        auth0Client: Auth0Client|null =  null;
        popupOpen =  false;
        error =  null;
        async loginWithPopup(o: any) {
            this.popupOpen = true;
            if (this.auth0Client == null) return;
            try {
                await this.auth0Client.loginWithPopup(o);
                this.user = await this.auth0Client.getUser();
                this.isAuthenticated = await this.auth0Client.isAuthenticated();
                this.error = null;
            } catch (e) {
                console.error(e);
                this.error = e;
            } finally {
                this.popupOpen = false;
            }
        }
        async handleRedirectCallback() {
            this.loading = true;
            if (this.auth0Client == null) return;
            try {
                await this.auth0Client.handleRedirectCallback();
                this.user = await this.auth0Client.getUser();
                this.isAuthenticated = true;
                this.error = null;
            } catch (e) {
                this.error = e;
            } finally {
                this.loading = false;
            }
        }
        loginWithRedirect(o:any) {
            if (this.auth0Client == null) return;
            return this.auth0Client.loginWithRedirect(o);
        }
        getIdTokenClaims(o:any) {
            if (this.auth0Client == null) return;
            return this.auth0Client.getIdTokenClaims(o);
        }
        getTokenSilently(o: any) {
            if (this.auth0Client == null) return;
            return this.auth0Client.getTokenSilently(o);
        }
        getTokenWithPopup(o:any) {
            if (this.auth0Client == null) return;
            return this.auth0Client.getTokenWithPopup(o);
        }
        logout(o:any) {
            if (this.auth0Client == null) return;
            return this.auth0Client.logout(o);
        }
        async created() {

            this.auth0Client = await createAuth0Client({
                // @ts-ignore
                domain: options.domain,
                // @ts-ignore
                client_id: options.clientId,
                // @ts-ignore
                audience: options.audience,
                // @ts-ignore
                redirect_uri: redirectUri
            });

            try {
                if (
                    window.location.search.includes("code=") &&
                    window.location.search.includes("state=")
                ) {
                    const {appState} = await this.auth0Client.handleRedirectCallback();
                    this.error = null;
                    // @ts-ignore
                    onRedirectCallback(appState);
                }
            } catch (e) {
                this.error = e;
            } finally {
                this.isAuthenticated = await this.auth0Client.isAuthenticated();
                this.user = await this.auth0Client.getUser();
                this.loading = false;
            }
        }
    }

    instance = new Vue(authWrapper);

    return instance;
};


export const Auth0Plugin = {
    install(Vue:any, options:any) {
        Vue.prototype.$auth = useAuth0(options);
    }
};
