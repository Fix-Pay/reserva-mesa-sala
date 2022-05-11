<template>
    <header>
        <div class="caixa">
            <h1>
                <slot class="titulo">
                </slot>
            </h1>
            <button class="login-btn" @click="showLogin">Log In</button>
        </div>
    </header>
</template>

<script>
export default {
  methods: {
    showLogin() {
      this.$router.push('/login');
    }
  }
}
</script>

<style>
header {
  background: #ffff;
  border-bottom: #606566 solid 2px;
}
.caixa {
  position: relative;
  width: 980px;
  margin: 0 auto;
}
.img-header {
  width: 20%;
}
.login-btn {
  font-weight: bold;
  color:white;
  position: absolute;
  top: 10%;
  right: 0px;
  width: 12%;
  height: 40px;
  background-color: #4FB8CE;
  border: 1px solid white;
  border-radius: 6px;
  cursor: pointer;
}

.login-btn:hover {
  background: aqua;
  transition: width 2s, height 2s, background-color 2s, transform 2s;
}

</style>
