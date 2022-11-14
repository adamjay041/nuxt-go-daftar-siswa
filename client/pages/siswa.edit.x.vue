<template>
  <b-modal id="mahasiswa-edit" hide-footer centered @hidden="onModalHidden" @show="resetModal">
    <template #modal-title> Edit Mahasiswa</template>
    <alert-invalid-form v-if="JSON.stringify(errors) != '{}'" />
    <div class="d-block">
      <div class="row">
        <div class="col-lg-12">
          <div class="row">
            <div class="col-lg-12">
              <div class="form-group">
                <label class="d-block">Major Study</label>
                <div>
                  <multiselect
                    :options="dropdown.PositionLevel"
                    aria-invalid="true"
                    :close-on-select="true"
                    :clear-on-select="false"
                    :preserve-search="true"
                    open-direction="bottom"
                    placeholder="Select Level"
                    :class="
                      errors.major
                        ? 'is-invalid'
                        : !model.major
                        ? null
                        : 'is-valid'
                    "
                   v-model="model.major"
                  />
                </div>
                <div
                  class="invalid-feedback d-block"
                  v-if="errors.major"
                >
                  {{ errors.major[0] }}
                </div>
              </div>
            </div>
          </div>
          <div class="row">
            <div class="col-lg-12">
              <div class="form-group">
                <label class="d-block">Name Student</label>
                <div>
                  <input
                    type="text"
                    aria-invalid="true"
                    class="form-control"

                    :class="
                      errors.name
                        ? 'is-invalid'
                        : !model.name
                        ? null
                        : 'is-valid'
                    "
                   v-model="model.name"
                  />
                </div>
                <div
                  class="invalid-feedback d-block"
                  v-if="errors.name"
                >
                  {{ errors.name[0] }}
                </div>
              </div>
            </div>
          </div>

          <div class="row">
            <div class="col-lg-6">
              <div class="form-group">
                <label class="d-block">Link Picture</label>
                <div>
                  <input
                    type="text"

                    aria-invalid="true"
                    class="form-control"
                    :disabled="model.PositionLevel == 'Fresh Grad/Less than 1 Year Experience'"

                    :class="
                      errors.l_pict
                        ? 'is-invalid'
                        : !model.l_pict
                        ? null
                        : 'is-valid'
                    "
                   v-model="model.l_pict"
                  />
                </div>
                <div
                  class="invalid-feedback d-block"
                  v-if="errors.l_pict"
                >
                  {{ errors.l_pict[0] }}
                </div>
              </div>
            </div>
            <div class="col-lg-6">
              <div class="form-group">
                <label class="d-block">Kelas</label>
                <div>
                  <multiselect
                    :options="dropdown.EmploymentStatus"
                    aria-invalid="true"
                    :close-on-select="true"
                    :clear-on-select="false"
                    :preserve-search="true"
                    open-direction="bottom"
                    placeholder="Select Status"


                    :class="
                      errors.Kelas
                        ? 'is-invalid'
                        : !model.Kelas
                        ? null
                        : 'is-valid'
                    "
                   v-model="model.Kelas"

                  />
                </div>
                <div
                  class="invalid-feedback d-block"
                  v-if="errors.Kelas"
                >
                  {{ errors.Kelas[0] }}
                </div>
              </div>
            </div>
          </div>

          <div class="row">
            <div class="col-lg-6">
              <div class="form-group">
                <label class="d-block">Email</label>
                <div>
                  <input
                    type="text"
                    aria-invalid="true"
                    class="form-control"
                    readonly
                    :class="
                      errors.email
                        ? 'is-invalid'
                        : !model.email
                        ? null
                        : 'is-valid'
                    "
                   v-model="model.email"
                  />
                </div>
                <div
                  class="invalid-feedback d-block"
                  v-if="errors.email"
                >
                  {{ errors.email[0] }}
                </div>
              </div>
            </div>
            <div class="col-lg-6">
              <div class="form-group">
                <label class="d-block">Password</label>
                <div>
                  <input
                    required
                    type="password"
                    aria-invalid="true"
                    class="form-control"


                    :class="
                      errors.password
                        ? 'is-invalid'
                        : !model.password
                        ? null
                        : 'is-valid'
                    "
                   v-model="model.password"
                  />
                </div>
                <div
                  class="invalid-feedback d-block"
                  v-if="errors.password"
                >
                  {{ errors.password[0] }}
                </div>
              </div>
            </div>
          </div>

        </div>
      </div>
    </div>
    <b-button class="mt-3" pill variant="primary" block @click="submit">
      Save
    </b-button>
    <loading v-if="isLoading" />
  </b-modal>
</template>

<script>
import moment from "moment";
export default {
  props: ["refresh","datas"],
  data: () => ({
    // isLoading: false,
    ModelCode: "RCCandidateExperiences",
    dropdown: {
      PositionLevel: [
        "Sistem Informatika (IT)",
        "Teknik Informatika (SI)",
      ],
      EmploymentStatus: ["S3N", "S3M", "S3O"],
      Role: ["admin", "student"]
    },
    model:
    {
      "name": "",
      // "npm" :"",
      "password": "",
      // "periode": "",
      "Kelas" : "",
      "l_pict" : "",
      "email" : "",
      // "status" : "",
      // "role" : "",
      "major" : "",
    },
    errors: {},
    mandatoryFields: [
      // {
      //   key: "name",
      //   msg: "name"
      // },
      // {
      //   key: "periode",
      //   msg: "periode"
      // },
      // {
      //   key: "kelas",
      //   msg: "kelas"
      // },
      // {
      //   key: "email",
      //   msg: "email"
      // },
      // {
      //   key: "role",
      //   msg: "role"
      // },
      // {
      //   key: "major",
      //   msg: "major"
      // },
    ],
  }),
  computed: {
    userData() {
      return this.$store.state.user.dataUser;
    },
    isLoading() {
      return this.$store.state.isBusy;
    }
  },
  methods: {
  objCopy(pObject) {
    return JSON.parse(JSON.stringify(pObject));
  },
    getToday(check) {
      if (check == true) {
        var days = new Date();
        this.model.EndDate = moment(days);
        this.model.EndInformation = moment(this.model.EndDate).format(
          "YYYY/MM/DD"
        );
      } else {
        this.model.EndDate = null;
        this.model.EndInformation = "YYYY/MM/DD";
      }
    },
    onlyNumber($event) {
      let keyCode = $event.keyCode ? $event.keyCode : $event.which;
      if ((keyCode < 48 || keyCode > 57) && keyCode !== 46) {
        $event.preventDefault();
      }
    },
    submit: function () {
      const data = this.objCopy(this.model);
      let errMsg = [];
        // console.log(this.model.PeriodInformation[0])

      // if (this.model.PositionLevel != "Fresh Grad/Less than 1 Year Experience") {
        for (let i in data) {
          this.mandatoryFields.forEach((el) => {
            if (el.key === i) {
              if (this.model[i] === "" || this.model[i] === null) {
                errMsg.push(`${el.msg} tidak boleh kosong`);
              }
            }
          });
        }
      // }

      if (errMsg.length > 0) {
        toastDanger(errMsg[0]);
        return;
      } else {
        this.isLoading = true;
        // console.log(data)
        // data.npm = `${data.periode.split("-")[0]}${data.periode.split("-")[1]}`
        // data.periode = moment(data.periode).format("YYYY-MM-DD")
        // if (this.model.PositionLevel != "Fresh Grad/Less than 1 Year Experience") {
        //   let startDate = moment(this.model.PeriodInformation[0]).format("MM/YYYY");
        //   let endDate = moment(this.model.PeriodInformation[1]).format("MM/YYYY");


        //   model.PeriodInformation = `${startDate} - ${endDate}`;
        //   model.StartInformation = startDate;
        //   model.EndInformation = endDate;
        // }

        // model.CandidateCode = this.userData.Code;
        // model.Name = this.userData.Name;

        // this.$store.commit("siswa/setFormSubmitDefault")
        // this.$store.commit("siswa/setTitleModal", val)
        var obj = {
          "id": this.datas.id,
          "data": data
        }

        this.$store.dispatch("siswa/putSiswa", obj)
        // .then((r) => {
        //   console.log(r)
            this.$bvModal.hide("mahasiswa-add");
        //     toastSuccess("Daftarkan Siswa successful!");

        //   })
        //   .catch((err) => {
        //     this.isLoading = false;
        //     if (err.response) this.errors = err.response.data.Errors;
        //   })
        //   .finally(() => (this.isLoading = false));

        // this.$api
        //   .submitModel(this.ModelCode, model)
          // .then((r) => {
          //   this.$bvModal.hide("experience-add");
          //   toastSuccess("Add experience record successful!");
          //   // this.$router.go(0);
          // })
          // .catch((err) => {
          //   this.isLoading = false;
          //   if (err.response) this.errors = err.response.data.Errors;
          // })
          // .finally(() => (this.isLoading = false));
      }
    },
    onModalHidden: function () {
      this.model = {
        StartDate: null,
        EndDate: null,
        StartInformation: "MM/YYYY",
        EndInformation: "MM/YYYY",
      };
      this.errors = {};
      if (this["refresh"]) this["refresh"]();
    },
    resetModal: function() {
      this.model = {
        "name": this.datas.name,
        // "npm" :"",
        "password": "",
        // "periode": "",
        "Kelas" : this.datas.kelas,
        "l_pict" : this.datas.l_pict,
        "email" : this.datas.email,
        // "status" : "",
        // "role" : "",
        "major" : this.datas.major,
      }
    }
  },
};
</script>
