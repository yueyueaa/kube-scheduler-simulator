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
import {
  defineComponent,
  inject,
  reactive,
  watch,
} from "@nuxtjs/composition-api";
import { ResourcesForImport } from "~/api/v1/export";
import yaml from "js-yaml";
import SnackBarStoreKey from "../StoreKey/SnackBarStoreKey";
import { ExportAPIKey } from "~/api/APIProviderKeys";

interface SelectedItem {
  dialog: boolean;
  filedata: ResourcesForImport;
  isImportButtonDisabled: boolean;
}

export default defineComponent({
  setup() {
    const data = reactive({
      dialog: false,
      isImportButtonDisabled: true,
    } as SelectedItem);

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

    watch(data, (newValue, _) => {
      if (!newValue.dialog) {
        // reset filedata
        data.filedata = {} as ResourcesForImport;
        data.isImportButtonDisabled = true;
      }
    });

    const ImportScheduler = async () => {
      exportAPI
        .importScheduler(data.filedata as ResourcesForImport)
        .catch((e) => setServerErrorMessage(e))
        .finally(() => {
          data.dialog = false;
        });
    };
    function readfile(e: { target: { files: FileList | null } }) {
      if (e.target.files === null) return;
      const file = e.target.files[0];
      const reader = new FileReader();

      reader.onload = function () {
        try {
          const filedata = <ResourcesForImport>yaml.load(
            reader.result as string
          );
          data.filedata = filedata;
          data.isImportButtonDisabled = false;
        } catch (e) {
          setServerErrorMessage("Failed to load the selected file.");
        }
      };
      reader.onabort = function () {
        console.log("file read aborted");
      };
      reader.readAsText(file);
    }

    return {
      data,
      ImportScheduler,
      readfile,
    };
  },
});
</script>
