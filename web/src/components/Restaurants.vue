<template>
  <md-layout md-gutter>
    <!-- ##dialog start -->
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
        <md-button class="md-primary" @click="addOrUpdateRestaurant('restaurantModal')">{{dialog.action}}</md-button>
        <md-button class="md-primary" @click="closeModal('restaurantModal')">Cancel</md-button>
      </md-dialog-actions>
    </md-dialog>
    <!-- ##dialog end -->
    <md-layout md-flex-small="10" md-flex-medium="10" md-flex-large="10" md-flex-xlarge="20"></md-layout>
    <md-layout md-flex-small="80" md-flex-medium="80" md-flex-large="80" md-flex-xlarge="60" class="content" v-if="!isLoading">
      <md-whiteframe md-elevation="2" md-tag="section" class="whiteframe-container">
        <md-toolbar>
          <h1 class="md-title"><md-icon>business</md-icon> <span>Restaurants</span></h1>
          <span class="md-title-gutter"></span>
          <md-button class="md-icon-button md-raised md-accent" id="custom" @click="openModal('restaurantModal')">
            <md-icon>add</md-icon>
          </md-button>
        </md-toolbar>
        <md-list>
          <md-list-item v-for="(row, rowIndex) in restaurants" class="restaurant-item">
            <div class="image-placeholder" @click="info(row)">
              <div class="lorem-image" v-bind:style="'background-image: url(\'http://lorempixel.com/128/128/food?'+rowIndex+'\')'"></div>
            </div>
            <div class="md-list-text-container" @click="info(row)">
              <span>{{row.name}}</span>
              <span>{{row.phone || '%phone number%'}}</span>
            </div>
            <md-button class="md-icon-button md-list-action" @click="openEditModal('restaurantModal',rowIndex)">
              <md-icon>create</md-icon>
            </md-button>
            <md-button class="md-icon-button md-list-action">
              <md-icon>schedule</md-icon>
            </md-button>
          </md-list-item>
        </md-list>
      </md-whiteframe>
    </md-layout>
    <md-layout md-flex-small="10" md-flex-medium="10" md-flex-large="10" md-flex-xlarge="20"></md-layout>
  </md-layout>
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
      isLoading() {
        return this.$store.getters.isLoading;
      },
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
      addOrUpdateRestaurant(ref) {
        this.$refs[ref].close();
        this.$store.dispatch('addOrUpdateRestaurant', this.restaurant);
      },
      info(row) {
        const path = `/restaurant/${row.id}`;
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
    .md-section{
      width: 100%;
    }
    .whiteframe-container {
      width:100%;
    }
    .content {
      text-align: center;
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
