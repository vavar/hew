<template>
  <div class="phone-viewport">
    <md-list>
      <md-list-item>
        <md-icon>home</md-icon> <router-link to="/"><span>Home</span></router-link>
      </md-list-item>

      <md-list-item >
        <md-icon>find_in_page</md-icon> <router-link to="/history"><span>Order History</span></router-link>
      </md-list-item>

      <md-list-item >
        <md-icon>restaurant</md-icon> <router-link to="/restaurants"><span>Restaurants</span></router-link>
      </md-list-item>

      <md-list-item >
        <md-icon>build</md-icon> <router-link to="/admin/events"><span>Manage Events</span></router-link>
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
