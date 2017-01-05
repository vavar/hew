<template>
  <div class="page">
    <md-card>
      <md-card-content>
        <h2>Login</h2>
        <div v-show="error" style="color:red; word-wrap:break-word;">{{ error }}</div>
        <p v-if="$route.query.redirect">
          You need to login first.
        </p>
        <form @submit.prevent="login">
          <p v-if="error" class="error">Bad login information</p>
          <md-input-container>
            <label>Email</label>
            <md-input name="username" required v-model="data.body.username" placeholder="email"></md-input>
          </md-input-container>
          <md-input-container class="md-input-invalid">
            <label>Password</label>
            <md-input name="password" required v-model="data.body.password" placeholder="password" type="password"></md-input>
          </md-input-container>
          <md-button class="md-primary md-raised" type="submit">log in</md-button>
          <router-link tag="md-button" class="md-primary md-raised" to="/register">sign up</router-link>
        </form>
      </md-card-content>
    </md-card>
  </div>
</template>

<script>

export default {
  data() {
    return {
      data: {
        body: {
          username: '',
          password: '',
        },
      },
      error: null,
    };
  },
  methods: {
    login() {
      this.$auth.login({
        body: this.data.body,
        redirect: {name: 'home'},
        success(res) {
          const user = res.data.data;
          this.$auth.user(user);
        },
        error(res) {
          this.error = res.data;
        }
      });
    }
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
