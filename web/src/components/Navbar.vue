<template>
  <div class="phone-viewport">
    <md-toolbar>
      <md-button class="md-icon-button" @click="toggleLeftSidenav">
        <md-icon>menu</md-icon>
      </md-button>
      <h2 class="md-title" @click="gotoHome()">
          <md-avatar>
            <img src="../assets/logo.png" />
          </md-avatar>
          <span>Hew</span>
      </h2>
      <div class="text-right">
        <label>{{ $auth.user().email }}</label>
        <router-link tag="md-button" v-if="$auth.ready()" v-on:click="$auth.logout()" class="md-raised md-accent">Log out</router-link>
        <router-link tag="md-button" v-if="!$auth.check()" to="/login" class="md-raised md-warn">Sign in</router-link>
      </div>
    </md-toolbar>

    <md-sidenav class="md-left" ref="leftSidenav" @open="open('Left')" @close="close('Left')">
      <md-toolbar>
        <div class="md-toolbar-container">
          <h3 class="md-title">Navigation</h3>
        </div>
      </md-toolbar>
      <navlist></navlist>
    </md-sidenav>
  </div>
</template>

<script>
  import Navlist from './Navlist';
  import auth from '../auth';

  export default {
    name: 'navbar',
    data() {
      return {
        loggedIn: auth.loggedIn(),
      };
    },
    created() {
      auth.onChange = (loggedIn) => {
        this.loggedIn = loggedIn;
      };
    },
    methods: {
      toggleLeftSidenav() {
        this.$refs.leftSidenav.toggle();
      },
      open() {
      },
      close() {
      },
      gotoHome() {
        this.$router.push({ path: '/' });
      },
      onLogout() {
        this.$auth.logout({
        makeRequest: true,
        params: {},
        success: function () {},
        error: function () {},
        redirect: '/login',
        // etc...
    });
      }
    },
    components: {
      Navlist,
    },
  };
</script>

<style>
  .text-right {
    position: absolute;
    right: 8px;
    top: 7px;
  }
</style>
