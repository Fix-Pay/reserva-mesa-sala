<template>
  <div class="container-login">
    <form method="post">
      <div class="imgcontainer">
        <img src="./2.png" alt="Avatar" class="avatar">
      </div>

      <div class="container">
        <input type="text" placeholder="Username" name="uname" required>
        <input type="password" placeholder="Password" name="psw" required>
        <button type="submit">Login</button>
      </div>
    </form>
  </div>
</template>

<script>
export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: "Login"
}
</script>

<style scoped>
form {
  width: 93%;

}
input, button, span {
  font-family: "Montserrat", sans-serif;
}

input[type=text], input[type=password] {
  width: 110%;
  padding: 12px 20px;
  margin: 5px 0;
  display: flex;
  align-items: center;
  border-radius: 4px;
}

button {
  background-color: #4FB8CE;
  color: white;
  padding: 10px;
  margin: 20px;
  border-radius: 20px;
}

button:hover {
  opacity: 0.8;
}

.container-login {
  width: 100%;
  justify-content: center;
}


.imgcontainer {
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

span.psw {
  float: right;
  padding-top: 16px;
}

@media screen and (max-width: 300px) {
  span.psw {
    display: block;
    float: none;
  }
}
</style>
