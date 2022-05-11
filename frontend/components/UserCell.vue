<template>
  <div class="cell text-left" :style="styleValue">
    <v-btn class="mx-1" icon>
      <v-avatar size="25px">
        <v-img alt="Avatar" style="object-fit: cover;" :src="$config.mediaURL + '/users/' + user.image">
          <template v-slot:placeholder>
            <v-row class="fill-height ma-0" align="center" justify="center">
              <v-progress-circular indeterminate color="grey lighten-5" />
            </v-row>
          </template>
        </v-img>
      </v-avatar>
    </v-btn>
    {{ user.name }}
  </div>
</template>

<script>
import Vue from 'vue'
export default Vue.extend({
  props: ['styleValue', 'user']
})
</script>