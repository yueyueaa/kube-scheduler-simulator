<template>
</template>

<script lang="ts">
import { V1PriorityClass } from "@kubernetes/client-node";
import { computed, inject, defineComponent } from "@nuxtjs/composition-api";
import {} from "../../lib/util";
import PriorityClassStoreKey from "../../StoreKey/PriorityClassStoreKey";
export default defineComponent({
  setup() {
    const store = inject(PriorityClassStoreKey);
    if (!store) {
      throw new Error(`${PriorityClassStoreKey.description} is not provided`);
    }

    const onClick = (priorityclass: V1PriorityClass) => {
      store.select(priorityclass, false);
    };

    const priorityclasses = computed(() => store.priorityclasses);
    return {
      priorityclasses,
      onClick,
    };
  },
});
</script>
