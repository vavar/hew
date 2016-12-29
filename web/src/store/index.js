import Vue from 'vue';
import Vuex from 'vuex';


/* global console */
/* eslint no-console: ["error", { allow: ["warn", "error","log"] }] */
const BASE_URL = 'http://hew.abct.io';
const RESTAURANTS_URL = `${BASE_URL}/api/restaurants`;
const ACTIVITIES_URL = `${BASE_URL}/api/activities`;
const MENU_URL = `${BASE_URL}/api/menus`;

function fetchRestaurants(state, context) {
  Vue.http.get(RESTAURANTS_URL).then((response) => {
    state.restaurants = response.body;
    state.restaurants.forEach((rest) => {
      state.restaurantMap[rest.id] = rest;
    });
    context.commit('loadingState', { isLoading: false });
  });
}

function fetchActivities(state, context) {
  Vue.http.get(ACTIVITIES_URL).then((response) => {
    state.restaurants = response.body;
    context.commit('loadingState', { isLoading: false });
  });
}

function fetchRestaurantInfo(state, context) {
  const info = `${RESTAURANTS_URL}/${context.id}`;
  Vue.http.get(info).then((response) => {
    const rest = response.body;
    state.restaurantMap[rest.id] = rest;
    state.restaurants.push(rest);
    state.activeRestaurant = rest;
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
    activeRestaurant: {},
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
      context.commit('loadingState', { isLoading: true });
      context.commit('fetchRestaurants', context);
    },
    getRestaurant(context, {id}) {
      if (!id) {
        return;
      }
      const state = context.state;
      const commit = context.commit;
      if (state.restaurantMap[id]) {
        activeRestaurant = state.restaurantMap[id];
        return;
      }
      commit('loadingState', { isLoading: true });
      commit('fetchRestaurantInfo', {store:context, id});
    },
    getActivities(context) {
      context.commit('loadingState', { isLoading: true });
      context.commit('fetchActivities', this);
    },
    addOrUpdateMenu(context, menu) {
      menu.price = +menu.price;
      if (menu.id) {
        Vue.http.put(MENU_URL, menu).then((response) => {
          console.log(response.body , menu.restaurant_id);
          context.commit('fetchRestaurantInfo',{store:context, id:""+menu.restaurant_id});
        });
      } else {
        Vue.http.post(MENU_URL, menu).then(() => {
          context.commit('fetchRestaurantInfo',{store:context, id:""+menu.restaurant_id});
        });
      }
    },
    addOrUpdateRestaurant(context, restaurant) {
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
