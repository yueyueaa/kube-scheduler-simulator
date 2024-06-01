<template>
</template>

<script lang="ts">
import { V1PersistentVolume } from "@kubernetes/client-node";
import { computed, inject, defineComponent } from "@nuxtjs/composition-api";
import {} from "../../lib/util";
import PersistentVolumeStoreKey from "../../StoreKey/PVStoreKey";
export default defineComponent({
  setup() {
    const store = inject(PersistentVolumeStoreKey);
    if (!store) {
      throw new Error(`${PersistentVolumeStoreKey.description} is not provided`);
    }

    const onClick = (pv: V1PersistentVolume) => {
      store.select(pv, false);
    };

    const pvs = computed(() => store.pvs);
    return {
      pvs,
      onClick,
    };
  },
});
</script>
