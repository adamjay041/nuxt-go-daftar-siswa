<template>
<div>

  <b-modal
      id="modal-prevent-closing"
      ref="modal"
      hide-footer
      :title="titleM"
      @show="resetModal"
    >

    <b-overlay
    :show="busy" rounded="lg" opacity="0.6" @hidden="onHidden"
    >
      <template #overlay>
        <div class="d-flex align-items-center">
          <b-spinner small type="grow" variant="secondary"></b-spinner>
          <b-spinner type="grow" variant="dark"></b-spinner>
          <b-spinner small type="grow" variant="secondary"></b-spinner>
          <!-- We add an SR only text for screen readers -->
          <span class="sr-only">Please wait...</span>
        </div>
      </template>

      <!-- <input v-model="formSubmit.first_name"> -->
        <form ref="form" @submit.stop.prevent="handleSubmit">
          <b-form-group
            label="First Name"
            label-for="First_name-input"
            invalid-feedback="First_Name is required"
            :state="first_name"
            :aria-hidden="busy ? 'true' : null"
          >
            <b-form-input
              id="First_name-input"
              v-model="formSubmit.First_Name"
              :state="first_name"
              required
            ></b-form-input>
          </b-form-group>
          <b-form-group
            label="Last Name"
            label-for="Last_name-input"
            invalid-feedback="Last_Name is required"
            :state="last_name"
          >
            <b-form-input
              id="Last_name-input"
              v-model="formSubmit.Last_Name"
              :state="last_name"
              vale
              required
            ></b-form-input>
          </b-form-group>
          <b-form-group
            label="Full Name"
            label-for="Full_name-input"
            invalid-feedback="Full_Name is required"
            :state="full_name"
          >
            <b-form-input
              readonly
              id="name-input"
              v-model="formSubmit.Full_Name"
              :state="full_name"
              required
            ></b-form-input>
          </b-form-group>
          <b-form-group
            label="Kelas"
            label-for="Kelas-input"
            invalid-feedback="Kelas is required"
            :state="kelas"
          >
            <b-form-select
              id="Kelas-input"
              v-model="formSubmit.Class"
              :state="kelas"
              :options="options"
              required
            ></b-form-select>
          </b-form-group>
          <b-form-group
            label="NPM"
            label-for="npm input"
            invalid-feedback="npm is required"
            :state="npm"
          >
            <b-form-input
              id="npm-input"
              v-model="formSubmit.NPM"
              :state="npm"
              required
              @blur="validateNpm"
            ></b-form-input>
            <!-- <span class="floating-placeholder" v-if="msg.email">{{msg.email}}</span> -->
          </b-form-group>
          <b-form-group
            label="Email"
            label-for="email-input"
            invalid-feedback="email is required"
            :state="email"
          >
            <b-form-input
              id="email-input"
              type="email"
              v-model="formSubmit.Email"
              :state="email"
              required
            ></b-form-input>
          </b-form-group>
          <b-form-group
            label="password"
            label-for="password-input"
            invalid-feedback="password is required"
          >
            <b-form-input
              id="password-input"
              v-model="formSubmit.Password"
              required
              type="password"
            ></b-form-input>
          </b-form-group>

      </form>
      <b-button class="mt-2" variant="outline-primary" block @click="onClick" :disabled="busy">Submit</b-button>
    </b-overlay>

  </b-modal>
</div>
</template>



<script>
export default {
  props: ['dataEdit','id'],
  data() {
    return {
      timeout: null,
      msg: [],
      options: [
        {value: "S3N", text: "S3N"},
        {value: "S2N", text: "S2N"},
        {value: "S1N", text: "S1N"}
      ],
      formSubmit: {
        First_Name: "",
        Last_Name: "",
        Full_Name: "",
        Class: "",
        NPM : "",
        Email: "",
        Password: ""
      },
      title : "",
      // formState: {
        first_name: null,
        last_name: null,
        full_name: null,
        kelas: null,
        npm: null,
        email: null,
      // },
      submittedNames: []
    }
  },
  methods: {
    checkFormValidity() {
      const valid = this.$refs.form.checkValidity()
      console.log("ini valid", valid)
      this.first_name = valid
      this.last_name = valid
      this.full_name = valid
      this.kelas = valid
      this.npm = valid
      this.email = valid
      return valid
    },
    resetModal() {
      if(this.titleM == "Edit Data Mahasiswa") {
        this.formSubmit = {
         First_Name: this.dataEdit['First_Name'],
         Last_Name: this.dataEdit['Last_Name'],
         Full_Name: this.dataEdit['Full_Name'],
         Class: this.dataEdit['Class'],
         NPM :this.dataEdit['NPM'],
         Email: this.dataEdit['Email']
       }
      } else {
        this.formSubmit =  {
          First_Name: "",
          Last_Name: "",
          Full_Name: "",
          Class: "",
          NPM :"",
          Email: ""
        }
      }
    },
    handleSubmit() {
      // Exit when the form isn't valid
      if (!this.checkFormValidity()) {
        return false
      }
      if(this.titleM == "Tambah Daftar Mahasiswa") {
        this.$store.dispatch("siswa/postSiswa", this.formSubmit)
        this.$nextTick(() => {
          this.$bvModal.hide('modal-prevent-closing')
        })
      }
      else if (this.titleM == "Edit Data Mahasiswa") {
        var obj = {
          "id": this.id,
          "data": this.formSubmit
        }
        this.$store.dispatch("siswa/putSiswa", obj)
        this.$nextTick(() => {
          this.$bvModal.hide('modal-prevent-closing')
        })
      }
      else {
        this.$store.dispatch("siswa/createSiswa", this.formSubmit)

        this.$nextTick(() => {
          this.$bvModal.hide('modal-prevent-closing')
        })
      }
    },
    onClick(e) {
        e.preventDefault()
        // alert(JSON.stringify(this.form))
      // this.busy = true
      // Simulate an async request
        this.handleSubmit()
    },
    validateNpm(e) {
      if(this.formSubmit.NPM.length < 14) {
        console.log('ini diaaa')
        // msg['email'] = "panjang npm harus kurang dari 14"
      }
    },
    onHidden() {
      // Return focus to the button
      // this.$refs.button.focus()
    }
  },
  computed: {
    titleM() {
      return this.$store.state.siswa.titleModal
    },
    busy() {
      return this.$store.state.siswa.isBusy
    },
    validation(val) {
      return this.formSubmit[val] != undefined || this.formSubmit[val] == ""
    }
  },
  watch:  {
    'formSubmit.First_Name'(newF, oldF) {
      this.formSubmit.Full_Name = newF
    },
    'formSubmit.Last_Name'(newF, oldF) {
      this.formSubmit.Full_Name = this.formSubmit.First_Name + " " + newF
    }
  }
}
</script>

<style>

</style>
