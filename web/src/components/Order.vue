<template>
  <div class="full-width">
    <md-dialog md-open-from="#fab" md-close-to="#fab" ref="menuModal">
      <md-dialog-title>Add {{ restaurantName }}KhunWee menu</md-dialog-title>
      <md-dialog-content>
        <form>
          <md-input-container>
            <label>Menu</label>
            <md-select
              placeholder="Name"
              :name="'name'"
              v-model="menu.name">
              <md-option v-for="(value, rowIndex) in menu" :key="rowIndex" value={value.name}>{{ value.name }} ( {{ value.price }} baht )</md-option>
            </md-select>
          </md-input-container>
          <md-input-container>
            <label>Quantity</label>
            <md-select
              placeholder="Quantity"
              v-model="menu.quantity">
              <md-option v-for="(value, rowIndex) in 10" :key="rowIndex" value={value}>{{ value }}</md-option>
            </md-select>
          </md-input-container>
        </form>
      </md-dialog-content>

      <md-dialog-actions>
        <md-button class="md-primary" @click="addMenu('menuModal')">Add</md-button>
        <md-button class="md-primary" @click="closeModal('menuModal')">Cancel</md-button>
      </md-dialog-actions>
    </md-dialog>
      <md-toolbar class="md-accent">
        <h2 class="md-title" style="flex: 1">List</h2>
        <md-button v-if="loggedIn" class="md-raised md-button md-default" id="addMenu" @click="openModal('menuModal')">
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
          <md-table-row v-for="(row, rowIndex) in menu" :key="rowIndex" :md-item="row" md-auto-select>
            <md-table-cell v-for="(column, columnIndex) in row" :key="columnIndex" :md-numeric="columnIndex !== 'user' && columnIndex !== 'name'">
              {{ column }}
            </md-table-cell>
            <md-table-cell>
              <md-button class="md-icon-button">
                <md-icon>delete</md-icon>
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

export default {
  name: 'order',
  props: ['restaurant-name'],
  data: () => ({
    loggedIn: auth.loggedIn(),
    loading: false,
    post: null,
    error: null,
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
      this.error = this.post = null;
      this.loading = true;
      // replace getPost with your data fetching util / API wrapper
      // getPost(this.$route.params.id, (err, post) => {
      //   this.loading = false;
      //   if (err) {
      //     this.error = err.toString();
      //   } else {
      //     this.post = post;
      //   }
      // });
    },
    openModal(ref) {
      this.$refs[ref].open();
    },
    closeModal(ref) {
      this.$refs[ref].close();
    },
    onSort() {

    },
    onPagination() {

    },
  },
};
</script>

<style>
  .full-width {
    width: 100%;
  }
</style>
