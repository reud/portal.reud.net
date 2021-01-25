import Vue from "vue";
import App from "./App.vue";
import vuetify from "./plugins/vuetify";

const VueScrollTo = require("vue-scrollto");

Vue.config.productionTip = false;
Vue.use(VueScrollTo, {
  offset: -64
});

new Vue({
  vuetify,
  render: h => h(App)
}).$mount("#app");
