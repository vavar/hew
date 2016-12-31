<template>
  <div class="page">
    <md-card>
      <md-card-content>
        <h2>Register your new account</h2>
        <form novalidate @submit.prevent="register">
          <md-input-container :class="{'md-input-invalid': errors.has('email') }">
            <label>Email</label>
            <md-input v-validate data-vv-name="email" data-vv-rules="required|email" data-vv-delay="1000" required v-model="data.body.email" placeholder="email"></md-input>
            <span v-show="errors.has('email')" class="md-error">{{ errors.first('email') }}</span>
          </md-input-container>
          <md-input-container :class="{'md-input-invalid': errors.has('username') }">
            <label>Username</label>
            <md-input v-validate data-vv-name="username" data-vv-rules="required|min:3" data-vv-delay="1000" required v-model="data.body.username" placeholder="username"></md-input>
            <span v-show="errors.has('email')" class="md-error">{{ errors.first('username') }}</span>
          </md-input-container>
          <md-input-container :class="{'md-input-invalid': errors.has('password') }">
            <label>Password</label>
            <md-input v-validate data-vv-name="password" data-vv-rules="required|min:4" data-vv-delay="1000" required v-model="data.body.password" placeholder="password" type="password"></md-input>
            <span v-show="errors.has('password')" class="md-error">{{ errors.first('password') }}</span>
          </md-input-container>
          <md-button class="md-primary md-raised" type="submit">Submit</md-button>
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
          email: '',
          username: '',
          password: '',
          organization_id: 1,
          role: 'user',
        },
      error: null,
      }
    };
  },
  methods: {
    register() {
      this.$auth.register({
        body: this.data.body,
        success: function () {
          this.$router.replace(this.$route.query.redirect || '/login');
        },
        error: function (res) {
          this.error = res.data;
        }
      });
    },
  },
};
</script>
