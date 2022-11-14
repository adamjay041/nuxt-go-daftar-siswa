<template>
  <div>
    <b-navbar toggleable="sm" type="light" variant="white" class="bd-navbar" sticky>
      <!-- <b-navbar-toggle target="nav-text-collapse"></b-navbar-toggle> -->
      <b-navbar-brand>
        <img src="https://unindra.ac.id/assets/front/img/header_logo_16462082631435350658.png" alt="Kitten" >
      </b-navbar-brand>
      <b-collapse id="nav-text-collapse" is-nav>
        <b-navbar-nav>
          <b-nav-text>Daftar Mahasiswa</b-nav-text>
        </b-navbar-nav>
      </b-collapse>
      <b-nav-form>
        <b-button size="md" class="mr-1" variant="outline-danger" @click="logout">
          Logout  <font-awesome-icon :icon="['fas', 'arrow-right-from-bracket']" />
        </b-button>
      </b-nav-form>
    </b-navbar>
  </div>
</template>

<script>
export default {
  methods : {
    logout() {
      clearCookies()
      this.$router.push("/login")
    }
  }
}
</script>

<style>

</style>
