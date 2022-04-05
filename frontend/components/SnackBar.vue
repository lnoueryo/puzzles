<template>
  <div class="text-center">
    <v-snackbar v-model="snackbar" :timeout="timeout">
      {{ snackbarText }}
      <template v-slot:action="{ attrs }">
        <v-btn color="blue" text v-bind="attrs" @click="$store.commit('snackbar', false)">
          Close
        </v-btn>
      </template>
    </v-snackbar>
  </div>
</template>

<script>
export default {
  data: () => ({
    timeout: 2000
  }),
  computed: {
    snackbar: {
      get() {
        return this.$store.getters['snackbar'];
      },
      set(v) {
        this.$store.commit('snackbar', v);
        if(!v) return this.$store.commit('snackbarText', '');
      }
    },
    snackbarText() {
      return this.$store.getters['snackbarText'];
    }
  }
}
</script>