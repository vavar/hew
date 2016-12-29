<template>
  <div class="page content">
    <md-dialog md-open-from="#fab" md-close-to="#fab" ref="activityModal">
      <md-dialog-title>{{dialog.header}}</md-dialog-title>

      <md-dialog-content>
        <form>
          <md-input-container>
            <label>Name</label>
            <md-input required v-model="activity.name" placeholder="Restaurant Name"></md-input>
          </md-input-container>
        </form>
      </md-dialog-content>

      <md-dialog-actions>
        <md-button class="md-primary">{{dialog.action}}</md-button>
        <md-button class="md-primary" @click="closeModal('activityModal')">Cancel</md-button>
      </md-dialog-actions>
    </md-dialog>

    <md-table-card>
      <md-toolbar>
        <h1 class="md-title">Manage Events</h1>
        <md-button class="md-icon-button" id="custom" @click="openModal('activityModal')">
          <md-icon>add</md-icon>
        </md-button>
      </md-toolbar>

      <md-table md-sort="closed-at" md-sort-type="desc">
        <md-table-header>
          <md-table-row>
            <md-table-head md-sort-by="name">Name</md-table-head>
            <md-table-head md-sort-by="closed-at">Closed At</md-table-head>
            <md-table-head>&nbsp;</md-table-head>
          </md-table-row>
        </md-table-header>

        <md-table-body>
          <md-table-row v-for="(row, rowIndex) in activities" :key="rowIndex" :md-item="row" md-auto-select>
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
    </md-table-card>
  </div>
</template>

<script>
export default {
  name: 'manage-events',
  data: () => ({
    dialog: { header: '', action: '' },
    activity: { name: '', closed_at: '' },
  }),
  methods: {
    openModal(ref) {
      this.dialog.header = 'Create a new event';
      this.dialog.action = 'Add';
      this.$refs[ref].open();
    },
    closeModal(ref) {
      this.$refs[ref].close();
    },
  },
};
</script>
