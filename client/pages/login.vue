<template>
<div class="content-center">
  <div class="auth-box">
    <div>
      <form>
        <div class="text-center">

          <img
            style="width: 110px; margin: 20px"
            src="https://lms.unindra.ac.id/media_public/get_gambar/d2llSWJaazRBRUtSZ2lFQm15V3RiZi9GV3oyd09zSGdTdXlOVjVyR0Zsc20yWkxNc3J4dUNFdi9MbENKd2ZYSg=="
            width="80"
            class="center"
          />
           <h3>Sign In</h3>
          <p class="text-muted">Enter your Email and password</p>
        </div>

        <b-alert show variant="danger" v-if="isAuthenticationFail">
          <strong>Authentication Failed!</strong> Invalid email or password.
        </b-alert>
        <div class="form-group py-3 m-0 has-success">
          <input
            required
            validated="true"
            v-model="model.email"
            type="text"
            name="email"
            class="form-control placeholder-dark-75"
            placeholder="Email"
          />
        </div>

        <div class="form-group py-3 m-0 has-invalid">
          <input
            required
            password="password"
            v-model="model.password"
            type="password"
            class="form-control placeholder-dark-75"
            placeholder="password"
            name="password"
          />
        </div>

        <button
          type="submit"
          class="btn btn-block btn-lg btn-primary"
          @click="signIn"
        >
          <a>Login</a>
        </button>

        <!-- <div
          class="form-group d-flex flex-wrap justify-content-between align-items-center mt-3"
        >
          <a class="text-muted text-hover-primary" href="/auth/reset-password"
            >Forgot Password?
          </a>
        </div> -->
        <!-- <div class="marginTop4-2BNfKC">
          <span class="needAccount-23l_Wh">Need an Account?</span>
          <a class="" v-b-modal.modal-prevent-closing @click="addSiswa('Buat Akun Mahasiswa')">Sign Up </a>
        </div> -->
        <div class="d-flex flex-column-auto justify-content-between mt-5">
          <div class="text-dark-50 font-weight-bold order-2 order-sm-1 my-2">
            © {{ new Date().getFullYear() }} Sistem Informasi
          </div>
          <div class="d-flex order-1 order-sm-2 my-2"></div>
        </div>
      </form>
      <loading v-if="isLoading" />
    </div>
  </div>
</div>
</template>

<script>
import modalForm from '../components/modalForm.vue'
export default {
  name: "Login",
  components: { modalForm },
  data() {
    return {
      isAuthenticationFail: false,
      isLoading: false,
      model: {
        email: "",
        password: "",
      },
      editData: '',
      id: '',
    }
  },
  methods: {
    signIn (e){
       e.preventDefault();
       this.isAuthenticationFail = false;
      try {
        this.$store.dispatch("siswa/loginSiswa", this.model)
         this.isAuthenticationFail = false;
       } catch (error) {
         this.isAuthenticationFail = true;
       }
      // this.isLoading = true
      // this.isLoading = false
    },
    addSiswa(val) {
      this.$store.commit("siswa/setFormSubmitDefault")
      this.$store.commit("siswa/setTitleModal", val)
    },
  },
}
</script>

<style lang="scss">


.content-center {
  height: 100%;
  display: flex;
  align-content: center;
  align-items: center;
  justify-content: center;
  background-image: url("https://lms.unindra.ac.id/assets/template/login_colek/images/img-wayang.png");
  background-repeat: no-repeat;
  background-size: cover;
  background-position: center;
}

#__nuxt {
  height: 100%;
}

#__layout {
  height: 100%;
}


.auth-box {
  width: 450px;
  // height: 100%;
  // margin: auto;
  // justify-content: center;
  // align-items: center;
  // height: 80vh;
  background: #fff;
  padding: 50px;
  border-radius: 12px;
  // margin-top: 100px;
  // padding-bottom: 25px;
}
.text-dark-50 {
  color: #7e8299!important;
}
</style>
