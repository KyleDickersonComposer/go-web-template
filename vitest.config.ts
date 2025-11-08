/// <reference types="vitest" />
import { defineConfig } from "vitest/config";

export default defineConfig({
	test: {
		environment: "jsdom",
		globals: true,
		include: ["ts/**/**.test.ts"],
		exclude: [
			"node_modules",
			"dist",
			"**/*.js",
			"**/*.d.ts",
			"**/__tests__/**",
		],
	},
});
