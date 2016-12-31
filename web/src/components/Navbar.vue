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
      <span class="md-title-gutter"></span>
      <div v-if="$auth.check()" class="profile-panel">
        <md-avatar>
          <img src="http://lorempixel.com/128/128/people">
        </md-avatar>
        <ul>
          <li><span>{{$auth.user().email}}</span></li>
          <li><a @click="onLogout()" class="md-raised md-accent">Log out</a></li>
        </ul>
      </div>
      <div v-else>
        <md-button @click="toLogin()" class="md-raised md-warn">Sign in</md-button>
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

  export default {
    name: 'navbar',
    data() {
      return {
        loggedIn: true,
      };
    },
    created() {
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
      toLogin() {
        this.$router.push({ path: '/login' });
      },
      onLogout() {
        this.$auth.logout({
          makeRequest: true,
          params: {},
          success: function () { },
          error: function () { },
          redirect: '/login',
        });
      }
    },
    components: {
      Navlist,
    },
  };
</script>

<style>
  .profile-panel {
    width: 140px;
  }
  .profile-panel .md-avatar {
    float:left;
    margin-right: 10px;
  }
  .profile-panel a {
    cursor: pointer;
  }
  .profile-panel ul {
    list-style: none;
    margin:0px;
    padding:0px;
    margin-left: 10px;
  }
  .profile-panel li {
    text-align: left;
  }
  .text-right {
    position: absolute;
    right: 8px;
    top: 7px;
  }
</style>
