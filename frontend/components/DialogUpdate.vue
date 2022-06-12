<template>
  <v-row justify="center">
    <v-dialog v-model="value" scrollable max-width="700px">
      <v-card>
        <v-card-title><slot></slot></v-card-title>
        <v-divider></v-divider>
          <v-card-text class="py-6 text-center" style="height: 400px;">
            <p v-if="changeNum != 0">{{changeNum}}件の変更があります。</p>
            <div class="d-flex" style="width: 100%">
              <div style="width: 50%">
                <p>変更前</p>
              </div>
              <div style="width: 50%">
                <p>変更後</p>
              </div>
            </div>
            <div v-for="(content, i) in editedData" :key="i">
              <div class="d-flex" style="width: 100%" v-if="content['image']">
                <div style="width: 50%" v-if="content.change">
                  <div>{{ content.title }}</div>
                  <div>
                    <v-img style="max-width: 300px" :src="$config.mediaURL + '/projects/' + content.oldData" :lazy-src="require('~/assets/image/project.png')">
                      <template v-slot:placeholder>
                        <v-row class="fill-height ma-0" align="center" justify="center">
                          <v-progress-circular indeterminate color="grey lighten-5" />
                        </v-row>
                      </template>
                    </v-img>
                  </div>
                </div>
                <div style="width: 50%" v-if="content.change">
                  <div :class="[{'red--text': content.change}, {'text--primary': !content.change}]">{{ content.title }}</div>
                  <div>
                    <img style="max-width: 300px" :src="content.newData" alt="">
                  </div>
                </div>
              </div>
              <div class="d-flex" style="width: 100%" v-else>
                <div style="width: 50%">
                  <div>{{ content.title }}</div>
                  <p>{{ content.oldData }}</p>
                </div>
                <div style="width: 50%">
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
    value: {
      type: Boolean
    },
    form: {
      type: Array
    },
    loading: {
      type: Boolean
    }
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