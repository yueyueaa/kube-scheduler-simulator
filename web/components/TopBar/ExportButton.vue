<template>
  <v-dialog v-model="data.dialog" >
    <v-card>
      <v-divider></v-divider>
      <v-card-actions>
        <v-spacer></v-spacer>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts">
import { saveAs } from "file-saver";
import { defineComponent, inject, reactive } from "@nuxtjs/composition-api";
import SnackBarStoreKey from "../StoreKey/SnackBarStoreKey";
import yaml from "js-yaml";
import { ExportAPIKey } from "~/api/APIProviderKeys";

export default defineComponent({
  setup() {
    const data = reactive({
      dialog: false,
    });

    const exportAPI = inject(ExportAPIKey);
    if (!exportAPI) {
      throw new Error(`${ExportAPIKey.description} is not provided`);
    }

    const snackbarstore = inject(SnackBarStoreKey);
    if (!snackbarstore) {
      throw new Error(`${SnackBarStoreKey.description} is not provided`);
    }
    const setServerErrorMessage = (error: string) => {
      snackbarstore.setServerErrorMessage(error);
    };

    const ExportScheduler = async () => {
      try {
        const c = await exportAPI.exportScheduler();
        if (c) {
          const blob = new Blob([yaml.dump(c)], {
            type: "application/yaml",
          });
          saveAs(blob, "export.yml");
          data.dialog = false;
        }
      } catch (e: any) {
        setServerErrorMessage(e);
      }
    };
    return {
      data,
      ExportScheduler,
    };
  },
});
</script>
