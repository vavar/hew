import Vue from 'vue';
import Vuex from 'vuex';

/* global console */
/* eslint no-console: ["error", { allow: ["warn", "error","log"] }] */
const BASE_URL = 'http://hew.abct.io';
const RESTAURANTS_URL = `${BASE_URL}/api/restaurants`;
const ACTIVITIES_URL = `${BASE_URL}/api/activities`;

const fetchRestaurants = function fetchRestaurants(context) {
  Vue.http.get(RESTAURANTS_URL).then((response) => {
    context.restaurants = response.body;
  });
};

const fetchActivities = function fetchActivities(context) {
  Vue.http.get(ACTIVITIES_URL).then((response) => {
    context.restaurants = response.body;
  });
};

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    restaurants: [],
    activities: [],
  },
  mutations: {
    fetchRestaurants,
    fetchActivities,
  },
  actions: {
    getRestaurants({ commit }) {
      commit('fetchRestaurants');
    },
    getActivities({ commit }) {
      commit('fetchActivities');
    },
    updateRestaurant(context, restaurant) {
      if (restaurant.id) {
        Vue.http.put(RESTAURANTS_URL, restaurant).then((response) => {
          console.log(response.body);
          context.commit('fetchRestaurants');
        });
      } else {
        Vue.http.post(RESTAURANTS_URL, restaurant).then(() => {
          context.commit('fetchRestaurants');
        });
      }
    },
  },
});
