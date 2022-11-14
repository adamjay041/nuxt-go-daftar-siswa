<template>
  <b-modal id="mahasiswa-add" hide-footer centered @hidden="onModalHidden">
    <template #modal-title> Tambah Mahasiswa</template>
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
          <!-- <div class="row">
            <div class="col-lg-12">
              <div class="form-group">
                <label class="d-block">NPM</label>
                <div>
                  <input
                    type="text"
                    aria-invalid="true"
                    class="form-control"


                    :class="
                      errors.npm
                        ? 'is-invalid'
                        : !model.npm
                        ? null
                        : 'is-valid'
                    "
                    v-model="model.npm"
                  />
                </div>
                <div
                  class="invalid-feedback d-block"
                  v-if="errors.npm"
                >
                  {{ errors.npm[0] }}
                </div>
              </div>
            </div>
          </div> -->
          <div class="row">
            <div class="col-lg-6">
              <div class="form-group">
                <label class="d-block">Period</label>
                <div>
                  <date-picker
                    type="date"
                    format="YYYY-MM-DD"
                    :placeholder="model.periode"
                    :editable="false"
                    :default-value="new Date()"


                    v-model="model.periode"
                  />
                </div>
                <div class="invalid-feedback d-block" v-if="errors.periode">
                  {{ errors.periode[0] }}
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
          <!-- <div class="row">
            <div class="col-lg-6"> -->
              <!-- <div class="form-group">
                <label class="d-block">Time Period</label>
                <div>
                  <date-picker
                    type="date"
                    range
                    format="MM-YYYY"
                    placeholder="MM/YYYY"
                    :class="
                      errors.PeriodInformation
                        ? 'is-invalid'
                        : !model.PeriodInformation
                        ? null
                        : 'is-valid'
                    "
                    v-model="model.PeriodInformation"
                    :disabled="model.StillOnGoing == true || model.PositionLevel == 'Fresh Grad/Less than 1 Year Experience'"
                  />
                </div>
                <div
                  class="invalid-feedback d-block"
                  v-if="errors.PeriodInformation"
                >
                  {{ errors.PeriodInformation[0] }}
                </div>
              </div> -->
              <!-- <div class="form-group">
                <label class="d-block">Start Date</label>
                <div>
                  <date-picker
                    type="date"
                    format="MM-YYYY"
                    :placeholder="model.StartInformation"
                    :editable="false"
                    :default-value="new Date()"
                    :disabled="model.PositionLevel == 'Fresh Grad/Less than 1 Year Experience'"

                    v-model="model.StartDate"
                  />
                </div>
                <div class="invalid-feedback d-block" v-if="errors.StartDate">
                  {{ errors.StartDate[0] }}
                </div>
              </div>
            </div> -->
            <!-- <div class="col-lg-6" >
              <div class="form-group">
                <span class="d-inline d-flex justify-content-between">
                  <label>End Date</label>
                  <b-form-checkbox
                    :disabled="model.PositionLevel == 'Fresh Grad/Less than 1 Year Experience'"

                    id="StillOnGoing"
                    v-model="model.StillOnGoing"
                    name="StillOnGoing"
                    @input="getToday(model.StillOnGoing)"
                  >
                    Untill Now
                  </b-form-checkbox>
                </span>
                <div >
                  <date-picker
                    type="date"
                    format="MM-YYYY"
                    :placeholder="model.EndInformation"
                    :editable="false"
                    :default-value="new Date()"
                    v-model="model.EndDate"

                    :disabled="model.StillOnGoing == true"
                  />
                </div>
                <div class="invalid-feedback d-block" v-if="errors.EndDate">
                  {{ errors.EndDate[0] }}
                </div>
              </div>
            </div> -->
          <!-- </div> -->
          <div class="row">
            <div class="col-lg-6">
              <div class="form-group">
                <label class="d-block">Email</label>
                <div>
                  <input
                    type="text"
                    aria-invalid="true"
                    class="form-control"

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
                <label class="d-block">Role</label>
                <div>
                  <multiselect
                    :options="dropdown.Role"
                    aria-invalid="true"
                    :close-on-select="true"
                    :clear-on-select="false"
                    :preserve-search="true"
                    open-direction="bottom"
                    placeholder="Select Status"


                    :class="
                      errors.kelas
                        ? 'is-invalid'
                        : !model.role
                        ? null
                        : 'is-valid'
                    "
                    v-model="model.role"

                  />
                </div>
                <div
                  class="invalid-feedback d-block"
                  v-if="errors.EmploymentStatus"
                >
                  {{ errors.EmploymentStatus[0] }}
                </div>
              </div>
            </div>
          </div>
          <!-- <div class="row">
            <div class="col-lg-6">
              <div class="form-group">
                <label class="d-block">Team Member Count</label>
                <div>
                  <input
                    type="text"
                    @keypress="onlyNumber($event)"
                    aria-invalid="true"
                    class="form-control"
                    :disabled="model.PositionLevel == 'Fresh Grad/Less than 1 Year Experience'"

                    :class="
                      errors.TeamMember
                        ? 'is-invalid'
                        : !model.TeamMember
                        ? null
                        : 'is-valid'
                    "
                    v-model="model.TeamMember"
                  />
                </div>
                <div class="invalid-feedback d-block" v-if="errors.TeamMember">
                  {{ errors.TeamMember[0] }}
                </div>
              </div>
            </div>
            <div class="col-lg-6">
              <div class="form-group">
                <label class="d-block">Salary</label>
                <div>
                  <input
                    type="text"
                    @keypress="onlyNumber($event)"
                    aria-invalid="true"
                    class="form-control"
                    :disabled="model.PositionLevel == 'Fresh Grad/Less than 1 Year Experience'"

                    :class="
                      errors.Salary
                        ? 'is-invalid'
                        : !model.Salary
                        ? null
                        : 'is-valid'
                    "
                    v-model="model.Salary"
                  />
                </div>
                <div class="invalid-feedback d-block" v-if="errors.Salary">
                  {{ errors.Salary[0] }}
                </div>
              </div>
            </div>
          </div>
          <div class="row">
            <div class="col-lg-12">
              <div class="form-group">
                <label class="d-block">Other Benefits</label>
                <div>
                  <textarea
                    type="text"
                    aria-invalid="true"
                    class="form-control"
                    :disabled="model.PositionLevel == 'Fresh Grad/Less than 1 Year Experience'"

                    :class="
                      errors.BenefitNotes
                        ? 'is-invalid'
                        : !model.BenefitNotes
                        ? null
                        : 'is-valid'
                    "
                    v-model="model.BenefitNotes"
                  />
                </div>
                <div
                  class="invalid-feedback d-block"
                  v-if="errors.BenefitNotes"
                >
                  {{ errors.BenefitNotes[0] }}
                </div>
              </div>
            </div>
          </div>
          <div class="row">
            <div class="col-lg-12">
              <div class="form-group">
                <label class="d-block">Responsibilities</label>
                <div>
                  <textarea
                    type="text"
                    aria-invalid="true"
                    class="form-control"
                    :disabled="model.PositionLevel == 'Fresh Grad/Less than 1 Year Experience'"

                    :class="
                      errors.FunctionResponsibility
                        ? 'is-invalid'
                        : !model.FunctionResponsibility
                        ? null
                        : 'is-valid'
                    "
                    v-model="model.FunctionResponsibility"
                  />
                </div>
                <div
                  class="invalid-feedback d-block"
                  v-if="errors.FunctionResponsibility"
                >
                  {{ errors.FunctionResponsibility[0] }}
                </div>
              </div>
            </div>
          </div>
          <div class="row">
            <div class="col-lg-12">
              <div class="form-group">
                <label class="d-block">Reason of Resignation</label>
                <div>
                  <textarea
                    type="grid"
                    aria-invalid="true"
                    class="form-control"
                    :disabled="model.PositionLevel == 'Fresh Grad/Less than 1 Year Experience'"

                    :class="
                      errors.ResignationReason
                        ? 'is-invalid'
                        : !model.ResignationReason
                        ? null
                        : 'is-valid'
                    "
                    v-model="model.ResignationReason"
                  />
                </div>
                <div
                  class="invalid-feedback d-block"
                  v-if="errors.ResignationReason"
                >
                  {{ errors.ResignationReason[0] }}
                </div>
              </div>
            </div>
          </div> -->
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
  props: ["refresh"],
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
    model: {
      "name": "",
      "npm" :"",
      "password": "",
      "periode": "",
      "Kelas" : "",
      "l_pict" : "",
      "email" : "",
      "status" : "active",
      "role" : "",
      "major" : "",
    },
    errors: {},
    mandatoryFields: [
      {
        key: "name",
        msg: "name"
      },
      {
        key: "periode",
        msg: "periode"
      },
      {
        key: "kelas",
        msg: "kelas"
      },
      {
        key: "email",
        msg: "email"
      },
      {
        key: "role",
        msg: "role"
      },
      {
        key: "major",
        msg: "major"
      },
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
        console.log(data)
        data.npm = `${data.periode.split("-")[0]}${data.periode.split("-")[1]}${Math.random() * 100}`
        data.periode = moment(data.periode).format("YYYY-MM-DD")
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


        this.$store.dispatch("siswa/postSiswa", data)
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
  },
};
</script>
