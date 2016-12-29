<template>
  <md-layout md-gutter>
    <md-layout md-gutter md-flex="70" md-flex-offset="15">
      <md-card class="activity-card" v-for="act in activities">
        <md-toolbar md-theme="white" class="md-dense">
          <h2 class="md-title">
            <md-icon>event</md-icon>
            <span>{{act.name}}</span>
          </h2>
          <span class="gutter"></span>
        </md-toolbar>
        <md-card-area>
          <md-tabs class="md-transparent activity-tabs" md-right>
            <md-tab v-bind:md-label="res.name" v-for="(res,index) in act.restaurants" class="restaurant-tab">
              <md-layout md-gutter>
                <md-layout md-row md-gutter>
                  <md-layout md-flex="60">
                    <order-today></order-today>
                  </md-layout>
                  <md-layout>
                  </md-layout>
                </md-layout>
              </md-layout>
            </md-tab>
          </md-tabs>
        </md-card-area>
      </md-card>
    </md-layout>
  </md-layout>
</template>

<script>
  import OrderToday from './OrderToday';

  export default {
    name: 'home',
    data() {
      return {
        activities: [
          {
            name: 'Monday Lunch',
            restaurants: [
              { id: 1, name: 'ครัวคุณวี', menus: [{ id: 1, name: 'ข้าวกระเพราไก่', price: 80 }] },
              { id: 2, name: 'DJ Poom', menus: [{ id: 1, name: 'ข้าวกระเพราไก่', price: 80 }] },
            ],
          },
          {
            name: 'Tuesday Lunch',
            restaurants: [{ id: 1, name: 'ครัวคุณวี' }, { id: 2, name: 'DJ Poom' }],
          },
        ],
      };
    },
    // computed: {
    //   activities() {
    //     return this.$store.state.activities;
    //   },
    // },
    components: {
      OrderToday,
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
  margin-top: -48px;
}
.gutter {
  flex: 1;
}
</style>
