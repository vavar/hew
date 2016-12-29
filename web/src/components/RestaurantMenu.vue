<template>
  <md-layout md-gutter v-if="!isLoading">
    <md-dialog md-open-from="#fab" md-close-to="#fab" ref="menuModal">
      <md-dialog-title>Add new menu</md-dialog-title>

      <md-dialog-content>
        <form>
          <md-input-container>
            <label>Menu</label>
            <md-textarea></md-textarea>
          </md-input-container>
          <md-input-container>
            <label>Price</label>
            <md-textarea></md-textarea>
          </md-input-container>
        </form>
      </md-dialog-content>

      <md-dialog-actions>
        <md-button class="md-primary" @click="addMenu('menuModal')">Add</md-button>
        <md-button class="md-primary" @click="closeModal('menuModal')">Cancel</md-button>
      </md-dialog-actions>
    </md-dialog>

    <md-layout md-gutter md-row>
      <md-layout md-flex="60" md-flex-offset="20">
        <md-table-card>
          <md-toolbar>
            <h1 class="md-title">{{restaurant.name}}</h1>
            <md-button class="md-icon-button" id="custom" @click="openModal('menuModal')">
              <md-icon>playlist_add</md-icon>
            </md-button>
          </md-toolbar>
          <md-table md-sort="restaurant" md-sort-type="desc" @sort="onSort">
            <md-table-header>
              <md-table-row>
                <md-table-head md-sort-by="name">Menu</md-table-head>
                <md-table-head md-sort-by="price">price</md-table-head>
                <md-table-head>action</md-table-head>
              </md-table-row>
            </md-table-header>

            <md-table-body>
              <md-table-row v-for="(row, rowIndex) in restaurant.menus" :key="rowIndex" :md-item="row" md-auto-select>
                <md-table-cell v-for="(column, columnIndex) in [row.name, row.price]" :key="columnIndex">
                  {{ column }}
                </md-table-cell>
                <md-table-cell>
                  <md-button class="md-icon-button">
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
    </md-layout>
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
      menu: [
        {
          name: 'Frozen yogurt',
          price: '159',
        },
      ],
    }),
    computed: {
      restaurant(){
        console.log(this.restaurantID,this.$store.state.restaurantMap);
        return this.$store.state.restaurantMap[this.restaurantID];
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
        console.log( 'fetchData' , this.$route.params.id);
        this.error = this.post = null;
        this.loading = true;
        this.restaurantID = this.$route.params.id;
        this.$store.dispatch('getRestaurant', {id:this.restaurantID});
      },
      openModal(ref) {
        this.$refs[ref].open();
      },
      closeModal(ref) {
        this.$refs[ref].close();
      },
      addMenu() {

      },
      onSort() {

      },
      onPagination() {

      },
    },
  };
</script>

<style>
  .md-table-card {
    width: 100%;
  }
  .md-title-gutter {
    flex: 1;
  }
</style>
