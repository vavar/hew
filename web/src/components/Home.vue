<template>
  <md-layout md-gutter v-if="!isLoading">
    <md-layout md-flex-small="10" md-flex-medium="10" md-flex-large="10" md-flex-xlarge="20"></md-layout>
    <md-layout md-flex-small="80" md-flex-medium="80" md-flex-large="80" md-flex-xlarge="60">
      <md-card class="activity-card" v-for="act in activities">
        <md-toolbar class="md-dense">
          <h2 class="md-title">
            <md-icon>event</md-icon>
            <span>{{act.name}}<sup>[{{act.id}}]</sup></span>
          </h2>
          <span class="gutter"></span>
        </md-toolbar>
        <md-table-card>
          <md-tabs class="md-transparent activity-tabs">
            <md-tab v-bind:md-label="res.name" v-for="(res,index) in restaurantList(act)" class="restaurant-tab">
              <order :restaurant="res" :activityID="act.id"></order>
            </md-tab>
          </md-tabs>
        </md-table-card>
      </md-card>
    </md-layout>
    <md-layout md-flex-small="10" md-flex-medium="10" md-flex-large="10" md-flex-xlarge="20"></md-layout>
  </md-layout>
</template>

<script>
  import Order from './Order';
  import _ from 'lodash';

  export default {
    name: 'home',
    data() {
      return {
      };
    },
    computed: {
      activities() {
        return this.$store.state.openActivities;
      },
      isLoading() {
        return this.$store.getters.isLoading;
      },
    },
    methods: {
      fetchData() {
        this.$store.dispatch('getHomeActivities', 1);
      },

      restaurantList(activity) {
        if (!activity || !activity.restaurants) {
          return;
        }
        return _.reduce(activity.restaurants, (restaurants, restaurant) => {
          const menus = restaurant.menus;

          // if ( !activity.order_items ) {
          restaurant.orders = (!activity.order_items) ? [] : _.reduce(activity.order_items, (_orders, order) => {
            menus.forEach((menu) => {
              if (menu.id === order.menu_id) {
                _orders.push({ user_id: order.user_id, menu: menu.name });
              }
            })
            return _orders;
          }, []);
          restaurants.push(restaurant);
          return restaurants;
          // }

          // restaurant.orders = _.reduce(activity.order_items, (_orders, order) => {
          //   menus.forEach((menu) => {
          //     if (menu.id === order.menu_id) {
          //       _orders.push({ user_id: order.user_id, menu: menu.name });
          //     }
          //     return _orders;
          //   })
          // },[]);
          // restaurants.push(restaurant);
          // return restaurants;
        }, []);
      },
    },
    created() {
      this.fetchData();
    },
    components: {
      Order,
    },
  };
</script>
<style>
.md-tab-header-container {
  min-width: 250px;
  font-size: 1.5em;
}
.menus {
  width: 100%;
}
.activity-button-bar{
  text-align: left;
}
</style>
<style scoped>
.activity-card {
  width: 100%;
  margin-bottom: 20px;
}
.activity-tabs {
  /*margin-top: -48px;*/
}
.gutter {
  flex: 1;
}
</style>
