<template>
  <div>
    <input type="file" multiple @input="upload">
    <v-btn @click="send">upload</v-btn>
    <v-btn @click="download">download</v-btn>
      <v-select
        v-model="downloadFiles"
        :items="tables"
        label="Table"
        multiple
      ></v-select>
    <v-card v-for="(table, i) in errors" :key="i">
      <v-card-text>
        <div v-for="(error, j) in table" :key="j">
          {{ error }}
        </div>
      </v-card-text>
    </v-card>
    <div class="text-center" v-if="loading" style="position: fixed;top: 0;bottom: 0;right: 0;left: 0;margin: auto;z-index: 20">
      <v-progress-circular
        indeterminate
        color="primary"
         style="position: fixed;top: 0;bottom: 0;right: 0;left: 0;margin: auto;z-index: 20"
      ></v-progress-circular>
    </div>
  </div>
</template>

<script>
export default {
  data: () => ({
    downloadFiles: [],
    uploadFiles: '',
    errors: [],
    loading: false
  }),
  computed: {
    tables() {
      return [
        'activities',
        'activity_contents',
        'authorities',
        'comments',
        'fields',
        'milestones',
        'versions',
        'organization_authorities',
        'organizations',
        'priorities',
        'project_authorities',
        'projects',
        'statuses',
        'tasks',
        'types',
        'users',
      ]
    }
  },
  methods: {
    async download() {
      this.loading = true;
      let response
      try {
        response = await this.$axios.post('/api/data/download', {request: this.downloadFiles},{
          responseType: "blob",
          headers: {
            'content-type': 'multipart/form-data',
          }
        });
        const url = window.URL.createObjectURL(new Blob([response.data],{type:'application/zip'}))
        // HTML要素のaタグを生成
        const link = document.createElement('a')
        link.href = url
        // aタグのdownload属性を設定
        link.setAttribute('download', `a.zip`)
        // 生成したaタグを設置し、クリックさせる
        document.body.appendChild(link)
        link.click()
        // URLを削除
        window.URL.revokeObjectURL(url)
      } catch (error) {
        response = error.response
      }
      console.log(response)
      this.loading = false;
    },
    upload(e) {
      this.uploadFiles = e.target.files;
    },
    async send() {
      this.loading = true;
      const formData = new FormData()
      for (let i = 0; i < this.uploadFiles.length; i++) {
        formData.append("files", this.uploadFiles[i]);
      }
      try {
        const response = await this.$axios.post('/api/data/upload', formData, {headers: {'Content-Type': 'multipart/form-data'}})
        console.log(response)
        this.errors = response.data;
      } catch (error) {
        console.log(error.response)
      }
      this.loading = false;
    }
  }
}
</script>