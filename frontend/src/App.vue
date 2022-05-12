<template>
  <div class="corpo">
    <router-view></router-view>

  </div>
</template>

<script>
export default {
  name: `App`
}
</script>

<style>
.corpo {
  margin: 0;
  padding: 0;
  min-height: 100vh;
  min-width: 100vw;
}

</style>
