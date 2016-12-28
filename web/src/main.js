// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue';
import VueResource from 'vue-resource';
import Vuex from 'vuex';
import VueMaterial from 'vue-material';
import VueRouter from 'vue-router';

import auth from './auth';
import store from './store';
import App from './App';
import Home from './components/Home';
import Login from './components/Login';
import Order from './components/Order';
import Restaurant from './components/Restaurant';
import ManageEvents from './components/ManageEvents';

Vue.use(VueMaterial);
Vue.use(VueResource);
Vue.use(Vuex);

Vue.material.registerTheme('default', {
  primary: 'blue',
  accent: 'orange',
  warn: 'red',
  background: 'white',
});
Vue.use(VueRouter);

function requireAuth(to, from, next) {
  if (!auth.loggedIn()) {
    next({
      path: '/login',
      query: { redirect: to.fullPath },
    });
  } else if (to.fullPath.startsWith('/admin') && !auth.isAdmin()) {
    next('/');
  } else {
    next();
  }
}

const router = new VueRouter({
  mode: 'history',
  linkActiveClass: 'active',
  base: __dirname,
  routes: [
    { path: '/', component: Home },
    { path: '/login', component: Login },
    { path: '/order', component: Order, beforeEnter: requireAuth },
    { path: '/restaurant', component: Restaurant, beforeEnter: requireAuth },
    { path: '/admin/events', component: ManageEvents, beforeEnter: requireAuth },
    { path: '/logout',
      beforeEnter(to, from, next) {
        auth.logout();
        next('/');
      },
    },
  ],
});

/* eslint-disable no-new */
new Vue({
  router,
  store,
  el: '#app',
  template: '<App/>',
  components: { App },
});
