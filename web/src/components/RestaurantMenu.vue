<template>
  <div class="page content">
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
    <md-table-card>
      <md-toolbar>
        <h1 class="md-title">Restaurant Name :{{ $route.params.id }}</h1>
        <md-button class="md-icon-button" id="custom" @click="openModal('menuModal')">
          <md-icon>playlist_add</md-icon>
        </md-button>
      </md-toolbar>
      <md-table md-sort="restaurant" md-sort-type="desc" @sort="onSort">
        <md-table-header>
          <md-table-row>
            <md-table-head md-sort-by="restaurant">Menu</md-table-head>
            <md-table-head md-sort-by="price">price</md-table-head>
            <md-table-head>action</md-table-head>
          </md-table-row>
        </md-table-header>

        <md-table-body>
          <md-table-row v-for="(row, rowIndex) in menu" :key="rowIndex" :md-item="row" md-auto-select>
            <md-table-cell v-for="(column, columnIndex) in row" :key="columnIndex">
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
        md-size="5"
        md-total="10"
        md-page="1"
        md-label="Rows"
        md-separator="of"
        :md-page-options="[5, 10, 25, 50]"
        @pagination="onPagination"></md-table-pagination>
    </md-table-card>
  </div>
</template>

<script>
export default {
  name: 'restaurant-menu',
  data: () => ({
    loading: false,
    post: null,
    error: null,
    menu: [
      {
        name: 'Frozen yogurt',
        price: '159',
      },
    ],
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
    addMenu() {

    },
    onSort() {

    },
    onPagination() {

    },
  },
};
</script>
