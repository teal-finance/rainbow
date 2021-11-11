const colors = require('tailwindcss/colors');

const rainbowColors = {
  '50': '#f6fdff',
  '100': '#ecfcff',
  '200': '#d0f7fe',
  '300': '#b4f1fe',
  '400': '#7ce7fd',
  '500': '#44ddfc',
  '600': '#3dc7e3',
  '700': '#33a6bd',
  '800': '#298597',
  '900': '#216c7b'
};

const rainbowDarkColors = {
  '50': '#f4f8f8',
  '100': '#e9f0f2',
  '200': '#c8dade',
  '300': '#a6c4ca',
  '400': '#6498a3',
  '500': '#216c7b',
  '600': '#1e616f',
  '700': '#19515c',
  '800': '#14414a',
  '900': '#10353c'
}

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
        'rainbow': rainbowColors,
        'primary': {
          DEFAULT: rainbowColors['500'],
          dark: colors.gray[700],
        },
        'primary-r': {
          DEFAULT: "#FFFFFF",
          dark: "#FFFFFF",
        },
        'secondary': {
          DEFAULT: rainbowColors['700'],
          dark: rainbowDarkColors['700'],
        },
        'secondary-r': {
          DEFAULT: "#FFFFFF",
          dark: "#FFFFFF",
        },
      }
    },
  },
}
