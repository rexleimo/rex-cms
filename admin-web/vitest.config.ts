import {configDefaults,defineConfig} from 'vitest/config';

export default defineConfig({
  test: {
    ...configDefaults,
    environment:"jsdom",
    globals:true,
    setupFiles: ['./src/setup.ts'],
    css: false
  }
})