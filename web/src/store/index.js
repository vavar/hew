import Vue from 'vue';
import Vuex from 'vuex';
import Promise from 'bluebird';

/* global console */
/* eslint no-console: ["error", { allow: ["warn", "error","log"] }] */
const BASE_URL = 'http://localhost:8080';
// const BASE_URL = 'http://hew.abct.io';
const RESTAURANTS_URL = `${BASE_URL}/api/restaurants`;
const ACTIVITIES_URL = `${BASE_URL}/api/activities`;
const MENU_URL = `${BASE_URL}/api/menus`;
const USER_URL = `${BASE_URL}/api/users`;
const ORDER_URL = `${BASE_URL}/api/orders`;

function fetchRestaurants(state, context) {
  Vue.http.get(RESTAURANTS_URL).then((response) => {
    state.restaurants = response.body;
    state.restaurants.forEach((rest) => {
      state.restaurantMap[rest.id] = rest;
    });
    context.commit('loadingState', { isLoading: false });
  });
}

function fetchUsers(state, context) {
  // context.commit('loadingState', { isLoading: false });
  Vue.http.get(USER_URL).then((response) => {
    state.users = response.body;
    state.users.forEach((rest) => {
      state.userMap[rest.id] = rest;
    });
  });
}

function fetchActivities(state, context) {
  Vue.http.get(`${ACTIVITIES_URL}?org=${context.organizationId}&status=open`)
    .then((response) => {
      state.openActivities = response.body;
      return Vue.http.get(`${ACTIVITIES_URL}?org=${context.organizationId}&status=closed`);
    })
    .then((response) => {
      state.closedActivities = response.body;
      context.store.commit('loadingState', { isLoading: false });
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

function updateRestaurantsInActivity(activity, restaurantBool) {
  const url = `${ACTIVITIES_URL}/restaurants`;

  return Promise.coroutine(function* () {
    const restaurantIds = Object.keys(restaurantBool);
    for (let i = 0; i < restaurantIds.length; i += 1) {
      let restaurantId = restaurantIds[i];

      if (!restaurantBool[restaurantId]) {
        continue;
      }

      if (!activity.restaurants || !activity.restaurants.some(restaurant => restaurant.id === restaurantId)) {
        yield Vue.http.post(url, {
          activity_id: activity.id,
          restaurant_id: parseInt(restaurantId),
        });
      }
    }

    if (!activity.restaurants) {
      return;
    }

    for (let i = 0; i < activity.restaurants.length; i += 1) {
      let restaurant = activity.restaurants[i];
      if (!restaurantBool[restaurant.id]) {
        yield Vue.http.delete(url, {
          body : {
            activity_id: activity.id,
            restaurant_id: restaurant.id,
          }
        });
      }
    }
  })();
}

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    isLoading: false,
    user: {},
    restaurantMap: {},
    users:[],
    userMap: {},
    restaurants: [],
    openActivities: [],
    closedActivities: [],
    activeRestaurant: {},
  },
  getters: {
    isLoading(state) {
      return state.isLoading;
    },
    restaurantByID(state, id) {
      return state.restaurantMap[id];
    },
    userByID(state, id ) {
      return state.userMap[id];
    }
  },
  mutations: {
    fetchRestaurants,
    fetchRestaurantInfo,
    fetchActivities,
    fetchUsers,
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
        state.activeRestaurant = state.restaurantMap[id];
        return;
      }
      commit('loadingState', { isLoading: true });
      commit('fetchRestaurantInfo', {store:context, id});
    },
    getHomeActivities(context, organizationId) {
      context.commit('loadingState', { isLoading: true });
      context.commit('fetchUsers', { store: context, organizationId});
      context.commit('fetchActivities', { store: context, organizationId });
    },
    getUsers(context, organizationId) {
      context.commit('loadingState', { isLoading: true });
      context.commit('fetchUsers', { store: context, organizationId });
    },
    getActivities(context, organizationId) {
      context.commit('loadingState', { isLoading: true });
      context.commit('fetchActivities', { store: context, organizationId });
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
    updateActivity(context, { activity, restaurantBool }) {
      if (activity.id) {
        Vue.http.put(ACTIVITIES_URL, activity)
          .then((response) => {
            return updateRestaurantsInActivity(response.body, restaurantBool);
          })
          .then(() => {
            context.commit('fetchActivities', { store: context, organizationId: activity.organization_id });
          });
      } else {
        Vue.http.post(ACTIVITIES_URL, activity)
          .then((response) => {
            return updateRestaurantsInActivity(response.body, restaurantBool);
          })
          .then(() => {
            context.commit('fetchActivities', { store: context, organizationId: activity.organization_id });
          });
      }
    },
    addUser(context, user) {
      Vue.http.post(USER_URL, user).then(() => {});
    },
    addOrder(context, order) {
      Vue.http.post(ORDER_URL, order).then(()=>{
        context.commit('fetchActivities', { store: context, organizationId:1 });
      });
    }
  },
});
