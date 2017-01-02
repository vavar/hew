<template>
  <div class="full-width">
      <md-toolbar>
        <h2 class="md-title md-left-align">Order List</h2>
        <add-order v-if="$auth.check()" :restaurant="restaurant" :activityID="activityID"></add-order>
      </md-toolbar>
      <md-table md-sort="restaurant" md-sort-type="desc" @sort="onSort">
        <md-table-header>
          <md-table-row>
            <md-table-head md-sort-by="user">User</md-table-head>
            <md-table-head md-sort-by="restaurant">Menu</md-table-head>
            <md-table-head md-sort-by="restaurant">Price</md-table-head>
            <md-table-head v-if="$auth.check()">action</md-table-head>
          </md-table-row>
        </md-table-header>
        <md-table-body>
          <md-table-row v-for="(row, rowIndex) in restaurant.orders" :key="rowIndex" :md-item="row" md-auto-select>
            <md-table-cell>
              <div>{{ lazyUserInfo(row.user_id) }}</div>
            </md-table-cell>
            <md-table-cell>
              <div>{{row.menu}}</div>
            </md-table-cell>
            <md-table-cell>
              <div>{{row.price || '-'}}</div>
            </md-table-cell>
            <md-table-cell v-if="$auth.check()">
              <md-button class="md-icon-button" v-if="isOwnerOrder(row.user_id)">
                <md-icon v-on:click="removeUser(menu)">delete</md-icon>
              </md-button>
            </md-table-cell>
          </md-table-row>
        </md-table-body>
      </md-table>

      <md-table-pagination
        md-size="10"
        md-total="20"
        md-page="1"
        md-label="Rows"
        md-separator="of"
        :md-page-options="[10, 20, 40]"
        @pagination="onPagination">
      </md-table-pagination>
  </div>
</template>

<script>
import AddOrder from './AddOrder';

export default {
  name: 'order',
  props: ['restaurant','activityID'],
  components: {
      AddOrder,
  },
  data: () => ({
    newMenu: {
      user: 'ton',
      name: '',
      price: '100',
      quantity: '',
    },
  }),
  created() {
    // fetch the data when the view is created and the data is
    // already being observed
    this.fetchData();
  },
  watch: {
    // call again the method if the route changes
    $route: 'fetchData',
  },
  methods: {
    fetchData() {
    },
    removeMenu() {
      this.child(this.menu['.key']).remove();
    },
    isOwnerOrder(id){
      return id === this.$auth.user().id;
    },
    onSort() {

    },
    lazyUserInfo( userID ) {
      const users = this.$store.state.users.filter( user=> user.id === userID);
      return users.length > 0 ? users[0].username : "";
    },
    onPagination() {

    },
  },
};
</script>

<style>
  .add-order-dialog {
    min-width: 600px;
  }
  .full-width {
    width: 100%;
  }
</style>
