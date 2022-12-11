const colors = require('tailwindcss/colors');

const rainbowColors = {
  '50': {
    DEFAULT: '#f6fdff',
    dark: '#f4f8f8',
  },
  '100': {
    DEFAULT: '#ecfcff',
    dark: '#e9f0f2',
  },
  '200': {
    DEFAULT: '#d0f7fe',
    dark: '#c8dade',
  },
  '300': {
    DEFAULT: '#b4f1fe',
    dark: '#a6c4ca',
  },
  '400': {
    DEFAULT: '#7ce7fd',
    dark: '#6498a3',
  },
  '500': {
    DEFAULT: '#44ddfc',
    dark: '#216c7b',
  },
  '600': {
    DEFAULT: '#3dc7e3',
    dark: '#1e616f',
  },
  '700': {
    DEFAULT: '#33a6bd',
    dark: '#19515c',
  },
  '800': {
    DEFAULT: '#298597',
    dark: '#14414a',
  },
  '900': {
    DEFAULT: '#216c7b',
    dark: '#10353c',
  }
};

module.exports = {
  content: [
    './index.html',
    './src/**/*.{vue,js,ts,jsx,tsx}',
    './node_modules/@snowind/**/*.{vue,js,ts}',
  ],
  options: {
    safelist: ['hidden', 'block', 'sm:block', 'sm:hidden'],
  },
  darkMode: 'class',

  plugins: [
    require('@tailwindcss/forms'),
    require('@snowind/plugin'),
    require('tailwindcss-semantic-colors'),
  ],

  theme: {
    extend: {
      semanticColors: {
        background: {
          light: {
            bg: colors.white,
            txt: colors.gray[600],
          },
          dark: {
            bg: "#111827",
            txt: colors.neutral[100]
          }
        },
        secondary: {
          light: {
            bg: '#298597',
            txt: colors.white,
          },
          dark: {
            bg: '#14414a',
            txt: colors.neutral[100]
          }
        }
      },
      colors: {
        'rainbow': rainbowColors,
        'primary': {
          DEFAULT: rainbowColors['700'].DEFAULT,
          dark: colors.gray[700],
        },
        'primary-r': {
          DEFAULT: "#FFFFFF",
          dark: "#FFFFFF",
        },
        'secondary': {
          DEFAULT: rainbowColors['700'].DEFAULT,
          dark: rainbowColors['700'].dark,
        },
        'secondary-r': {
          DEFAULT: "#FFFFFF",
          dark: "#FFFFFF",
        },
        'header': {
          DEFAULT: colors.white,
          dark: "#0e2b32"
        },
        'bid': {
          DEFAULT: '#00DD66',
          dark: '#008822',
        },
        'ask': {
          DEFAULT: '#DD0066',
          dark: '#880022',
        },
        'tr-border': {
          DEFAULT: colors.white,
          dark: '#ffffff0a',
        },
      }
    },
  },
}
