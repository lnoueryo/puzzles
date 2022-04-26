<template>
  <div>
    <v-card class="my-4 px-8 py-4">
      <v-card-actions>
        <v-btn text :to="{name: 'project-id-task', params: {id: $route.params.id}}">
          <slot name="back"></slot>
        </v-btn>
        <v-spacer></v-spacer>
      </v-card-actions>
      <v-form ref="form" v-model="formReady">
        <v-card-text>
          <v-row class="mb-6" no-gutters>
            <v-col sm="12" md="12">
              <v-text-field
                ref="title"
                v-model="title"
                :rules="[rules.required]"
                label="課題のタイトル"
                prepend-icon="mdi-format-list-bulleted-square"
                required
              ></v-text-field>
            </v-col>
          </v-row>
          <markdown @currentText="detail = $event" :text="detail"></markdown>
          <v-row>
            <v-col sm="5" md="6">
              <v-select
                ref="assignee"
                v-model="assignee"
                :items="projectAuthority.project.authority_users"
                item-text="user.name"
                item-value="user_id"
                label="担当者"
                prepend-icon="mdi-account-tie"
              ></v-select>
            </v-col>
            <v-col sm="5" md="6">
              <v-select
                ref="status"
                v-model="status"
                :items="statuses"
                item-text="name"
                item-value="id"
                label="現在の状態"
                prepend-icon="mdi-playlist-check"
              ></v-select>
            </v-col>
          </v-row>
          <v-row>
            <v-col sm="6" md="6">
            <v-select
                ref="type"
                v-model="type"
                :items="types"
                label="タスクの種類"
                item-text="name"
                item-value="id"
                prepend-icon="mdi-priority-high"
              ></v-select>
            </v-col>
            <v-col sm="6" md="6">
            <v-select
                ref="priority"
                v-model="priority"
                :items="priorities"
                label="優先順位"
                item-text="name"
                item-value="id"
                prepend-icon="mdi-priority-high"
              ></v-select>
            </v-col>
          </v-row>
          <v-row>
            <v-col sm="6" md="6">
            <v-menu
              v-model="deadlineDate"
              :close-on-content-click="false"
              :nudge-right="40"
              transition="scale-transition"
              offset-y
              min-width="auto"
            >
              <template v-slot:activator="{ on, attrs }">
                <v-text-field
                  ref="deadline"
                  v-model="deadline"
                  label="タスクの期日"
                  prepend-icon="mdi-calendar"
                  :rules="[v => !!v || '必ず選択してください']"
                  readonly
                  v-bind="attrs"
                  v-on="on"
                ></v-text-field>
              </template>
              <v-date-picker v-model="deadline" :min="min" @input="deadlineDate = false"></v-date-picker>
            </v-menu>
            </v-col>
            <v-col sm="6" md="6">
            <v-text-field
              ref="estimated_time"
              v-model="estimated_time"
              :rules="[isNumber(estimated_time)]"
              label="推定時間"
              required
              prepend-icon="mdi-calendar-clock"
            ></v-text-field>
            </v-col>
          </v-row>
          <v-row align="center">
            <v-col sm="6" md="6">
              <v-select
                id="field"
                ref="field"
                v-model="field"
                :items="projectAuthority.project.fields"
                item-text="name"
                item-value="id"
                label="分野"
                prepend-icon="mdi-shape"
                :disabled="isEmptyArr(projectAuthority.project.fields)"
              >
              </v-select>
            </v-col>
            <v-col sm="6" md="6">
              <v-select
                ref="milestone"
                v-model="milestone"
                :items="projectAuthority.project.milestones"
                item-text="name"
                item-value="id"
                label="マイルストーン"
                prepend-icon="mdi-flag-triangle"
                :disabled="isEmptyArr(projectAuthority.project.fields)"
              ></v-select>
            </v-col>
          </v-row>
          <small v-if="showConfigLink">分野、マイルストーンは<nuxt-link :to="{name: 'project-id-edit', params: {id: $route.params.id}}">設定画面より</nuxt-link>作成頂けます</small>
        </v-card-text>
        <!-- <v-divider class="mt-12"></v-divider> -->
        <v-card-actions>
          <v-btn text :to="{name: 'project-id-task', params: {id: $route.params.id}}">
            <slot name="back"></slot>
          </v-btn>
          <v-spacer></v-spacer>
          <v-btn color="primary" :loading="loading" text @click="onSubmit" :disabled="!formReady || loading">
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
import { isEmptyArr, isNumber, changeToISOFormat } from '~/modules/utils'
import * as lib from '~/modules/store'
declare module 'vue/types/vue' {
  interface Vue {
    updateValue: (v: {}) => void;
  }
}
export default Vue.extend({
  props: {
    value: {
      type: Object
    },
    loading: {
      type: Boolean,
    }
  },
  data: () => ({
    formReady: false,
    deadlineDate: false,
    rules: {
      length: (len: number) => (v: string) => (v || '').length <= len || `最大20文字までです`,
      required: (v: string) => !!v || '必ずご記入ください',
      requiredSelect: (v: lib.User[]) => v.length != 0 || '1名は選択してください',
    },
  }),
  computed: {
    ...mapGetters('task', [
      'statuses',
      'types',
      'priorities',
    ]),
    ...mapGetters([
      'projectAuthority',
    ]),
    isEmptyArr,
    isNumber,
    changeToISOFormat,
    title: {
      get() {
        return this.value.title;
      },
      set(title) {
        this.updateValue({title});
      }
    },
    detail: {
      get() {
        return this.value.detail;
      },
      set(detail) {
        this.updateValue({detail});
      }
    },
    assignee: {
      get() {
        return this.value.assignee_id;
      },
      set(assignee_id) {
        this.updateValue({assignee_id});
      }
    },
    status: {
      get() {
        return this.value.status_id;
      },
      set(status_id) {
        this.updateValue({status_id});
      }
    },
    type: {
      get() {
        return this.value.type_id;
      },
      set(type_id) {
        this.updateValue({type_id});
      }
    },
    priority: {
      get() {
        return this.value.priority_id;
      },
      set(priority_id) {
        this.updateValue({priority_id});
      }
    },
    field: {
      get() {
        return this.value.field_id;
      },
      set(field_id) {
        this.updateValue({field_id});
      }
    },
    milestone: {
      get() {
        return this.value.milestone_id;
      },
      set(milestone_id) {
        this.updateValue({milestone_id});
      }
    },
    deadline: {
      get() {
        return this.value.deadline;
      },
      set(deadline) {
        this.updateValue({deadline});
      }
    },
    estimated_time: {
      get() {
        return this.value.estimated_time;
      },
      set(estimated_time) {
        this.updateValue({estimated_time});
      }
    },
    min() {
      return this.changeToISOFormat('')
    },
    showConfigLink() {
      if(this.projectAuthority.authority != '管理者') {
        return;
      }
      const fields = this.projectAuthority.project.fields;
      const milestones = this.projectAuthority.project.milestones;
      return this.isEmptyArr(fields) || this.isEmptyArr(milestones)
    },
  },
  methods: {
    updateValue(value: {}) {
      const newValue = {...this.value, ...value};
      this.$emit('input', newValue);
    },
    onSubmit() {
      this.$emit('submit');
    }
  }
})
</script>

<style lang="scss" scoped>
.v-application--is-ltr .v-text-field .v-label {
    color: cadetblue;
}
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