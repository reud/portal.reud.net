import Vue from 'vue'
import App from './App.vue'
import router from './router'
import vuetify from './plugins/vuetify';
import {Auth0Plugin} from "@/auth";
import {domain,clientId,audience} from "@/../auth_config.json";

Vue.config.productionTip = false;


Vue.use(Auth0Plugin, {
    domain,
    clientId,
    audience,
    // @ts-ignore
    onRedirectCallback: appState => {
        router.push(
            appState && appState.target
                ? appState.targetUrl
                : window.location.pathname
        );
    }
});

new Vue({
    router,
    vuetify,
    render: h => h(App)
}).$mount('#app');

