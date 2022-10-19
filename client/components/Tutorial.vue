<!-- Please remove this file from your project -->
<template>
<div>
  <Header />
  <b-card class="3px 2py" style="margin:30px;">
    <div style="display:flex;justify-content:space-between;padding-bottom:20px;">
      <!-- <b-col md="4">md="4"</b-col> -->
      <!-- <b-col md="4">md="4" offset-md="4"</b-col> -->
      <h3>Daftar Mahasiswa</h3>
      <b-button variant="success" size="md" v-b-modal.modal-prevent-closing @click="addSiswa('Tambah Daftar Mahasiswa')">Tambah Daftar Mahasiswa </b-button>
    </div>
    <!-- <div align="right" class="justify">

    </div> -->
    <div class="overflow-auto">
      <b-table  :items="getlist" :fields="fields"  outlined bordered responsive class="md" :per-page="perPage" :current-page="currentPage" show-empty>
          <template #cell(name)="row">
           <p> {{ row.value.first }} {{ row.value.last }} </p>
          </template>
          <template #cell(actions)="row">
            <div v-if="row.item.Role == 'student'">
              <b-button size="md"  class="mr-1" variant="primary" v-b-modal.modal-prevent-closing  @click="commitEdit(row, 'Edit Data Mahasiswa' )" >
                Edit  <b-icon icon="pencil" font-scale="1"></b-icon>
              </b-button>
              <b-button size="md"  class="mr-1" variant="danger" @click="deleteData(row)">
                Delete  <b-icon icon="trash" font-scale="1"></b-icon>
              </b-button>
            </div>
          </template>
      </b-table>
      <b-pagination
        v-model="currentPage"
        :total-rows="getlist.length"
        :per-page="perPage"
        aria-controls="my-table"
        align="right"
      ></b-pagination>
      <p class="mt-3" align="right">Current Page: {{ currentPage }}</p>
    </div>
  </b-card>
  <!-- modal update -->

  <modal-form :dataEdit="editData" :id="id"></modal-form>



</div>

</template>

<script>
import modalForm from './modalForm.vue'

  export default {
  components: { modalForm },
    data() {
      return {
        editData: '',
        id: '',
        fields: [
          {key: "Id", label: "Id", class: 'text-center'},
          {key: "First_Name", label: "Fist Nama", class: 'text-center'},
          {key: "Last_Name", label: "Last Nama", class: 'text-center'},
          {key: "Full_Name", label: "Full Nama", class: 'text-center'},
          {key: "Class", label: "Kelas", class: 'text-center'},
          {key: "NPM", label: "NPM", class: 'text-center'},
          {key: "Email", label: "Email", class: 'text-center'},
          {key: "actions", label: "Actions", class: 'text-center'
          },



          // { key: 'name', label: 'Person full name', class: 'text-center' },
          // { key: 'age', label: 'Person age',  class: 'text-center' },
          // {
          //   key: 'isActive',
          //   label: 'Is Active',

          //   class: 'text-center',

          //   sortByFormatted: true,
          //   filterByFormatted: true
          // },
          // { key: 'actions', label: 'Actions',class: 'text-center' }
        ],
        currentPage: 1,
        perPage: 8,

      }
    },
    methods: {
      commitEdit(val, titel) {
        console.log(val.item['Id'])
        // val.item['NPM'] = parseInt(val.item['NPM'])
        this.editData = val.item
        this.id = val.item['Id']
        this.$store.commit("siswa/setFormSubmitDefault")
        this.$store.commit("siswa/setTitleModal", titel)
        this.$store.commit("siswa/setFormSubmit", val.item)
      },
      addSiswa(val) {
        this.$store.commit("siswa/setFormSubmitDefault")
        this.$store.commit("siswa/setTitleModal", val)
      },
      deleteData(val) {
        this.$store.dispatch('siswa/deleteSiswa',val.item['Id'])
        this.getData()
      },
      getData() {
        this.$store.dispatch("siswa/getListSiswa")
      }
    },
    computed: {
      getlist() {
        return this.$store.state.siswa.listSiswa
      }
    },

  }
</script>
