import Vue from 'vue';
import Vuex from 'vuex';

/* global console */
/* eslint no-console: ["error", { allow: ["warn", "error","log"] }] */

const fetchRestaurants = function fetchRestaurants(context) {
  // console.log(state);
  Vue.http.get('http://hew.abct.io/api/restaurants').then((response) => {
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
  },
});
