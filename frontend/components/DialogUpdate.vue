<template>
  <v-row justify="center">
    <v-dialog v-model="value" scrollable max-width="700px">
      <v-card>
        <v-card-title><slot></slot></v-card-title>
        <v-divider></v-divider>
          <v-card-text class="py-6 text-center card-height">
            <p v-if="changeNum != 0">{{changeNum}}件の変更があります。</p>
            <div class="d-flex w100">
              <div class="w50">
                <p>変更前</p>
              </div>
              <div class="w50">
                <p>変更後</p>
              </div>
            </div>
            <div v-for="(content, i) in editedData" :key="i">
              <div class="d-flex w100" v-if="content['image']">
                <div class="w50" v-if="content.change">
                  <div>{{ content.title }}</div>
                  <div>
                    <v-img class="image-width" :src="$config.mediaURL + content.oldData" :lazy-src="require('~/assets/image/project.png')" v-if="content.oldData">
                      <template v-slot:placeholder>
                        <v-row class="fill-height ma-0" align="center" justify="center">
                          <v-progress-circular indeterminate color="grey lighten-5" />
                        </v-row>
                      </template>
                    </v-img>
                  </div>
                </div>
                <div class="w50" v-if="content.change">
                  <div :class="[{'red--text': content.change}, {'text--primary': !content.change}]">{{ content.title }}</div>
                  <div>
                    <img class="image-width" :src="content.newData" alt="">
                  </div>
                </div>
              </div>
              <div class="d-flex w100" v-else>
                <div class="w50">
                  <div>{{ content.title }}</div>
                  <p>{{ content.oldData }}</p>
                </div>
                <div class="w50">
                  <div :class="[{'red--text': content.change}, {'text--primary': !content.change}]">{{ content.title }}</div>
                  <p>{{ content.newData }}</p>
                </div>
              </div>
            </div>
          </v-card-text>
          <v-divider></v-divider>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="blue darken-1" text @click="$emit('input', false)">
            戻る
          </v-btn>
          <v-btn color="blue darken-1" text @click="onClickSubmit" :loading="loading">
            保存
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script>
export default {
  props: {
    value: Boolean,
    form: Array,
    loading: Boolean
  },
  data: () => ({
    changeNum: 0,
  }),
  computed: {
    editedData() {
      let changeNum = 0;
      const editedData = this.form.map((content) => {
        content.change = this.edited(content.newData, content.oldData);
        if(content.change) changeNum += 1;
        return content;
      });
      this.changeNum = changeNum;
      return editedData;
    }
  },
  methods: {
    edited(newValue, oldValue) {
      return newValue != oldValue;
    },
    onClickSubmit() {
      this.$emit('loading', true);
      if(this.changeNum === 0) return this.$router.back();
      this.$emit('submit')
    }
  }
}
</script>


<style lang="scss" scoped>
.custom-loader {
  animation: loader 1s infinite;
  display: flex;
}
.card-height {
  height: 400px;
}
.image-width {
  max-width: 300px;
}
@-moz-keyframes loader {
  from {
    transform: rotate(0);
  }
  to {
    transform: rotate(360deg);
  }
}
@-webkit-keyframes loader {
  from {
    transform: rotate(0);
  }
  to {
    transform: rotate(360deg);
  }
}
@-o-keyframes loader {
  from {
    transform: rotate(0);
  }
  to {
    transform: rotate(360deg);
  }
}
@keyframes loader {
  from {
    transform: rotate(0);
  }
  to {
    transform: rotate(360deg);
  }
}
</style>