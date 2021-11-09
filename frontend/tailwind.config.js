const colors = require('tailwindcss/colors');

module.exports = {
  purge: {
    content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
    options: {
      safelist: ['hidden', 'block', 'sm:block', 'sm:hidden'],
    }
  },
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
          dark: colors.gray[700],
        },
        'primary-r': {
          DEFAULT: "#FFFFFF",
          dark: "#FFFFFF",
        },
      }
    },
  },
}
