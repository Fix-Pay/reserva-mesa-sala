<template>
  <form @submit.prevent="handleSubmit">
    <h3> Sign Up </h3>

    <div class="form-group">
      <label> Name </label>
      <input type="text" class="form-control" v-model="first_name" placeholder="First Name"/>
    </div>

    <div class="form-group">
      <label> E-Mail </label>
      <input type="email" class="form-control" v-model="email" placeholder="Email"/>
    </div>

    <div class="form-group">
      <label> Password </label>
      <input type="password" class="form-control" v-model="password" placeholder="Password"/>
    </div>

    <div class="form-group">
      <label> Confirm Password </label>
      <input type="password" class="form-control" v-model="password_confirm" placeholder="Confirm Password"/>
    </div>

    <div class="form-group">
      <label> Occupation </label>
      <select v-model="occupation">
        <option value="" disabled selected hidden>Select your ccupation</option>
        <option value="senior">Senior</option>
        <option value="pleno">Pleno</option>
        <option value="junior">Junior</option>
        <option value="trainee">Trainee</option>
      </select>
    </div>

    <button class="btn btn-primary btn-block"> Sign Up </button>
  </form>
</template>

<script>
import axios from 'axios';
export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: "Register",
  data() {
    return {
      first_name: '',
      email: '',
      password: '',
      password_confirm: '',
      occupation: ''
    }
  },
  methods: {
    handleSubmit() {
      const data = {
        first_name: this.first_name,
        email: this.email,
        password: this.password,
        password_confirm: this.password_confirm,
        occupation: this.occupation
      };
      axios.post('', data) // URL da API
        .then(
          res => {
            console.log(res)
          }
        ).catch(
          err => {
            console.log(err)
          }
        )
    }
  }
}
</script>

<style scoped>

</style>
