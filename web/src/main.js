// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue';
import VueResource from 'vue-resource';
import VueMaterial from 'vue-material';
import VueRouter from 'vue-router';
import VeeValidate from 'vee-validate';

import store from './store';
import App from './App';
import Home from './components/Home';
import Login from './components/Login';
import Order from './components/Order';
import OrderHistory from './components/OrderHistory';
import AddOrder from './components/AddOrder';
import Countdown from './components/countdown/Countdown';
import Register from './components/Register';
import Restaurants from './components/Restaurants';
import ManageEvents from './components/ManageEvents';
import RestaurantMenu from './components/RestaurantMenu';

Vue.use(VueMaterial);
Vue.use(VueResource);
Vue.use(VeeValidate);
Vue.use(VueRouter);

Vue.component('add-order', AddOrder);
Vue.component('countdown', Countdown);
Vue.filter('two_digits', function (value) {
  if (value.toString().length <= 1) {
    return "0" + value.toString();
  }
  return value.toString();
});

Vue.material.registerTheme('default', {
  primary: 'blue',
  accent: 'orange',
  warn: 'red',
  background: 'white',
});

Vue.http.options.root = 'http://localhost:8080';
// Vue.http.options.root = 'http://hew.abct.io';

Vue.router = new VueRouter({
  // mode: 'history',
  // linkActiveClass: 'active',
  // base: __dirname,
  routes: [
    { path: '/', name: 'home', component: Home },
    { path: '/login', name: 'login', component: Login, meta: { auth: false } },
    { path: '/register', name: 'register', component: Register, meta: { auth: false } },
    { path: '/history', name: 'history', component: OrderHistory, meta: { auth: true } },
    { path: '/restaurants', name: 'restaurants', component: Restaurants, meta: { auth: true } },
    { path: '/restaurants/:id', name: 'restaurant-info', component: RestaurantMenu },
    { path: '/admin/events', component: ManageEvents, meta: { auth: 'admin' }, },
  ],
});

/* eslint-disable no-new */
new Vue({
  router: Vue.router,
  store,
  el: '#app',
  template: '<App/>',
  components: { App },
});
