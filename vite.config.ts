// vite.config.ts
import { defineConfig } from "vite";

export default defineConfig({
	root: ".",
	publicDir: "static",
	build: {
		outDir: "dist",
		emptyOutDir: false,
		minify: false,
		rollupOptions: {
			input: {
				main: "ts/main.ts",
			},
			output: {
				entryFileNames: "[name].js",
			},
		},
	},
});
