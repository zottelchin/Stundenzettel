<template>
     <router-view id="container"></router-view>
</template>

<script>
import Vue from "vue";
import VueRouter from "vue-router";
import Stundenzettel from "./Stundenzettel"
Vue.use(VueRouter);
const router = new VueRouter({
    mode: "history",
    routes: [
        { path: "/", redirect: to => {
            return "/" + (new Date).getFullYear() + "/" + ((new Date).getMonth()+1)
        }},
        { path: "/:year/:month", component: Stundenzettel }
    ]
});

export default {
    router
}
</script>