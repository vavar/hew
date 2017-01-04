<template>
  <div class="full-width">
    <md-dialog-alert :md-content="alert.content" :md-ok-text="alert.ok" @open="onOpen" @close="onClose" ref="dialog3">
    </md-dialog-alert>

    <md-toolbar class="md-transparent">
      <h2 class="md-title md-left-align ">Order List</h2>
      <span style="flex:1"></span>
      <add-order v-if="$auth.check()" :activity="activity"></add-order>
    </md-toolbar>
    <md-table md-sort="restaurant" md-sort-type="desc" @sort="onSort">
      <md-table-header>
        <md-table-row>
          <md-table-head md-sort-by="user">User</md-table-head>
          <md-table-head md-sort-by="restaurant">Restaurant</md-table-head>
          <md-table-head md-sort-by="restaurant">Menu</md-table-head>
          <md-table-head md-sort-by="restaurant">Price</md-table-head>
          <md-table-head v-if="$auth.check()">action</md-table-head>
        </md-table-row>
      </md-table-header>
      <md-table-body>
        <md-table-row v-for="(row, rowIndex) in orderList(activity)" :key="rowIndex" :md-item="row" md-auto-select>
          <md-table-cell>
            <div>{{ lazyUserInfo(row.user_id) }}</div>
          </md-table-cell>
          <md-table-cell>
            <div>{{row.restaurant}}</div>
          </md-table-cell>
          <md-table-cell>
            <div>{{row.menu}}</div>
          </md-table-cell>
          <md-table-cell>
            <div>{{row.price || '-'}}</div>
          </md-table-cell>
          <md-table-cell v-if="$auth.check()">
            <md-button class="md-icon-button" v-if="isOwnerOrder(row.user_id)" v-on:click="removeUser(row.user_id, row.menu_id)">
              <md-icon >delete</md-icon>
            </md-button>
          </md-table-cell>
        </md-table-row>
      </md-table-body>
    </md-table>
  </div>
</template>

<script>
  import AddOrder from './AddOrder';

  export default {
    name: 'order',
    props: ['activity'],
    components: {
      AddOrder,
    },
    data: () => ({
      alert: {
        content: 'Comming Soon !!',
        ok: 'Cool!'
      },
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
      isOwnerOrder(id) {
        return id === this.$auth.user().id || this.$auth.check('admin');
      },
      onSort() {

      },

      onOpen() {
        console.log('Opened');
      },
      onClose(type) {
        console.log('Closed', type);
      },
      removeUser(userID, menuID) {
        this.openDialog("dialog3");
      },
      openDialog(ref) {
        this.$refs[ref].open();
      },
      closeDialog(ref) {
        this.$refs[ref].close();
      },
      orderList(activity) {
        if (!activity || !activity.restaurants || !activity.order_items || !activity.order_items.length) {
          return [];
        }
        return (!activity.order_items) ? [] : _.reduce(activity.order_items, (_orders, order) => {
          activity.restaurants.forEach((restaurant) => {
            const menus = restaurant.menus;
            menus.forEach((menu) => {
              if (menu.id === order.menu_id) {
                _orders.push({
                  user_id: order.user_id,
                  restaurant: restaurant.name,
                  menu_id: menu.id,
                  menu: menu.name,
                  price: menu.price
                });
              }
            });
          });
          return _orders;
        }, []);
      },
      lazyUserInfo(userID) {
        const users = this.$store.state.users.filter(user => user.id === userID);
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
