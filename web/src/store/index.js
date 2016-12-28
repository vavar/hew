import Vue from 'vue';
import Vuex from 'vuex';

/* global console */

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    restaurants: [],
  },
  // getters: {
  //   restaurants: fetchRestaurants,
  // },
  mutations: {
    fetchRestaurants(context) {
      Vue.http.get('http://hew.abct.io/api/restaurants').then((response) => {
        context.dispatch('restaurants', response.json());
      }, () => {
        context.dispatch('restaurants', []);
      });
    },
  },
  actions: {
    getRestaurants({ commit }) {
      commit('fetchRestaurants');
    },
  },
});
