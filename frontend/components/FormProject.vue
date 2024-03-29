<template>
  <div>
    <v-card class="mx-auto my-4 px-8 py-4" max-width="700px">
      <div class="relative">
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
          v-model="field.name"
          v-for="(field, i) in fields"
          :key="'field_' + i"
          :rules="[rules.length(20)]"
          filled
          color="#295caa"
          :label="'分野 ' + (i + 1)"
          type="text"
          append-icon="mdi-close"
          @click:append="onDeleteField(i)"
        ></v-text-field>
        <v-btn class="mb-8" block @click="fields.push({name: ''})" :disabled="!fields[fields.length - 1].name">追加</v-btn>
        <v-text-field
          v-model="milestone.name"
          v-for="(milestone, i) in milestones"
          :key="'milestone_' + i"
          :rules="[rules.length(20)]"
          filled
          color="#295caa"
          :label="'マイルストーン ' + (i + 1)"
          type="text"
          append-icon="mdi-close"
          @click:append="onDeleteMilestone(i)"
        >
        </v-text-field>
        <v-btn class="mb-8" block @click="milestones.push({name: ''})" :disabled="!milestones[milestones.length - 1].name">追加</v-btn>
        <v-text-field
          v-model="version.name"
          v-for="(version, i) in versions"
          :key="'version_' + i"
          :rules="[rules.length(20)]"
          filled
          color="#295caa"
          :label="'バージョン ' + (i + 1)"
          type="text"
          append-icon="mdi-close"
          @click:append="onDeleteVersion(i)"
        >
        </v-text-field>
        <v-btn class="mb-8" block @click="versions.push({name: ''})" :disabled="!versions[versions.length - 1].name">追加</v-btn>
        <!-- <v-select
          v-model="authorityUsers"
          :items="projectUserItems"
          label="管理者"
          item-text="user.name"
          item-value="user.id"
          :rules="[rules.requiredSelect]"
          item-disabled="disabled"
          multiple
          v-if="!isEmptyObj(project)"
        >
        </v-select> -->
        <v-cropper v-model="image" :pixel="900" ratio="16:9" :width="450" :currentImage="currentImage"></v-cropper>
        <div class="px-4 py-2 red--text accent-3 text-center error-height">{{ this.error }}</div>
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
import * as model from '~/modules/model'
interface ProjectAuthority extends model.ProjectAuthority {
  disabled: boolean
  project: model.Project
}
export default Vue.extend({
  props: {
    value: Object,
    loading: Boolean
  },
  data:() => ({
    isAuthorized: false,
    formReady: false,
    rules: {
      length: (len: number) => (v: string) => (v || '').length <= len || `最大20文字までです`,
      required: (v: string) => !!v || '必ずご記入ください',
      requiredSelect: (v: model.User[]) => v.length != 0 || '1名は選択してください',
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
    fields: {
      get(): model.Field[] {
        return this.value.fields;
      },
      set(fields) {
        this.updateValue({fields});
      }
    },
    milestones: {
      get(): model.Milestone[] {
        return this.value.milestones;
      },
      set(milestones) {
        this.updateValue({milestones});
      }
    },
    versions: {
      get(): model.Version[] {
        return this.value.versions;
      },
      set(versions) {
        this.updateValue({versions});
      }
    },
    // authorityUsers: {
    //   get(): model.ProjectAuthority[] {
    //     return this.value.authority_users.filter((authority_user: model.ProjectAuthority) => authority_user.auth_id == 1);
    //   },
    //   set(_authority_users: number[]) {
    //     const original_authority_users = JSON.parse(JSON.stringify(this.project.authority_users))
    //     const authority_users = original_authority_users.map((authority_user: model.ProjectAuthority) => {
    //       if (_authority_users.includes(authority_user.user.id)) {
    //         authority_user.auth_id = 1;
    //       } else {
    //         authority_user.auth_id = 2;
    //       }
    //       return authority_user;
    //     })

    //     this.updateValue({authority_users});
    //   }
    // },
    image: {
      get() {
        return this.value.image_data;
      },
      set(image_data) {
        this.updateValue({image_data});
      }
    },
    currentImage() {
      return this.value.image_data || this.$config.mediaURL + '/projects/' + this.value.image
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
    onSubmit() {
      this.$emit('submit');
    },
    onDeleteField(index: number) {
      if((this.fields as model.Field[]).length == 1) return;
      (this.fields  as model.Field[]).splice(index, 1)
    },
    onDeleteMilestone(index: number) {
      if((this.milestones as model.Milestone[]).length == 1) return;
      (this.milestones as model.Milestone[]).splice(index, 1)
    },
    onDeleteVersion(index: number) {
      if((this.versions as model.Milestone[]).length == 1) return;
      (this.versions as model.Milestone[]).splice(index, 1)
    },
  }
})
</script>

<style lang="scss" scoped>
.custom-loader {
  animation: loader 1s infinite;
  display: flex;
}
.error-height {
  height: 80px;
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