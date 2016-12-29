import Vue from 'vue';
import Vuex from 'vuex';

/* global console */
/* eslint no-console: ["error", { allow: ["warn", "error","log"] }] */
const BASE_URL = 'http://hew.abct.io';
const RESTAURANTS_URL = `${BASE_URL}/api/restaurants`;
const ACTIVITIES_URL = `${BASE_URL}/api/activities`;

function fetchRestaurants({state, commit}) {
  Vue.http.get(RESTAURANTS_URL).then((response) => {
    state.restaurants = response.body;
    state.restaurants.forEach((rest) => {
      state.restaurantMap[rest.id] = rest;
    });
    commit('loadingState', { isLoading: false });
  });
}

function fetchActivities({state, commit}) {
  Vue.http.get(ACTIVITIES_URL).then((response) => {
    state.restaurants = response.body;
    commit('loadingState', { isLoading: false });
  });
}

function fetchRestaurantInfo(state, context) {
  const info = `${RESTAURANTS_URL}/${context.id}`;
  Vue.http.get(info).then((response) => {
    const rest = response.body;
    console.log('fetch me', rest);
    state.restaurantMap[rest.id] = rest;
    state.restaurants.push(rest);
    context.store.commit('loadingState', { isLoading: false });
  });
}

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    isLoading: false,
    restaurantMap: {},
    restaurants: [],
    activities: [],
  },
  getters: {
    isLoading(state) {
      return state.isLoading;
    },
    restaurantByID(state, id) {
      return state.restaurantMap[id];
    }
  },
  mutations: {
    fetchRestaurants,
    fetchRestaurantInfo,
    fetchActivities,
    loadingState(state, { isLoading }) {
      state.isLoading = isLoading
    },
  },
  actions: {
    getRestaurants(context) {
      context.commit('fetchRestaurants', context);
    },
    getRestaurant(context, {id}) {
      if (!id) {
        return;
      }
      const state = context.state;
      const commit = context.commit;
      if (state.restaurants) {
        const restaurants = state.restaurants.filter((restaurant) => restaurant.id === id);
        console.log(`exists !!`);
        if (restaurants.length > 0) {
          return restaurants[0];
        }
      }
      console.log(`fetchRestaurantInfo !! ${id}`);
      commit('loadingState', { isLoading: true });
      commit('fetchRestaurantInfo', {store:context, id});
    },
    getActivities(context) {
      context.commit('fetchActivities', this);
    },
    updateRestaurant(context, restaurant) {
      if (restaurant.id) {
        Vue.http.put(RESTAURANTS_URL, restaurant).then((response) => {
          console.log(response.body);
          context.commit('fetchRestaurants',context);
        });
      } else {
        Vue.http.post(RESTAURANTS_URL, restaurant).then(() => {
          context.commit('fetchRestaurants',context);
        });
      }
    },
  },
});
