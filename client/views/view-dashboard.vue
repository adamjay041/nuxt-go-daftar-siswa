<template>
<div>


  <div class="sticky">
    <Header />
  </div>
  <div>

    <div class="card mt-4">
      <div class="card-body">
          <h3 class="profile-title">Daftar Mahasiswa</h3>
        <div class="row p-5">
          <div class="col-lg-12 text-right">
            <b-button
              variant="primary"
              class="px-5"
              pill
              @click="$bvModal.show('mahasiswa-add')"
              >Daftarkan Mahasiswa</b-button
            >
          </div>
          <div class="col-lg-12 mt-3">
            <table class="c-table">
              <thead>
                <tr>
                  <th></th>
                  <th>NPM</th>
                  <th>Name</th>
                  <th>Major Study</th>
                  <th>Period Information</th>
                  <th>Class</th>
                  <th>Status</th>
                  <div v-if="getRole == 'admin'">
                    <th></th>
                    <th></th>
                  </div>
                </tr>
              </thead>
              <tbody>
                <tr v-if="getlist.length == 0 && !isLoading">
                  <td colspan="10">
                    <img :src="require('@/assets/alert.png')" class="centers" />
                    <h5 class="text-center">
                      Silahkan input data terlebih dahulu
                    </h5>
                  </td>
                </tr>
                <tr v-for="(item, i) in getlist" :key="i">
                  <td>
                    <div class="img-profile-bar" style="margin-right:10px">
                      <img crossorigin="anonymous" width="50px" height="50px" :src="item.l_pict" alt="">
                    </div></td>
                  <td>{{ moreDetails(item.npm) }}</td>
                  <td>
                    {{
                      moreDetails(
                        item.name ? item.name : "-"
                      )
                    }}
                  </td>
                  <td>{{ moreDetails(item.major ? item.major : "-") }}</td>
                  <td>
                    {{
                      moreDetails(
                        item.generation ? item.generation : "-"
                      )
                    }}
                  </td>
                  <td>
                    {{
                      moreDetails(
                        item.kelas ? item.kelas : "-"
                      )
                    }}
                  </td>
                  <td>
                    <div  class="status" >
                      <span>
                          {{
                            moreDetails(
                              item.status ? item.status : "-"
                            )
                          }}
                      </span>
                    </div>
                  </td>
                  <td v-if="getRole == 'admin'">
                    <b-button
                      variant="primary"
                      class="px-2"
                      pill
                      @click="editItem(item)"
                      ><feather-icon
                        name="edit"
                        :width="20"
                        weight="3"
                        style="color: #fffff"
                    /></b-button>
                  </td>
                  <td v-if="getRole == 'admin'">
                    <b-button
                      variant="primary"
                      class="px-2"
                      pill
                      @click="removeItem(item)"
                      ><feather-icon
                        name="trash-2"
                        :width="20"
                        weight="3"
                        style="color: #fffff"
                    /></b-button>
                  </td>
                </tr>
              </tbody>
              <tr v-if="isLoading">
                <td colspan="10">
                  <loading-2 class="m-10 p-10" />
                </td>
              </tr>
            </table>
          </div>
        </div>
      </div>
    </div>
    <siswa-add :refresh="refreshList" />
    <siswa-edit :datas="selectedItem" :refresh="refreshList" />
  </div>
  <!-- <Card-info :data="list[0]" /> -->
</div>
</template>

<script>


  // <b-card style="margin:20px">
  //     <b-tabs content-class="mt-3" pills card>
  //       <b-tab title="List Daftar Siswa" active>
  //         <div class="center">
  //           <div v-for="(i,j) in list" :key="i.name">


  //               <Card-info v-if="j < perPage * currentPage && j >= perPage * (currentPage-1)" :data="i">

  //               </Card-info >
  //           </div>
  //         </div>
  //         <b-pagination
  //           v-model="currentPage"
  //           :per-page="perPage"
  //           :total-rows="list.length"
  //           aria-controls="my-table"
  //           align="right"
  //           style="padding:10px"
  //         ></b-pagination>
  //       </b-tab>
  //       <b-tab title="Daftar Siswa"><p>I'm the second tab</p></b-tab>
  //       <b-tab title="Edit Siswa"><p>I'm the tab with the very, very long title</p></b-tab>
  //     </b-tabs>
  //   <!-- </div> -->

  // </b-card>

import SiswaAdd from "../pages/siswa.add.x.vue"
import SiswaEdit from "../pages/siswa.edit.x.vue"

export default {
  name: 'PageSiswa',
  components : {SiswaAdd,SiswaEdit},
  mounted: {

  },
  data() {
    return {
      isLoading: true,
      dataList: [
        {
          npm: "202143501946",
          name: "adam jay pramusti",
          major: "Teknik Informatika",
          period: "2021/2022 Ganjil",
          generasi: "2021",
          kelas: "S3N",
          status :"Aktif"
        },
        {
          npm: "202143501946",
          name: "adam jay",
          major: "Teknik Informatika",
          period: "2021/2022 Ganjil",
          generasi: "2021",
          kelas: "S3N",
          status :"Aktif"
        },
        {
          npm: "202143501946",
          name: "adam jay",
          major: "Teknik Informatika",
          period: "2021/2022 Ganjil",
          generasi: "2021",
          kelas: "S3N",
          status :"Aktif"
        },
        {
          npm: "202143501946",
          name: "adam jay",
          major: "Teknik Informatika",
          period: "2021/2022 Ganjil",
          generasi: "2021",
          kelas: "S3N",
          status :"Aktif"
        },
                {
          npm: "202143501946",
          name: "adam jay pramusti",
          major: "Teknik Informatika",
          period: "2021/2022 Ganjil",
          generasi: "2021",
          kelas: "S3N",
          status :"Aktif"
        },
        {
          npm: "202143501946",
          name: "adam jay",
          major: "Teknik Informatika",
          period: "2021/2022 Ganjil",
          generasi: "2021",
          kelas: "S3N",
          status :"Aktif"
        },
        {
          npm: "202143501946",
          name: "adam jay",
          major: "Teknik Informatika",
          period: "2021/2022 Ganjil",
          generasi: "2021",
          kelas: "S3N",
          status :"Aktif"
        },
        {
          npm: "202143501946",
          name: "adam jay",
          major: "Teknik Informatika",
          period: "2021/2022 Ganjil",
          generasi: "2021",
          kelas: "S3N",
          status :"Aktif"
        }
      ],
      currentPage: 1,
      perPage: 6,
      selectedItem : null
    }

  },
  computed: {
      getlist() {
        return this.$store.state.siswa.listSiswa
      },
      getRole() {
        return this.$store.state.siswa.role
      }
  },

  methods: {
      moreDetails: function (item) {
      if (item) {
        if (item !== null) {
          if (item.length > 20) {
            return item.substring(0, 30) + "...";
          }
          return item;
        }
      }
    },
    getData() {
      this.$store.dispatch("siswa/getListSiswa")
    },
  removeItem(val) {
    console.log(val)
    this.$store.dispatch('siswa/deleteSiswa',val.id)
    this.getData()
  },
  editItem(item) {
    this.selectedItem = item;
    this.$bvModal.show("mahasiswa-edit");
  },

  refreshList: function () {},
  },
  mounted(){
    console.log('this mounted')
    this.getData()
  }
}
</script>


<style lang="scss">
  #__nuxt {
    height: 100%;
  }

  #__layout {
    height: 100%;
  }

  // .containers {
  //   height: 100%;
  //   background-repeat: no-repeat;
  //   background-size: cover;
  //   background-position: center;
  // }


  .centers {
    // overflow: auto;
    overflow: auto;
    // top: 0;
    // height: inherit;
    display: flex;
    align-content: center;
    align-items: center;
    justify-content: center;
    flex-wrap: wrap;
  }
  .sticky {
    overflow:scroll;
    background-color: #333;
    position:sticky;
    // top: 0;
    width: 100%;
  }
</style>
