<template>
  <div class="page content">
    <md-dialog md-open-from="#fab" md-close-to="#fab" ref="restaurantModal">
      <md-dialog-title>Create new restaurant</md-dialog-title>

      <md-dialog-content>
        <form>
          <md-input-container>
            <label>Name</label>
            <md-textarea></md-textarea>
          </md-input-container>
        </form>
      </md-dialog-content>

      <md-dialog-actions>
        <md-button class="md-primary" @click="addRestaurant('restaurantModal')">Add</md-button>
        <md-button class="md-primary" @click="closeModal('restaurantModal')">Cancel</md-button>
      </md-dialog-actions>
    </md-dialog>
    <md-table-card>
      <md-toolbar>
        <h1 class="md-title">Restaurant List</h1>
        <md-button class="md-icon-button" id="custom" @click="openModal('restaurantModal')">
          <md-icon>add</md-icon>
        </md-button>
      </md-toolbar> 

      <md-table md-sort="restaurant" md-sort-type="desc" @sort="onSort">
        <md-table-header>
          <md-table-row>
            <md-table-head md-sort-by="restaurant">Name</md-table-head>
            <md-table-head md-sort-by="phone">Phone No.</md-table-head>
            <md-table-head>action</md-table-head>
          </md-table-row>
        </md-table-header>

        <md-table-body>
          <md-table-row v-for="(row, rowIndex) in restaurant" :key="rowIndex" :md-item="row" md-auto-select>
            <md-table-cell v-for="(column, columnIndex) in row" :key="columnIndex">
              {{ column }}
            </md-table-cell>
            <md-table-cell>
              <md-button class="md-icon-button">
                <md-icon>edit</md-icon>
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
  name: 'restaurant',
  data: () => ({
    restaurant: [
      {
        name: 'Frozen yogurt',
        phone: '159',
      },
    ],
  }),
  methods: {
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
    .content {
      text-align: left;
    }
</style>