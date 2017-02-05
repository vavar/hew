<template>
  <md-layout md-gutter v-if="!isLoading">
    <md-dialog md-open-from="#fab" md-close-to="#fab" ref="menuModal">
      <md-dialog-title>{{dialog.header}}</md-dialog-title>

      <md-dialog-content>
        <form>
          <md-input-container>
            <label>Menu</label>
            <md-input required v-model="menu.name" placeholder="Menu Name"></md-input>
          </md-input-container>
          <md-input-container>
            <label>Price</label>
            <md-input v-model="menu.price" placeholder="Price"></md-input>
          </md-input-container>
        </form>
      </md-dialog-content>

      <md-dialog-actions>
        <md-button class="md-primary" @click="addOrUpdateMenu('menuModal')">{{dialog.action}}</md-button>
        <md-button class="md-primary" @click="closeModal('menuModal')">Cancel</md-button>
      </md-dialog-actions>
    </md-dialog>

    <md-layout md-flex-small="10" md-flex-medium="10" md-flex-large="10" md-flex-xlarge="20"></md-layout>
    <md-layout md-flex-small="80" md-flex-medium="80" md-flex-large="80" md-flex-xlarge="60" >
      <md-table-card>
        <md-toolbar class="md-accent">
          <h1 class="md-title md-left-align">
            <md-icon>store</md-icon> {{restaurant.name}}</h1>
          <md-button class="md-icon-button md-raised md-warn" id="custom" @click="openModal('menuModal')">
            <md-icon>playlist_add</md-icon>
          </md-button>
        </md-toolbar>
        <md-table md-sort="restaurant" md-sort-type="desc" @sort="onSort">
          <md-table-header>
            <md-table-row>
              <md-table-head md-sort-by="name">Menu</md-table-head>
              <md-table-head md-sort-by="price">price</md-table-head>
              <md-table-head >action</md-table-head>
            </md-table-row>
          </md-table-header>

          <md-table-body>
            <md-table-row v-for="(row, rowIndex) in restaurant.menus" :key="rowIndex" :md-item="row" md-auto-select>
              <md-table-cell v-for="(column, columnIndex) in [row.name, row.price]" :key="columnIndex">
                {{ column }}
              </md-table-cell>
              <md-table-cell>
                <md-button class="md-icon-button" @click="openEditModal('menuModal',row)">
                  <md-icon>create</md-icon>
                </md-button>
              </md-table-cell>
            </md-table-row>
          </md-table-body>
        </md-table>

        <md-table-pagination md-size="5" md-total="10" md-page="1" md-label="Rows" md-separator="of" :md-page-options="[5, 10, 25, 50]"
          @pagination="onPagination"></md-table-pagination>

      </md-table-card>
    </md-layout>
    <md-layout md-flex-small="10" md-flex-medium="10" md-flex-large="10" md-flex-xlarge="20"></md-layout>
  </md-layout>
</template>

<script>
  import { mapGetters } from 'vuex'
  export default {
    name: 'restaurant-menu',
    data: () => ({
      restaurantID: null,
      post: null,
      error: null,
      dialog: { header: '', action: '' },
      menu: { id: '', name: '', price: '', restaurant_id: '' },
    }),
    computed: {
      restaurant() {
        return this.$store.state.activeRestaurant;
      },
      isLoading() {
        return this.$store.getters.isLoading;
      }
    },
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
        this.error = this.post = null;
        this.loading = true;
        this.restaurantID = this.$route.params.id;
        this.$store.dispatch('getRestaurant', { id: this.restaurantID });
      },
      openModal(ref) {
        this.dialog.header = 'Create new Menu';
        this.dialog.action = 'Add';
        this.menu.id = undefined;
        this.menu.name = '';
        this.menu.price = '';
        this.menu.restaurant_id = +this.restaurant.id;
        this.$refs[ref].open();
      },
      openEditModal(ref, row) {
        this.dialog.header = 'Update Menu';
        this.dialog.action = 'Update';
        const menu = row;

        this.menu.id = menu.id;
        this.menu.name = menu.name;
        this.menu.price = menu.price;
        this.menu.restaurant_id = menu.restaurant_id;
        this.$refs[ref].open();
      },

      addOrUpdateMenu(ref) {
        this.$refs[ref].close();
        this.$store.dispatch('addOrUpdateMenu', this.menu);
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
  .md-left-align {
    text-align: left;
  }
  .md-table-card {
    width: 100%;
  }
  .md-title-gutter {
    flex: 1;
  }
</style>
