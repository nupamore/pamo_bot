module.exports = {
    root: true,
    env: {
        browser: true,
        es6: true
    },
    extends: [
        "eslint:recommended",
        "plugin:vue/essential",
        "prettier",
        "plugin:prettier/recommended"
    ],
    globals: {
        Atomics: "readonly",
        SharedArrayBuffer: "readonly",
        process: "readonly",
        require: "readonly",
        __dirname: "readonly",
        module: "readonly",
        BASE_URL: "readonly",
        USER_API_URL: "readonly"
    },
    parserOptions: {
        ecmaVersion: 2018,
        sourceType: "module"
    },
    plugins: ["prettier"],
    rules: {
        "prettier/prettier": [
            "error",
            {
                singleQuote: true,
                semi: false,
                useTabs: false,
                tabWidth: 4,
                trailingComma: "all",
                printWidth: 80,
                bracketSpacing: true,
                arrowParens: "avoid"
            }
        ],
        "no-console": process.env.NODE_ENV === "production" ? "error" : "off"
    },
    overrides: [
        {
            files: ["build/*.js"],
            globals: {
                process: "readonly",
                module: "readonly",
                __dirname: "readonly"
            }
        }
    ]
};
