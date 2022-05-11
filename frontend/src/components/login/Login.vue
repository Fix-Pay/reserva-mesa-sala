<template>
  <div class="container-login">
    <div class="img-container">
      <img src="./2.png" alt="Avatar" class="avatar">
    </div>

    <div class="container">
      <input type="text" placeholder="Username" name="uname" required>
      <input type="password" placeholder="Password" name="psw" required>
      <button type="submit">Login</button>
    </div>

    <div class="container-img">
      <img src="./6.png">
    </div>

  </div>
</template>

<script>
import Header from "../header/Header";
import axios from "axios";

export default {
  name: 'Login',
  components: {Header},
  methods: {
    getUser() {
      axios
        .get('localhost:8000/users')
        .then(( res ) => {
          this.user = res.data;
        })
        .catch((error) => {
          console.log(error);
        });
    },
  },
};
</script>

<style scoped>
input, button, span {
  font-family: "Montserrat", sans-serif;
}

input[type=text], input[type=password] {
  width: 90%;
  padding: 12px 20px;
  margin: 5px 0;
  display: flex;
  align-items: center;
  border-radius: 4px;
  border: 1px solid white;
  border-bottom: 1px solid gray;
}

button {
  background-color: #4FB8CE;
  border: 1px solid white;
  font-weight: bold;
  color: white;
  padding: 10px;
  margin: 20px;
  border-radius: 20px;
}

button:hover {
  opacity: 0.9;
}

.container-login {
  width: 100%;
  justify-content: center;
}

img {
  width: 100%;
  height: 160%;
}

.img-container {
  text-align: center;
  margin: 20px -50px -35px 0;
}

.avatar {
  width: 200px;
}

.container {
  padding: 16px;
  display: grid;
  align-items: center;
  justify-content: center;
}
</style>
