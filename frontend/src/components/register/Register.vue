<template>
  <form>
    <h3> Sign Up </h3>

    <div class="form-group">
      <label> First Name </label>
      <input type="text" class="form-control" placeholder="First Name"/>
    </div>

    <div class="form-group">
      <label> Last Name </label>
      <input type="text" class="form-control" placeholder="Last Name"/>
    </div>

    <div class="form-group">
      <label> E-Mail </label>
      <input type="email" class="form-control" placeholder="Email"/>
    </div>

    <div class="form-group">
      <label> Password </label>
      <input type="password" class="form-control" placeholder="Password"/>
    </div>

    <div class="form-group">
      <label> Confirm Password </label>
      <input type="password" class="form-control" placeholder="Confirm Password"/>
    </div>

    <button class="btn btn-primary btn-block"> Sign Up </button>
  </form>
</template>

<script>
export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: "Register"
}
</script>

<style scoped>

</style>
