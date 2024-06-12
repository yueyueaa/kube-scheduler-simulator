<template>
  <v-sheet class="transparent">
    <v-btn
      v-for="(rn, i) in resourceNames"
      :key="i"
      color="primary ma-2"
      dark
      @click="create(rn)"
    >
      新{{ rn }}
    </v-btn>
  </v-sheet>
</template>

<script lang="ts">
import { ref, inject, defineComponent } from "@nuxtjs/composition-api";
import {
  podTemplate,
  nodeTemplate,
  namespaceTemplate,
} from "./lib/template";
import {} from "./lib/util";
import PodStoreKey from "./StoreKey/PodStoreKey";
import NodeStoreKey from "./StoreKey/NodeStoreKey";
import NamespaceStoreKey from "./StoreKey/NamespaceStoreKey";
import {
  V1Node,
  V1Pod,
  V1Namespace,
} from "@kubernetes/client-node";

type Resource =
  | V1Pod
  | V1Node
  | V1Namespace;

interface Store {
  readonly selected: object | null;
  readonly count: number;
  select(_resource: Resource | null, _isNew: boolean): void;
}

export default defineComponent({
  setup() {
    var store: Store | null = null;

    const podstore = inject(PodStoreKey);
    if (!podstore) {
      throw new Error(`${PodStoreKey.description} is not provided`);
    }

    const nodestore = inject(NodeStoreKey);
    if (!nodestore) {
      throw new Error(`${NodeStoreKey.description} is not provided`);
    }

    const namespacestore = inject(NamespaceStoreKey);
    if (!namespacestore) {
      throw new Error(`${namespacestore} is not provided`);
    }

    const dialog = ref(false);
    const resourceNames = [
      "服务器节点",
      "任务容器",
      "命名空间",
    ];

    const create = (rn: string) => {
      var targetTemplate: Resource | null = null;
      switch (rn) {
        case "任务容器":
          store = podstore;
          targetTemplate = podTemplate();
          break;
        case "服务器节点":
          store = nodestore;
          targetTemplate = nodeTemplate();
          break;
        case "命名空间":
          store = namespacestore;
          targetTemplate = namespaceTemplate();
          break;
      }

      if (store) {
        store.select(targetTemplate, true);
      }
      dialog.value = false;
    };

    return {
      create,
      dialog,
      resourceNames,
    };
  },
});
</script>
