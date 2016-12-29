<template>
  <div class="page content">
    <md-dialog md-open-from="#fab" md-close-to="#fab" ref="restaurantModal">
      <md-dialog-title>{{dialog.header}}</md-dialog-title>

      <md-dialog-content>
        <form>
          <md-input-container>
            <label>Name</label>
            <md-input required v-model="restaurant.name" placeholder="Restaurant Name"></md-input>
          </md-input-container>
          <md-input-container>
            <label>Phone Number</label>
            <md-input v-model="restaurant.phone" placeholder="Phone Number"></md-input>
          </md-input-container>
        </form>
      </md-dialog-content>

      <md-dialog-actions>
        <md-button class="md-primary" @click="addRestaurant('restaurantModal')">{{dialog.action}}</md-button>
        <md-button class="md-primary" @click="closeModal('restaurantModal')">Cancel</md-button>
      </md-dialog-actions>
    </md-dialog>
    <md-whiteframe md-elevation="2">
      <md-toolbar>
        <h1 class="md-title">Restaurants</h1>
        <md-button class="md-icon-button" id="custom" @click="openModal('restaurantModal')">
          <md-icon>add</md-icon>
        </md-button>
      </md-toolbar>
      <md-list>
        <md-list-item v-for="(row, rowIndex) in restaurants" class="restaurant-item">
          <div class="image-placeholder" @click="info(row.id)">
            <div class="lorem-image" v-bind:style="'background-image: url(\'http://lorempixel.com/128/128/food?'+rowIndex+'\')'"></div>
          </div>
          <div class="md-list-text-container" @click="info(rowIndex)">
            <span>{{row.name}}</span>
            <span>{{row.phone || '%phone number%'}}</span>
          </div>
          <md-button class="md-icon-button md-list-action" @click="openEditModal('restaurantModal',rowIndex)">
            <md-icon>create</md-icon>
          </md-button>
          <md-button class="md-icon-button md-list-action" >
            <md-icon>schedule</md-icon>
          </md-button>
        </md-list-item>
      </md-list>
    </md-whiteframe>
  </div>
</template>

<script>

  export default {
    name: 'restaurants',
    data() {
      return {
        dialog: { header: '', action: '' },
        restaurant: { id: '', name: '', phone: '' },
      };
    },
    computed: {
      restaurants() {
        return this.$store.state.restaurants;
      },
    },
    created() {
      this.$store.dispatch('getRestaurants');
    },
    methods: {
      openModal(ref) {
        this.dialog.header = 'Create new restaurant';
        this.restaurant.id = undefined;
        this.restaurant.name = '';
        this.restaurant.phone = '';
        this.dialog.action = 'Add';
        this.$refs[ref].open();
      },
      openEditModal(ref, rowIndex) {
        this.dialog.header = 'Update restaurant';
        this.dialog.action = 'Update';
        const restaurant = this.restaurants[rowIndex];

        this.restaurant.id = restaurant.id;
        this.restaurant.name = restaurant.name;
        this.restaurant.phone = restaurant.phone;
        this.$refs[ref].open();
      },
      addRestaurant(ref) {
        this.$refs[ref].close();
        this.$store.dispatch('updateRestaurant', this.restaurant);
      },
      info(rowIndex) {
        console.log('row index', rowIndex);
        const path = `/restaurant/${rowIndex}`;
        this.$router.push({ path });
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
      margin-top: 36px;
    }
    .restaurant-item {
      padding:10px;
    }
    .image-placeholder {
        width: 130px;
        height: 130px;
        background: #ddd;
        border: 1px solid grey;
        margin: 10px;
        margin-left: 0px;
        margin-right:20px;
    }
    .lorem-image {
      width: 128px;
      height: 128px;
      background-repeat: no-repeat;
      background-size: contain;
    }
</style>
