module.exports = {
  plugins: ['@typescript-eslint'],
  extends: [
    'eslint:recommended',
    'plugin:@typescript-eslint/eslint-recommended',
    'plugin:@typescript-eslint/recommended',
    'prettier/@typescript-eslint'
  ],
  parser: '@typescript-eslint/parser',
  parserOptions: {
    sourceType: 'module'
  },
  rules: {
    '@typescript-eslint/explicit-function-return-type': 'off'
  },
  env: {
    node: true,
    es6: true
  }
}
