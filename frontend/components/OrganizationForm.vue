<template>
  <div onselectstart="return false;">
    <v-card class="mx-auto my-4 px-8 py-4" style="max-width: 700px;">
      <div style="position: relative">
        <v-card-actions class="py-6">
          <v-btn icon @click="$router.push(to)">
            <v-icon>mdi-arrow-left</v-icon>
          </v-btn>
          <v-spacer></v-spacer>
        </v-card-actions>
      </div>
      <v-form ref="form" v-model="formReady" class="pa-4 pt-6">
        <v-text-field
          v-model="name"
          :rules="[rules.required, rules.length(20)]"
          filled
          color="#295caa"
          label="プロジェクト名"
          type="text"
        ></v-text-field>
        <v-text-field
          v-model="description"
          filled
          color="#295caa"
          label="プロジェクトの概要"
          type="text"
        ></v-text-field>
        <v-text-field
          v-model="founded"
          filled
          color="#295caa"
          label="設立日"
          type="text"
        ></v-text-field>
        <v-text-field
          v-model="number"
          filled
          color="#295caa"
          label="電話番号"
          type="text"
        ></v-text-field>
        <cropper v-model="image" ratio="16:9" :width="450" :pixel="900" :currentImage="$config.mediaURL + '/organizations/' + value.image"></cropper>
        <div class="px-4 py-2 red--text accent-3 text-center" style="height: 80px">{{ this.error }}</div>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn text @click="$router.push(to)">
            <slot name="back"></slot>
          </v-btn>
          <v-btn :disabled="!formReady || loading" :loading="loading" class="white--text" color="#295caa" depressed @click="onSubmit">
            <slot name="submit"></slot>
          </v-btn>
        </v-card-actions>
      </v-form>
    </v-card>
  </div>
</template>

<script lang="ts">

import Vue from 'vue'
import { mapGetters } from 'vuex'
import { isEmptyObj, resizeFile } from '~/modules/utils'
import * as lib from '~/modules/store'
interface ProjectAuthority extends lib.ProjectAuthority {
  disabled: boolean
  project: lib.Project
}
export default Vue.extend({
  props: {
    value: {
      type: Object
    },
    loading: {
      type: Boolean
    }
  },
  data:() => ({
    isAuthorized: false,
    formReady: false,
    rules: {
      length: (len: number) => (v: string) => (v || '').length <= len || `最大20文字までです`,
      required: (v: string) => !!v || '必ずご記入ください',
      requiredSelect: (v: lib.User[]) => v.length != 0 || '1名は選択してください',
    },
    error: '',
  }),
  computed: {
    ...mapGetters([
      'project',
      'projectAuthority',
      'user',
    ]),
    isEmptyObj,
    name: {
      get() {
        return this.value.name;
      },
      set(name) {
        this.updateValue({name});
      }
    },
    description: {
      get() {
        return this.value.description;
      },
      set(description) {
        this.updateValue({description});
      }
    },
    founded: {
      get() {
        return this.value.founded;
      },
      set(founded) {
        this.updateValue({founded});
      }
    },
    number: {
      get() {
        return this.value.number;
      },
      set(number) {
        this.updateValue({number});
      }
    },
    image: {
      get() {
        return this.value.image_data;
      },
      set(image_data) {
        this.updateValue({image_data});
      }
    },
    imageSrc() {
      return this.value?.image_data || this.$config.mediaURL + '/projects/' + this.value.image;
    },
    projectUserItems() {
      return this.project.authority_users.map((user: ProjectAuthority) => {
        if (user.user_id == this.projectAuthority.user_id) {
          user.disabled = true;
        }
        return user;
      })
    },
    to() {
      return 'id' in this.$route.params ? {name: 'project-id', params: {id: this.$route.params.id}, query: {tab: 'project'}} : {name: 'index'};
    }
  },
  methods: {
    updateValue(value: {}) {
      const newProject = {...this.value, ...value};
      this.$emit('input', newProject);
    },
    async onChangeFile(e: File) {
      if(!e) return this.value.image_data = '';
      const image_data = await resizeFile(e);
      const newProject = {...this.value.project, ...{image_data}};
      const newValue = {...this.value, ...{project: newProject}};
      this.$emit('input', newValue);
    },
    onSubmit() {
      this.$emit('submit');
    },
  }
})
</script>

<style lang="scss">
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