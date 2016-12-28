import Vue from 'vue';
import Vuex from 'vuex';

/* global console */
/* eslint no-console: ["error", { allow: ["warn", "error","log"] }] */

const URL = 'http://hew.abct.io/api/restaurants';
// const URL = 'http://localhost:8080/api/restaurants';
const fetchRestaurants = function fetchRestaurants(context) {
  Vue.http.get(URL).then((response) => {
    context.restaurants = response.body;
  });
};

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    restaurants: [],
  },
  mutations: {
    fetchRestaurants,
  },
  actions: {
    getRestaurants({ commit }) {
      commit('fetchRestaurants');
    },
    updateRestaurant(context, restaurant) {
      if (restaurant.id) {
        console.log('update !');
        Vue.http.put(URL, restaurant).then((response) => {
          console.log(response.body);
          context.commit('fetchRestaurants');
        });
      } else {
        Vue.http.post(URL, restaurant).then(() => {
          context.commit('fetchRestaurants');
        });
      }
    },
  },
});
