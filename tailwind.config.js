/** @type {import('tailwindcss').Config} */

const defaultTheme = require("tailwindcss/defaultTheme");

module.exports = {
  content: ["internal/server/templates/*.html"],
  theme: {
    container: {
      // you can configure the container to be centered
      center: true,
      // or have default horizontal padding
      padding: '1rem',
      // default breakpoints but with 40px removed
      //
      // notice how the color changes at 768px but
      // the container size changes at 728px
      screens: {
        sm: '600px',
        md: '728px',
        lg: '984px',
        xl: '984px',
        '2xl': '984px',
      },
      extend: {
        fontFamily: {
          sans: ["Inter var", ...defaultTheme.fontFamily.sans],
        },
      },
    },
    plugins: [
      require("@tailwindcss/forms"),
      require("@tailwindcss/typography"),
    ],
  }
}