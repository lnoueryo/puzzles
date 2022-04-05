<template>
  <div>
    <input type="file" multiple @input="upload">
    <v-btn @click="send">upload</v-btn>
    <v-btn @click="download">download</v-btn>
  </div>
</template>

<script>
export default {
  data: () => ({
    files: '',
  }),
  methods: {
    async download() {
      const response = await this.$axios.get('/api/data/download', {
      responseType: "blob",
      headers: {
        'content-type': 'multipart/form-data',
      }
    });
      console.log(response)
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
    },
    upload(e) {
    this.files = e.target.files;
    // axios.post('upload_file', file, {
    //     headers: {
    //       'Content-Type': 'multipart/form-data'
    //     }
    // })
    },
    async send() {
      const formData = new FormData()
      for (let i = 0; i < this.files.length; i++) {
        formData.append("files", this.files[i]);
      }
      try {
        const response = await this.$axios.post('/api/data/upload', formData, {headers: {'Content-Type': 'multipart/form-data'}})
        console.log(response)
      } catch (error) {
        console.log(error.response)
      }
    }
  }
}
</script>