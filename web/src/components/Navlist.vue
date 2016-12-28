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
        <md-icon>restaurant</md-icon> <router-link to="/restaurant"><span>Restaurant</span></router-link>
      </md-list-item>
  </div>
</template>

<script>
  export default {
    name: 'navlist',
    methods: {},
    components: {},
  };
</script>
