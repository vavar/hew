<template>
  <div class="page">
    <md-card>
      <md-card-content>
        <h2>Login</h2>
        <p v-if="$route.query.redirect">
          You need to login first.
        </p>
        <form novalidate @submit.prevent="login">
          <p v-if="error" class="error">Bad login information</p>
          <md-input-container>
            <label>Email</label>
            <md-input required v-model="email" placeholder="email"></md-input>
          </md-input-container>
          <md-input-container class="md-input-invalid">
            <label>Password</label>
            <md-input required v-model="pass" placeholder="password" type="password"></md-input>
            <span class="md-error">Hint: pass</span>
          </md-input-container>
          <md-button class="md-primary md-raised" type="submit">login</md-button>
        </form>
      </md-card-content>
    </md-card>  
  </div>
</template>

<script>
import auth from '../auth';

export default {
  data() {
    return {
      email: 'joe@example.com',
      pass: '',
      error: false,
    };
  },
  methods: {
    login() {
      auth.login(this.email, this.pass, (loggedIn) => {
        if (!loggedIn) {
          this.error = true;
        } else {
          this.$router.replace(this.$route.query.redirect || '/');
        }
      });
    },
  },
};
</script>

<style>
.error {
  color: red;
}
.page {
  width: 960px;
  margin: 0 auto;
}
</style>
