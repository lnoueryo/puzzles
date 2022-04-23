<template>
  <div>
    <v-card class="mx-auto my-4 px-8 py-4" style="max-width: 700px;">
      <v-card-actions>
        <v-btn text :to="to">
          <slot name="back"></slot>
        </v-btn>
        <v-spacer></v-spacer>
      </v-card-actions>
      <v-form ref="form" v-model="formReady" class="pa-4 pt-6">
        <v-text-field
          v-model="name"
          :rules="[rules.required, rules.length(20)]"
          filled
          color="amber darken-3"
          label="プロジェクト名"
          type="text"
        ></v-text-field>
        <v-text-field
          v-model="description"
          filled
          color="amber darken-3"
          label="プロジェクトの概要"
          type="text"
        ></v-text-field>
        <v-text-field
          v-model="field.name"
          v-for="(field, i) in fields"
          :key="'field_' + i"
          :rules="[rules.length(20)]"
          filled
          color="amber darken-3"
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
          color="amber darken-3"
          :label="'マイルストーン ' + (i + 1)"
          type="text"
          append-icon="mdi-close"
          @click:append="onDeleteMilestone(i)"
        >
        </v-text-field>
        <v-btn class="mb-8" block @click="milestones.push({name: ''})" :disabled="!milestones[milestones.length - 1].name">追加</v-btn>
        <v-select
          v-model="authorityUsers"
          :items="projectUserItems"
          label="管理者"
          item-text="user.name"
          item-value="user.id"
          :rules="[rules.requiredSelect]"
          item-disabled="disabled"
          multiple
          v-if="!isEmptyObj(projectAuthority)"
        >
        </v-select>
        <cropper v-model="image" :width="450" :currentImage="$config.mediaURL + '/projects/' + value.project.image"></cropper>
        <div class="px-4 py-2 red--text accent-3 text-center" style="height: 80px">{{ this.error }}</div>
        <v-card-actions>
          <v-btn text :to="to">
            <slot name="back"></slot>
          </v-btn>
          <v-spacer></v-spacer>
          <v-btn :disabled="!formReady || loading" :loading="loading" class="white--text" color="amber darken-3" depressed @click="onSubmit">
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
      'projectAuthority',
      'user',
    ]),
    isEmptyObj,
    name: {
      get() {
        return this.value.project.name;
      },
      set(name) {
        this.updateValue({name});
      }
    },
    description: {
      get() {
        return this.value.project.description;
      },
      set(description) {
        this.updateValue({description});
      }
    },
    fields: {
      get(): lib.Field[] {
        return this.value.project.fields;
      },
      set(fields) {
        this.updateValue({fields});
      }
    },
    milestones: {
      get(): lib.Milestone[] {
        return this.value.project.milestones;
      },
      set(milestones) {
        this.updateValue({milestones});
      }
    },
    authorityUsers: {
      get(): lib.ProjectAuthority[] {
        return this.value.project.authority_users.filter((authority_user: lib.ProjectAuthority) => authority_user.auth_id == 1);
      },
      set(_authority_users: number[]) {
        const original_authority_users = JSON.parse(JSON.stringify(this.projectAuthority.project.authority_users))
        const authority_users = original_authority_users.map((authority_user: lib.ProjectAuthority) => {
          if (_authority_users.includes(authority_user.user.id)) {
            authority_user.auth_id = 1;
          } else {
            authority_user.auth_id = 2;
          }
          return authority_user;
        })

        this.updateValue({authority_users});
      }
    },
    image: {
      get() {
        return this.value.project.image_data;
      },
      set(image_data) {
        this.updateValue({image_data});
      }
    },
    imageSrc() {
      return this.value.project?.image_data || this.$config.mediaURL + '/projects/' + this.value.project.image;
    },
    projectUserItems() {
      return this.projectAuthority.project.authority_users.map((user: ProjectAuthority) => {
        if (user.user_id == this.projectAuthority.user_id) {
          user.disabled = true;
        }
        return user;
      })
    },
    to() {
      return 'id' in this.$route.params ? {name: 'project-id-task', params: {id: this.$route.params.id}} : {name: 'project'};
    }
  },
  methods: {
    updateValue(value: {}) {
      const newProject = {...this.value.project, ...value};
      this.$emit('input', {...this.value, ...{project: newProject}});
    },
    async onChangeFile(e: File) {
      if(!e) return this.value.project.image_data = '';
      const image_data = await resizeFile(e);
      const newProject = {...this.value.project, ...{image_data}};
      const newValue = {...this.value, ...{project: newProject}};
      this.$emit('input', newValue);
    },
    onSubmit() {
      this.$emit('submit');
    },
    onDeleteField(index: number) {
      if((this.fields as lib.Field[]).length == 1) return;
      (this.fields  as lib.Field[]).splice(index, 1)
    },
    onDeleteMilestone(index: number) {
      if((this.milestones as lib.Milestone[]).length == 1) return;
      (this.milestones as lib.Milestone[]).splice(index, 1)
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