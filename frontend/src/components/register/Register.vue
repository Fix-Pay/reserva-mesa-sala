<template>
  <form @submit.prevent="handleSubmit">
    <h3> Login </h3>

    <div class="form-group">
      <label> Name </label>
      <input type="text" class="form-control" v-model="name" placeholder="Name"/>
    </div>

    <div class="form-group">
      <label> E-Mail </label>
      <input type="email" class="form-control" v-model="email" placeholder="Email"/>
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
      name: '',
      email: ''
    }
  },
  mounted() {
    this.$session.set("username", "user123");
  },
  methods: {
    handleSubmit() {
      const data = {
        name: this.name,
        email: this.email,
      };
      axios.post('http://localhost:8080/register', data);

      this.$router.push('/');
    }
  }
}
</script>

<style scoped>

</style>
