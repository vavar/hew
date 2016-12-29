<template>
  <div class="phone-viewport">
    <md-list>
      <md-list-item>
        <md-icon>home</md-icon> <router-link to="/"><span>Home</span></router-link>
      </md-list-item>

      <md-list-item>
        <md-icon>local_dining</md-icon> <router-link to="/order"><span>Make an order</span></router-link>
      </md-list-item>

      <md-list-item>
        <md-icon>restaurant</md-icon> <router-link to="/restaurants"><span>Restaurants</span></router-link>
      </md-list-item>

      <md-list-item v-if="isAdmin">
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
  import auth from '../auth';

  export default {
    name: 'navlist',
    methods: {},
    computed: {
      isAdmin() {
        return auth.isAdmin();
      },
    },
    components: {},
  };
</script>
