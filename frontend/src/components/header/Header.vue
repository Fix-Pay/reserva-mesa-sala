<template>
<<<<<<< HEAD
  <div>
    <b-nav pills>
      <b-nav-item href="/"><img class="img-header" src="../../assets/logo-black.svg" alt=""></b-nav-item>
      <b-nav-item href="/logs">Reservados</b-nav-item>
      <b-nav-item href="/#">Disponível</b-nav-item>
      <b-nav-item-dropdown
        id="my-nav-dropdown"
        text="Identificação"
        toggle-class="nav-link-custom"
        right
      >
        <b-item>
          <div class="form-group">
            <input type="text" class="form-control" v-model="name" placeholder="Name"/>
          </div>
        </b-item>
        <b-item>
          <div class="form-group">
            <input type="email" class="form-control" v-model="email" placeholder="Email"/>
          </div>
        </b-item>
        <b-dropdown-divider></b-dropdown-divider>
        <b-dropdown-item><b-button variant="outline-primary">Submit</b-button></b-dropdown-item>
      </b-nav-item-dropdown>
    </b-nav>
  </div>
=======
    <header>
        <div class="caixa">
            <h1>
                <slot class="titulo">
                </slot>
            </h1>

             <a href="/logs"><button class="login-btn">reservados</button></a>
        </div>
    </header>
>>>>>>> feature/front1
</template>

<script>
export default {
}
</script>

<style>
header {
  background: #ffff;
<<<<<<< HEAD
=======
  border-bottom: #606566 solid 2px;
>>>>>>> feature/front1
}
.caixa {
  position: relative;
  width: 980px;
  margin: 0 auto;
}
.img-header {
<<<<<<< HEAD
  width: 200px;
}
b-nav-item-dropdown {
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
.reservas:hover {
  background: aqua;
  transition: width 2s, height 2s, background-color 2s, transform 2s;
}

</style>
=======
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
>>>>>>> feature/front1
