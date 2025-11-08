import js from "@eslint/js";
import globals from "globals";
import tseslint from "typescript-eslint";
import { defineConfig } from "eslint/config";

export default defineConfig([
	{
		files: ["ts/*.{js,mjs,cjs,ts,mts,cts}"],
		plugins: { js },
		extends: ["js/recommended"],
	},
	{
		files: ["ts/*.{js,mjs,cjs,ts,mts,cts}"],
		languageOptions: { globals: globals.browser },
		rules: {
			quotes: ["error", "double"],
			semi: ["error", "always"],
			"no-unused-vars": "off",
			"@typescript-eslint/no-unused-vars": [
				"warn",
				{
					argsIgnorePattern: "^_",
					varsIgnorePattern: "^_",
					destructuredArrayIgnorePattern: "^_",
				},
			],
			"@typescript-eslint/consistent-type-imports": [
				"warn",
				{ prefer: "type-imports", fixStyle: "inline-type-imports" },
			],
			"@typescript-eslint/no-unsafe-declaration-merging": "warn",
			"@typescript-eslint/prefer-nullish-coalescing": "off",
			"@typescript-eslint/no-confusing-void-expression": "off",
			"@typescript-eslint/restrict-template-expressions": "off",
			"@typescript-eslint/explicit-function-return-type": "error",
			"no-restricted-syntax": [
				"error",
				{
					selector: "ConditionalExpression",
					message: "Ternary operators are evil.",
				},
			],
		},
	},
	tseslint.configs.recommended,
]);
