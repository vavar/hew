<template>
  <md-layout md-gutter v-if="!isLoading">
    <md-dialog md-open-from="#fab" md-close-to="#fab" ref="activityModal">
      <md-dialog-title>{{dialog.header}}</md-dialog-title>
      <md-dialog-content class="manage-events-dialog">
        <form>
          <md-input-container>
            <label>Name</label>
            <md-input required v-model="activity.name" placeholder="Event Name"></md-input>
          </md-input-container>
          <md-input-container>
            <label>Date</label>
            <md-input required type="date" v-model="date"></md-input>
          </md-input-container>
          <md-input-container>
            <label>Closed At</label>
            <md-input required type="time" v-model="time"></md-input>
          </md-input-container>
          <div class="text-left">
            <label>Restaurants</label>
            <md-checkbox class="restaurant-option" v-for="restaurant in restaurants" v-model="restaurantBool[restaurant.id]">{{restaurant.name}}</md-checkbox>
          </div>
        </form>
      </md-dialog-content>
      <md-dialog-actions>
        <md-button class="md-primary" @click="updateActivity('activityModal')">{{dialog.action}}</md-button>
        <md-button class="md-primary" @click="closeModal('activityModal')">Cancel</md-button>
      </md-dialog-actions>
    </md-dialog>
    <md-layout md-flex-small="10" md-flex-medium="10" md-flex-large="10" md-flex-xlarge="20"></md-layout>
    <md-layout md-flex-small="80" md-flex-medium="80" md-flex-large="80" md-flex-xlarge="60">
      <md-card class="activity-card">
        <md-toolbar>
          <h1 class="md-title">Manage Open Events</h1>
          <span class="md-title-gutter"></span>
          <md-button class="md-icon-button md-raised md-accent" id="custom" @click="openModal('activityModal')">
            <md-icon>add</md-icon>
          </md-button>
        </md-toolbar>
        <md-card-content>
          <md-table>
            <md-table-header>
              <md-table-row>
                <md-table-head>Name</md-table-head>
                <md-table-head>Closed At</md-table-head>
                <md-table-head>&nbsp;</md-table-head>
              </md-table-row>
            </md-table-header>

            <md-table-body>
              <md-table-row v-for="(row, rowIndex) in openActivities" :key="rowIndex" :md-item="row" md-auto-select>
                <md-table-cell>{{row.name}}</md-table-cell>
                <md-table-cell>{{formatDate(row.closed_at)}}</md-table-cell>
                <md-table-cell>
                  <md-button class="md-icon-button" @click="openEditModal('activityModal', row)">
                    <md-icon>edit</md-icon>
                  </md-button>
                </md-table-cell>
              </md-table-row>
            </md-table-body>
          </md-table>
        </md-card-content>
      </md-card>

      <md-card class="activity-card">
        <md-toolbar>
          <h1 class="md-title">Past Events</h1>
        </md-toolbar>
        <md-card-content>
          <md-table md-sort="closed-at" md-sort-type="desc">
            <md-table-header>
              <md-table-row>
                <md-table-head md-sort-by="name">Name</md-table-head>
                <md-table-head md-sort-by="closed-at">Closed At</md-table-head>
              </md-table-row>
            </md-table-header>

            <md-table-body>
              <md-table-row v-for="(row, rowIndex) in closedActivities" :key="rowIndex" :md-item="row" md-auto-select>
                <md-table-cell>{{row.name}}</md-table-cell>
                <md-table-cell>{{formatDate(row.closed_at)}}</md-table-cell>
              </md-table-row>
            </md-table-body>
          </md-table>
        </md-card-content>
      </md-card>
    </md-layout>
    <md-layout md-flex-small="10" md-flex-medium="10" md-flex-large="10" md-flex-xlarge="20"></md-layout>
  </md-layout>
</template>

<script>
  import utils from '../utils';

  export default {
    name: 'manage-events',
    data: () => ({
      dialog: { header: '', action: '' },
      date: '',
      time: '',
      activity: { id: '', name: '', closed_at: '', organization_id: 1 },
      restaurantBool: {},
    }),
    computed: {
      isLoading() {
        return this.$store.getters.isLoading;
      },
      openActivities() {
        return this.$store.state.openActivities;
      },
      closedActivities() {
        return this.$store.state.closedActivities;
      },
      restaurants() {
        return this.$store.state.restaurants;
      },
    },
    created() {
      this.$store.dispatch('getActivities', this.activity.organization_id);
      this.$store.dispatch('getRestaurants');
    },
    methods: {
      openModal(ref) {
        const localTime = utils.getNextDayTime().replace('Z', '').split('T');
        this.dialog.header = 'Create a new event';
        this.dialog.action = 'Add';
        this.activity.id = undefined;
        this.activity.name = '';
        this.date = localTime[0];
        this.time = localTime[1];
        this.restaurantBool = {};
        this.restaurants.map((restaurant) => {
          this.restaurantBool[restaurant.id] = false;
        });
        this.$refs[ref].open();
      },
      openEditModal(ref, row) {
        const activity = row;
        const time = utils.formatDateForAPI(activity.closed_at).replace('Z', '').split('T');

        this.restaurantBool = {};
        if (row.restaurants) {
          this.restaurants.map((restaurant) => {
            if (row.restaurants.some(item => item.id === restaurant.id)) {
              this.restaurantBool[restaurant.id] = true;
            } else {
              this.restaurantBool[restaurant.id] = false;
            }
          });
        } else {
          this.restaurants.map((restaurant) => {
            this.restaurantBool[restaurant.id] = false;
          });
        }

        this.dialog.header = 'Update event';
        this.dialog.action = 'Update';
        this.activity.id = activity.id;
        this.activity.name = activity.name;
        this.date = time[0];
        this.time = time[1];
        this.$refs[ref].open();
      },
      closeModal(ref) {
        this.$refs[ref].close();
      },
      updateActivity(ref) {
        if (this.activity.name === '' || this.date === '' || this.time === '') {
          return;
        }
        this.activity.closed_at = utils.formatDateForAPI(`${this.date} ${this.time}`, true);
        this.$refs[ref].close();
        this.$store.dispatch('updateActivity', { activity: this.activity, restaurantBool: this.restaurantBool });
      },
      formatDate(date) {
        return utils.formatDate(date);
      },
    },
  };
</script>

<style>
.text-left {
  text-align: left;
}

.manage-events-dialog{
  min-width: 600px;
}
.restaurant-option {
  display: flex;
}
</style>
