<template>
    <header>
        <div class="caixa">
            <h1>
                <slot class="titulo">
                </slot>
            </h1>
            <button class="login-btn">log-in</button>
        </div>
    </header>
</template>

<script>
export default {
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
  position: absolute;
  top: 10%;
  right: 0px;
  width: 10%;
  height: 26px;
  background-color: #52bacd;
  border: 1px solid #52bacd;
  border-radius: 6px;
  cursor: pointer;
}

</style>