import "babel-polyfill";
import Vue from "vue";
import App from "./App.vue";
import 'remixicon/fonts/remixicon.css'

import RestAPI from './rest-api';
window.api = new RestAPI("/api/v1");

new Vue({
    el: '#app',
    render: h => h(App)
});