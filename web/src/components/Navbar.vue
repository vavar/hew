<template>
  <div class="phone-viewport">
    <md-toolbar>
      <md-button class="md-icon-button" @click="toggleLeftSidenav">
        <md-icon>menu</md-icon>
      </md-button>

      <h2 class="md-title">Hew App</h2>
    </md-toolbar>
    <md-sidenav class="md-left" ref="leftSidenav" @open="open('Left')" @close="close('Left')">
      <md-toolbar class="md-large">
        <div class="md-toolbar-container">
          <h3 class="md-title">Side Navigation</h3>
        </div>
      </md-toolbar>

      <navlist></navlist>
    </md-sidenav>
  </div>
</template>

<script>
  import Navlist from './Navlist';

  export default {
    name: 'navbar',
    methods: {
      toggleLeftSidenav() {
        this.$refs.leftSidenav.toggle();
      },
    },
    components: {
      Navlist,
    },
  };
</script>
