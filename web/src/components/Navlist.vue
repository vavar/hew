<template>
  <div class="phone-viewport">
    <md-list>
      <md-list-item>
        <md-icon>home</md-icon> <router-link to="/"><span>Home</span></router-link>
      </md-list-item>

      <md-list-item v-if="$auth.check()">
        <md-icon>restaurant</md-icon> <router-link to="/restaurants"><span>Restaurants</span></router-link>
      </md-list-item>

      <md-list-item v-if="$auth.check('admin')">
        <md-icon>build</md-icon> <span>Admin Tools</span>
        <md-list-expand>
          <md-list>
            <md-list-item>
              <span>Manage Restaurants</span>
            </md-list-item>
            <md-list-item>
              <router-link to="/admin/events"><span>Manage Events</span></router-link>
            </md-list-item>
            <md-list-item>
              <span>Order Reports</span>
            </md-list-item>
          </md-list>
        </md-list-expand>
      </md-list-item>
  </div>
</template>

<script>
  export default {
    name: 'navlist',
    methods: {},
    computed: {
      isAdmin() {
        return true;
      },
    },
    components: {},
  };
</script>
