<template>
  <md-layout md-gutter>
    <md-layout md-flex-small="10" md-flex-medium="10" md-flex-large="10" md-flex-xlarge="20"></md-layout>
    <md-layout md-flex-small="80" md-flex-medium="80" md-flex-large="80" md-flex-xlarge="60">
      <md-card class="activity-card">
        <md-toolbar>
          <h2 class="md-title">History</h2>
          <span class="md-title-gutter"></span>
          <div>
            <md-button v-if="$auth.check('user')" @click="exportExcel(history)" class="md-icon-button">
              <md-icon>import_export</md-icon>
            </md-button>
            <div>
        </md-toolbar>
        <md-card-content>
          <div v-if="isLoading">
            <md-spinner :md-size="60" :md-stroke="1.5" md-indeterminate class="md-warn"></md-spinner>
          </div>
          <div v-else>
            <md-subheader>Orders</md-subheader>
            <md-table v-once>
              <md-table-header>
                <md-table-row>
                  <md-table-head>Activity</md-table-head>
                  <md-table-head>Menu</md-table-head>
                  <md-table-head md-numeric>Price</md-table-head>
                </md-table-row>
              </md-table-header>

              <md-table-body>
                <md-table-row v-for="item in history">
                  <md-table-cell>{{item.activity_name}}</md-table-cell>
                  <md-table-cell>{{item.menu_name}}</md-table-cell>
                  <md-table-cell md-numeric>{{item.menu_price}}</md-table-cell>
                </md-table-row>
              </md-table-body>
            </md-table>
          </div>
        </md-card-content>
      </md-card>
    </md-layout>
    <md-layout md-flex-small="10" md-flex-medium="10" md-flex-large="10" md-flex-xlarge="20"></md-layout>
  </md-layout>
</template>
<script>
  export default {
    name: 'order-history',
    computed: {
      isLoading() {
        return this.$store.getters.isLoading;
      },
      history() {
        return this.$store.state.history;
      },
    },
    methods: {
      fetchData() {
        this.$store.dispatch('getOrderHistory', this.$auth.user().id);
      },
      exportExcel(history) {
        if (!history || history.length == 0) {
          return;
        }
        const data = _.reduce(history, (_orders, order) => {
          _orders.push([order.activity_name, order.menu_name, order.menu_price]);
          return _orders;
        }, []);
        exportToExcel("order_history", _.concat([["activity", "menu", "price"]], data));
      },
    },
    created() {
      this.fetchData();
    },
  }
</script>
