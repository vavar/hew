<template>
  <div class="page content" v-if="!isLoading">
    <md-dialog md-open-from="#fab" md-close-to="#fab" ref="activityModal">
      <md-dialog-title>{{dialog.header}}</md-dialog-title>

      <md-dialog-content>
        <form>
          <md-input-container>
            <label>Name</label>
            <md-input required v-model="activity.name" placeholder="Event Name"></md-input>
          </md-input-container>
          <md-input-container>
            <label>Closed At</label>
            <md-input required type="date" v-model="date"></md-input>
            <md-input required type="time" v-model="time"></md-input>
          </md-input-container>
        </form>
      </md-dialog-content>

      <md-dialog-actions>
        <md-button class="md-primary" @click="updateActivity('activityModal')">{{dialog.action}}</md-button>
        <md-button class="md-primary" @click="closeModal('activityModal')">Cancel</md-button>
      </md-dialog-actions>
    </md-dialog>

    <md-table-card>
      <md-toolbar>
        <h1 class="md-title">Manage Open Events</h1>
        <md-button class="md-icon-button" id="custom" @click="openModal('activityModal')">
          <md-icon>add</md-icon>
        </md-button>
      </md-toolbar>

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
              <md-button class="md-icon-button" @click="openEditModal('activityModal', rowIndex)">
                <md-icon>edit</md-icon>
              </md-button>
            </md-table-cell>
          </md-table-row>
        </md-table-body>
      </md-table>
    </md-table-card>

    <md-table-card style="margin-top: 20px">
      <md-toolbar>
        <h1 class="md-title">Past Events</h1>
      </md-toolbar>

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
            <md-table-cell>
            </md-table-cell>
          </md-table-row>
        </md-table-body>
      </md-table>
    </md-table-card>
  </div>
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
    }
  },
  created() {
    this.$store.dispatch('getActivities', this.activity.organization_id);
  },
  methods: {
    openModal(ref) {
      const localTime = utils.getLocalTime().replace('Z', '').split('T');
      this.dialog.header = 'Create a new event';
      this.dialog.action = 'Add';
      this.activity.id = undefined;
      this.activity.name = '';
      this.date = localTime[0];
      this.time = localTime[1];
      this.$refs[ref].open();
    },
    openEditModal(ref, rowIndex) {
      const activity = this.openActivities[rowIndex];
      const time = utils.formatDateForAPI(activity.closed_at).replace('Z', '').split('T');
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
      this.activity.closed_at = utils.formatDateForAPI(`${this.date} ${this.time}`, true);
      this.$refs[ref].close();
      this.$store.dispatch('updateActivity', this.activity);
    },
    formatDate(date) {
      return utils.formatDate(date);
    },
  },
};
</script>
