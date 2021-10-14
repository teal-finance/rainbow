module.exports = {
  purge: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  // purge: false
  // purge is disabled to be able to deploy the storybook on GitHub pages

  darkMode: 'class',

  jit: true,

  plugins: [
    require('@tailwindcss/forms'),
    require('@snowind/plugin'),
  ],

  theme: {
    extend: {
      colors: {
        'primary': {
          DEFAULT: "#44ddfc",
          dark: "#44ddfc",
        },
        'primary-r': {
          DEFAULT: "#FFFFFF",
          dark: "#FFFFFF",
        },
      }
    },
  },
}
