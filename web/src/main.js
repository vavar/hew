// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue';
import VueResource from 'vue-resource';
import VueMaterial from 'vue-material';
import VueRouter from 'vue-router';
import VeeValidate from 'vee-validate';

import auth from './auth';
import store from './store';
import App from './App';
import Home from './components/Home';
import Login from './components/Login';
import Order from './components/Order';
import Restaurants from './components/Restaurants';
import ManageEvents from './components/ManageEvents';
import RestaurantMenu from './components/RestaurantMenu';
import Signup from './components/Signup';

Vue.use(VueMaterial);
Vue.use(VueResource);
Vue.use(VeeValidate);

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
  // mode: 'history',
  // linkActiveClass: 'active',
  // base: __dirname,
  routes: [
    { path: '/', component: Home },
    { path: '/login', component: Login },
    { path: '/signup', component: Signup },
    { path: '/order', component: Order, beforeEnter: requireAuth },
    { path: '/restaurants', component: Restaurants, beforeEnter: requireAuth },
    { path: '/admin/events', component: ManageEvents, beforeEnter: requireAuth },
    { path: '/restaurant/:id', component: RestaurantMenu, beforeEnter: requireAuth },
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
