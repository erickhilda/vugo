import { createApp } from "vue";
import { createPinia } from "pinia";
import PrimeVue from "primevue/config";
import Nora from "@primeuix/themes/nora";
import { definePreset } from "@primeuix/themes";

import App from "./App.vue";
import router from "./router";

import "primeicons/primeicons.css";
import "./assets/css/main.css";

// Custom brand preset based on Nora
const BrandPreset = definePreset(Nora, {
  semantic: {
    primary: {
      50: "{green.50}",
      100: "{green.100}",
      200: "{green.200}",
      300: "{green.300}",
      400: "{green.400}",
      500: "{green.500}",
      600: "{green.600}",
      700: "{green.700}",
      800: "{green.800}",
      900: "{green.900}",
      950: "{green.950}",
    },
  },
  primitive: {
    green: {
      50: "#e6f2ef",
      100: "#cce5df",
      200: "#99ccbf",
      300: "#66b29f",
      400: "#33997f",
      500: "#004838",
      600: "#003d30",
      700: "#003328",
      800: "#002820",
      900: "#001e18",
      950: "#000f0c",
    },
  },
});

const app = createApp(App);

app.use(createPinia());
app.use(router);
app.use(PrimeVue, {
  theme: {
    preset: BrandPreset,
    options: {
      darkModeSelector: "false",
    },
  },
});

app.mount("#app");
