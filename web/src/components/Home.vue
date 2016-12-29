<template>
  <md-layout md-gutter>
    <md-layout md-column md-gutter md-flex="70" md-flex-offset="15">
      <md-card class="activity-card" v-for="act in activities">
        <md-toolbar md-theme="white" class="md-dense">
          <h2 class="md-title">{{act.name}}</h2>
          <span class="gutter"></span>
        </md-toolbar>
        <md-card-area>
          <md-tabs class="md-transparent activity-tabs" md-right>
            <md-tab v-bind:md-label="res.name" v-for="(res,index) in act.restaurants" md-gutter>
              <md-layout md-row md-flex-offset="5" >
                  <md-layout>
                    aa
                  </md-layout>
                  <md-layout md-flex="45">
                    b
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
            restaurants: [{ id: 1, name: 'ครัวคุณวี' }, { id: 2, name: 'DJ Poom' }],
          },
          {
            name: 'Tuesday Lunch',
            restaurants: [{ id: 1, name: 'ครัวคุณวี' }, { id: 2, name: 'DJ Poom' }],
          },
        ],
      };
    },
    components: {
      OrderToday,
    },
  };
</script>
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
