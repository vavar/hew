<template>
  <div class="full-width">
    <add-order :restaurant="restaurant"></add-order>
      <md-toolbar>
        <h2 class="md-title md-left-align">Order List</h2>
        <md-button v-if="$auth.ready()" class="md-raised md-button md-warn" id="addMenu" @click="openModal('menuModal')">
          <md-icon>add</md-icon><span> Add order</span>
        </md-button>
      </md-toolbar>
      <md-table md-sort="restaurant" md-sort-type="desc" @sort="onSort">
        <md-table-header>
          <md-table-row>
            <md-table-head md-sort-by="user">User</md-table-head>
            <md-table-head md-sort-by="restaurant">Menu</md-table-head>
            <md-table-head md-sort-by="price" md-numeric>price</md-table-head>
            <md-table-head md-sort-by="quantity" md-numeric>quantity</md-table-head>
            <md-table-head>action</md-table-head>
          </md-table-row>
        </md-table-header>

        <md-table-body>
          <md-table-row v-for="(row, rowIndex) in restaurant.menus" :key="rowIndex" :md-item="row" md-auto-select>
            <md-table-cell v-for="(column, columnIndex) in row" :key="columnIndex" :md-numeric="columnIndex !== 'user' && columnIndex !== 'name'">
              {{ column }}
            </md-table-cell>
            <md-table-cell>
              <md-button v-if="loggedIn" class="md-icon-button">
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
import auth from '../auth';
import AddOrder from './AddOrder';
export default {
  name: 'order',
  props: ['restaurant','orders'],
  components: {
      AddOrder,
  },
  data: () => ({
    loggedIn: auth.loggedIn(),
    newMenu: {
      user: 'ton',
      name: '',
      price: '100',
      quantity: '',
    },
    menu: [
      {
        user: 'ton',
        name: 'Frozen yogurt',
        price: '159',
        quantity: '1',
      },
      {
        user: 'ken',
        name: 'Fried rice',
        price: '200',
        quantity: '2',
      },
    ],
  }),
  created() {
    // fetch the data when the view is created and the data is
    // already being observed
    this.fetchData();
    auth.onChange = (loggedIn) => {
      this.loggedIn = loggedIn;
    };
  },
  watch: {
    // call again the method if the route changes
    $route: 'fetchData',
  },
  methods: {
    fetchData() {
    },
    openModal(ref) {
      this.$refs[ref].open();
    },
    closeModal(ref) {
      this.$refs[ref].close();
    },
    addMenu(ref) {
      console.log(this.newMenu);
      this.menu.push(this.newMenu);
      this.$refs[ref].close();
      this.newMenu.name = '';
      this.newMenu.price = '';
      this.newMenu.quantity = '';
    },
    removeMenu() {
      this.child(this.menu['.key']).remove();
    },
    onSort() {

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
