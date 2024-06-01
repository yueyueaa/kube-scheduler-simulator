<template>            
</template>

<script lang="ts">
import { V1PersistentVolumeClaim } from "@kubernetes/client-node";
import { computed, inject, defineComponent } from "@nuxtjs/composition-api";
import {} from "../../lib/util";
import PersistentVolumeClaimStoreKey from "../../StoreKey/PVCStoreKey";
export default defineComponent({
  setup() {
    const store = inject(PersistentVolumeClaimStoreKey);
    if (!store) {
      throw new Error(`${PersistentVolumeClaimStoreKey.description} is not provided`);
    }

    const onClick = (pvc: V1PersistentVolumeClaim) => {
      store.select(pvc, false);
    };

    const pvcs = computed(() => store.pvcs);
    return {
      pvcs,
      onClick,
    };
  },
});
</script>
