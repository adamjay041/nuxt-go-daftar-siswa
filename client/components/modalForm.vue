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
            :state="formState.first_name"
            :aria-hidden="busy ? 'true' : null"
          >
            <b-form-input
              id="First_name-input"
              v-model="formSubmit.First_Name"
              :state="formState.first_name"
              required
            ></b-form-input>
          </b-form-group>
          <b-form-group
            label="Last Name"
            label-for="Last_name-input"
            invalid-feedback="Last_Name is required"
            :state="formState.last_name"
          >
            <b-form-input
              id="Last_name-input"
              v-model="formSubmit.Last_Name"
              :state="formState.last_name"
              vale
              required
            ></b-form-input>
          </b-form-group>
          <b-form-group
            label="Full Name"
            label-for="Full_name-input"
            invalid-feedback="Full_Name is required"
            :state="formState.full_name"
          >
            <b-form-input
              readonly
              id="name-input"
              v-model="formSubmit.Full_Name"
              :state="formState.full_name"
              required
            ></b-form-input>
          </b-form-group>
          <b-form-group
            label="Kelas"
            label-for="Kelas-input"
            invalid-feedback="Kelas is required"
            :state="formState.kelas"
          >
            <b-form-input
              id="Kelas-input"
              v-model="formSubmit.Class"
              :state="formState.kelas"
              required
            ></b-form-input>
          </b-form-group>
          <b-form-group
            label="NPM"
            label-for="npm-input"
            invalid-feedback="npm is required"
            :state="formState.npm"
          >
            <b-form-input
              id="npm-input"
              v-model="formSubmit.NPM"
              :state="formState.npm"
              required
            ></b-form-input>
          </b-form-group>
          <b-form-group
            label="Email"
            label-for="email-input"
            invalid-feedback="email is required"
            :state="formState.email"
          >
            <b-form-input
            type="email"
              id="email-input"
              v-model="formSubmit.Email"
              :state="formState.email"
              required
            ></b-form-input>
          </b-form-group>
          <!-- <div v-if="titleM != 'Edit Data Mahasiswa'"> -->
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
          <!-- </div> -->

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
      formSubmit: {
        First_Name: "",
        Last_Name: "",
        Full_Name: "",
        Class: "",
        NPM :0,
        Email: "",
        Password: ""
      },
      title : "",
      formState: {
        first_nameState: null,
        last_nameState: null,
        full_nameState: null,
        kelasState: null,
        npmState: null,
        emailState: null
      },
      submittedNames: []
    }
  },
  methods: {
    checkFormValidity() {
      const valid = this.$refs.form.checkValidity()
      this.formState.first_nameState = valid
      this.formState.last_nameState = valid
      this.formState.full_nameState = valid
      this.formState.kelasState = valid
      this.formState.npmState = valid
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
          NPM :0,
          Email: ""
        }
      }
    },
    handleSubmit() {
      // Exit when the form isn't valid
      if (!this.checkFormValidity()) {
        return false
      }
      console.log("titel",this.titleM )
      // Push the name to submitted names
      // Hide the modal manually
      if(this.titleM == "Tambah Daftar Mahasiswa") {
        this.$store.dispatch("siswa/postSiswa", this.formSubmit)
        // if (this.busy == false) {
        this.$nextTick(() => {
          this.$bvModal.hide('modal-prevent-closing')
        })
        // }
        // this.$store.commit("siswa/setFormSubmitDefault")
      }
      else if (this.titleM == "Edit Data Mahasiswa") {
        var obj = {
          "id": this.id,
          "data": this.formSubmit
        }
        this.$store.dispatch("siswa/putSiswa", obj)
        // if (this.busy == false) {
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

      // this.setTimeout(() => {
      //   this.busy = false
      //   this.$store.commit("siswa/setFormSubmitDefault")


      // })
    },
    clearTimeout() {
      if (this.timeout) {
        clearTimeout(this.timeout)
        this.timeout = null
      }
    },
    setTimeout(callback) {
      this.clearTimeout()
      this.timeout = setTimeout(() => {
        this.clearTimeout()
        callback()
      }, 1000)
    },
    onHidden() {
      // Return focus to the button
      // this.$refs.button.focus()
    },
    onClick() {
      // this.busy = true
      // Simulate an async request
      this.handleSubmit()
    },
  },
  computed: {
    titleM() {
      return this.$store.state.siswa.titleModal
    },
    busy() {
      return this.$store.state.siswa.isBusy
    }
  },
  watch:  {
    'formSubmit.First_Name'(newF, oldF) {
      this.formSubmit.Full_Name = newF
    },
    'formSubmit.Last_Name'(newF, oldF) {
      this.formSubmit.Full_Name = this.formSubmit.First_Name + " " + newF
    }
  },
  beforeDestroy() {
    this.clearTimeout()
  },
}
</script>

<style>

</style>
